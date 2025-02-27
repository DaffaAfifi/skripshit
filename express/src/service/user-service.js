import {
  createUserValidation,
  updateUserValidation,
} from "../validation/user-validation.js";
import { validate } from "../validation/validation.js";
import db from "../application/database.js";
import bcrypt from "bcrypt";
import { ResponseError } from "../response/response-error.js";
import "dotenv/config";

// Fungsi untuk mendapatkan semua pengguna dari database
const getUsers = async (page, limit) => {
  const connection = await db.promise().getConnection();
  try {
    const offset = (page - 1) * limit;
    const [rows] = await connection.query(
      "SELECT nama, email, NIK, alamat, telepon, jenis_kelamin, kepala_keluarga, tempat_lahir, tanggal_lahir, jenis_usaha FROM users LIMIT ? OFFSET ?",
      [limit, offset]
    );
    return rows;
  } catch (error) {
    throw new ResponseError(400, error.message);
  } finally {
    connection.release();
  }
};

// Fungsi untuk membuat pengguna baru
const createUser = async (req, res) => {
  const connection = await db.promise().getConnection();
  try {
    const user = validate(createUserValidation, req);

    user.password = await bcrypt.hash(user.password, 10);

    const [result] = await connection.query(
      `INSERT INTO users (nama, email, password, NIK, alamat, telepon, jenis_kelamin, kepala_keluarga, tempat_lahir, tanggal_lahir, jenis_usaha, created_at, updated_at)
         VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
         ON DUPLICATE KEY UPDATE email = email`,
      [
        user.nama,
        user.email,
        user.password,
        user.NIK,
        user.alamat,
        user.telepon,
        user.jenis_kelamin,
        user.kepala_keluarga,
        user.tempat_lahir,
        user.tanggal_lahir,
        user.jenis_usaha,
        new Date(),
        new Date(),
      ]
    );

    if (result.affectedRows === 0) {
      throw new ResponseError(400, "Email or NIK already exists");
    }

    return result;
  } catch (error) {
    throw new ResponseError(400, error.message);
  } finally {
    connection.release();
  }
};

// Fungsi untuk mendapatkan data pengguna berdasarkan ID
const getUserById = async (id) => {
  const connection = await db.promise().getConnection();
  try {
    const [rows] = await connection.query(
      "SELECT nama, email, NIK, alamat, telepon, jenis_kelamin, kepala_keluarga, tempat_lahir, tanggal_lahir, jenis_usaha FROM users WHERE id = ?",
      [id]
    );

    if (rows.length === 0) {
      throw new ResponseError(404, "User not found");
    }

    return rows[0];
  } catch (error) {
    throw new ResponseError(400, error.message);
  } finally {
    connection.release();
  }
};

// Fungsi untuk mendapatkan berita yang disimpan oleh pengguna berdasarkan ID
const getSavedNews = async (id) => {
  const connection = await db.promise().getConnection();
  try {
    const [rows] = await connection.query(
      `SELECT 
          users.id, users.nama, users.email, news.id AS news_id, news.gambar, news.judul, news.subjudul, news.isi, news.created_at 
        FROM users 
        INNER JOIN saved_news ON users.id = saved_news.user_id 
        INNER JOIN news ON saved_news.news_id = news.id 
        WHERE users.id = ?`,
      [id]
    );

    if (rows.length === 0) {
      throw new ResponseError(404, "User or saved news not found");
    }

    const payload = {
      id: rows[0].id,
      nama: rows[0].nama,
      email: rows[0].email,
      berita_tersimpan: rows.map((row) => ({
        id: row.news_id,
        gambar: row.gambar,
        judul: row.judul,
        subjudul: row.subjudul,
        isi: row.isi,
        created_at: new Date(row.created_at).toLocaleDateString("en-GB"),
      })),
    };

    return payload;
  } catch (error) {
    throw new ResponseError(400, error.message);
  } finally {
    connection.release();
  }
};

// Mendapatkan fasilitas user (sertifikat, pelatihan, bantuan, dan alat)
const getFacilities = async (id) => {
  const connection = await db.promise().getConnection();
  try {
    const [rows] = await connection.query(
      `SELECT
          users.id, users.email, 
          sertificates.id AS id_sertifikat, sertificates.nama AS nama_sertifikat, user_sertificates.no_sertifikat, sertificates.tanggal_terbit, sertificates.kadaluarsa, sertificates.keterangan,
          trainings.id AS id_pelatihan, trainings.nama AS nama_pelatihan, trainings.penyelenggara, trainings.tanggal_pelaksanaan, trainings.tempat,
          assistance.id AS id_bantuan, assistance.nama AS nama_bantuan, assistance.koordinator, assistance.sumber_anggaran, assistance.total_anggaran, assistance.tahun_pemberian,
          assistance_tools.kuantitas,
          tools.id AS id_alat, tools.nama_item, tools.harga, tools.deskripsi
        FROM users
        LEFT JOIN user_sertificates ON users.id = user_sertificates.user_id
        LEFT JOIN sertificates ON user_sertificates.sertificates_id = sertificates.id
        LEFT JOIN user_trainings ON users.id = user_trainings.user_id
        LEFT JOIN trainings ON user_trainings.trainings_id = trainings.id
        LEFT JOIN assistance ON users.id = assistance.user_id
        LEFT JOIN assistance_tools ON assistance.id = assistance_tools.assistance_id
        LEFT JOIN tools ON assistance_tools.tools_id = tools.id
        WHERE users.id = ?`,
      [id]
    );

    if (rows.length === 0) {
      throw new ResponseError(404, "User or facilities not found");
    }

    const result = {
      id: rows[0].id,
      nama: rows[0].nama,
      email: rows[0].email,
      sertifikat: [],
      pelatihan: [],
      bantuan: [],
    };

    const helpMap = {};

    rows.forEach((row) => {
      if (
        row.id_sertifikat &&
        !result.sertifikat.some((c) => c.id === row.id_sertifikat)
      ) {
        result.sertifikat.push({
          id: row.id_sertifikat,
          nama: row.nama_sertifikat,
          no_sertifikat: row.no_sertifikat,
          tanggal_terbit: row.tanggal_terbit,
          kadaluarsa: row.kadaluarsa,
          keterangan: row.keterangan,
        });
      }

      if (
        row.id_pelatihan &&
        !result.pelatihan.some((t) => t.id === row.id_pelatihan)
      ) {
        result.pelatihan.push({
          id: row.id_pelatihan,
          nama: row.nama_pelatihan,
          koordinator: row.penyelenggara,
          tanggal_pelaksanaan: row.tanggal_pelaksanaan,
          tempat: row.tempat,
        });
      }

      if (row.id_bantuan) {
        if (!helpMap[row.id_bantuan]) {
          helpMap[row.id_bantuan] = {
            id: row.id_bantuan,
            nama: row.nama_bantuan,
            koordinator: row.koordinator,
            sumber_anggaran: row.sumber_anggaran,
            tahun_pemberian: row.tahun_pemberian,
            total_anggaran: row.total_anggaran,
            alat: [],
          };
        }
        if (
          row.id_alat &&
          !helpMap[row.id_bantuan].alat.some((tool) => tool.id === row.id_alat)
        ) {
          helpMap[row.id_bantuan].alat.push({
            id: row.id_alat,
            nama: row.nama_item,
            harga: row.harga,
            kuantitas: row.kuantitas,
          });
        }
      }
    });

    result.bantuan = Object.values(helpMap);

    return result;
  } catch (error) {
    throw new ResponseError(400, error.message);
  } finally {
    connection.release();
  }
};

// Mengupdate data user
const updateUser = async (id, data) => {
  const connection = await db.promise().getConnection();
  const userUpdates = validate(updateUserValidation, data);

  const updates = [];
  const values = [];

  Object.keys(userUpdates).forEach((key) => {
    updates.push(`${key} = ?`);
    values.push(userUpdates[key]);
  });

  if (updates.length === 0) {
    throw new ResponseError(400, "No valid fields to update");
  }

  const query = `UPDATE users SET ${updates.join(", ")} WHERE id = ?`;
  values.push(id);

  try {
    const result = await connection.query(query, values);

    if (result.affectedRows === 0) {
      throw new ResponseError(404, "User not found");
    }

    return result;
  } catch (error) {
    throw new ResponseError(400, error.message);
  } finally {
    connection.release();
  }
};

// Mendapatkan komentar dari berita yang disimpan oleh user
const getSavedNewsComment = async (id) => {
  const connection = await db.promise().getConnection();
  try {
    const [rows] = await connection.query(
      `SELECT 
          users.id, users.nama, users.email, news.id AS news_id, news.gambar, news.judul, news.subjudul, news.isi, news.created_at, comments.news_id AS comment_id, comments.comment, comments.created_at
        FROM users
        LEFT JOIN saved_news ON users.id = saved_news.user_id
        LEFT JOIN news ON saved_news.news_id = news.id
        LEFT JOIN comments ON news.id = comments.news_id
        WHERE users.id = ?`,
      [id]
    );

    if (rows.length === 0) {
      throw new ResponseError(404, "User or saved news not found");
    }

    const result = {
      id: rows[0].id,
      nama: rows[0].nama,
      email: rows[0].email,
      news: [],
    };

    const newsMap = {};

    rows.forEach((row) => {
      if (!newsMap[row.news_id]) {
        newsMap[row.news_id] = {
          id: row.news_id,
          gambar: row.gambar,
          judul: row.judul,
          subjudul: row.subjudul,
          isi: row.isi,
          created_at: row.created_at,
          comment: [],
        };
      }
      if (row.comment_id) {
        newsMap[row.news_id].comment.push({
          comment: row.comment,
          created_at: row.created_at,
        });
      }
    });

    result.news = Object.values(newsMap);

    return result;
  } catch (error) {
    throw new ResponseError(400, error.message);
  } finally {
    connection.release();
  }
};

export default {
  getUsers,
  createUser,
  getUserById,
  getSavedNews,
  getFacilities,
  updateUser,
  getSavedNewsComment,
};
