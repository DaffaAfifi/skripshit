import { loginUserValidation } from "../validation/user-validation.js";
import { validate } from "../validation/validation.js";
import db from "../application/database.js";
import bcrypt from "bcrypt";
import { ResponseError } from "../response/response-error.js";
import jwt from "jsonwebtoken";
import "dotenv/config";

// Fungsi untuk login pengguna
const login = async (req) => {
  const connection = await db.promise().getConnection();
  try {
    const loginRequest = validate(loginUserValidation, req);

    const [rows] = await connection.query(
      "SELECT id, email, nama, role_id, password FROM users WHERE email = ?",
      [loginRequest.email]
    );

    if (rows.length === 0) {
      throw new ResponseError(400, "Username or password wrong");
    }

    const user = rows[0];

    const hash = user.password.replace("$2y$", "$2a$");
    const isPasswordValid = await bcrypt.compare(loginRequest.password, hash);

    if (!isPasswordValid) {
      throw new ResponseError(400, "Username or password wrong");
    }

    const token = jwt.sign(
      { id: user.id, email: user.email, nama: user.nama, role: user.role },
      process.env.JWT_SECRET,
      { expiresIn: "2h" }
    );

    await connection.query(
      "INSERT INTO sessions (token, user_id, expiry, created_at, updated_at) VALUES (?, ?, ?, ?, ?)",
      [token, user.id, new Date(Date.now() + 7200000), new Date(), new Date()]
    );

    return token;
  } catch (error) {
    throw new ResponseError(400, error.message);
  } finally {
    connection.release();
  }
};

// Fungsi untuk logout pengguna
const logout = async (token) => {
  const connection = await db.promise().getConnection();
  try {
    const result = await connection.query(
      "DELETE FROM sessions WHERE token = ?",
      [token]
    );
    return result;
  } catch (error) {
    throw new ResponseError(400, error.message);
  } finally {
    connection.release();
  }
};

export default {
  login,
  logout,
};
