package service

import (
	"database/sql"
	"fmt"
	"gin-project/src/model"
	"gin-project/src/response"
	"reflect"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// GetUsers mengambil semua pengguna dari database
func GetUsers(db *sql.DB, page int, limit int) ([]model.User, error) {
	var users []model.User
	offset := (page - 1) * limit
	query := "SELECT nama, email, NIK, alamat, telepon, jenis_kelamin, kepala_keluarga, tempat_lahir, tanggal_lahir, jenis_usaha FROM users LIMIT ? OFFSET ?"

	stmt, err := db.Prepare(query)
	if err != nil {
		return users, response.NewResponseError(500, "Failed to prepare statement")
	}
	defer stmt.Close()

	rows, err := stmt.Query(limit, offset)
	if err != nil {
		return users, response.NewResponseError(400, err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.Name, &user.Email, &user.NIK, &user.Alamat, &user.Telepon, &user.Jenis_kelamin, &user.Kepala_keluarga, &user.Tempat_lahir, &user.Tanggal_lahir, &user.Jenis_usaha); err != nil {
			return users, response.NewResponseError(400, err.Error())
		}
		users = append(users, user)
	}

	return users, nil
}

// GetUserById mengambil data pengguna berdasarkan id
func GetUserById(id string, db *sql.DB) (model.User, error) {
	var user model.User
	query := "SELECT nama, email, NIK, alamat, telepon, jenis_kelamin, kepala_keluarga, tempat_lahir, tanggal_lahir, jenis_usaha FROM users WHERE id = ?"

	stmt, err := db.Prepare(query)
	if err != nil {
		return user, response.NewResponseError(500, "Failed to prepare statement")
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&user.Name, &user.Email, &user.NIK, &user.Alamat, &user.Telepon, &user.Jenis_kelamin, &user.Kepala_keluarga, &user.Tempat_lahir, &user.Tanggal_lahir, &user.Jenis_usaha)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, response.NewResponseError(404, "User not found")
		}
		return user, response.NewResponseError(400, err.Error())
	}

	return user, nil
}

// CreateUser membuat pengguna baru di database
func CreateUser(request model.CreateUserRequest, db *sql.DB) error {
	newPassword, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	stmt, err := db.Prepare("INSERT INTO users (nama, email, password, NIK, alamat, telepon, jenis_kelamin, kepala_keluarga, tempat_lahir, tanggal_lahir, jenis_usaha) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return response.NewResponseError(500, err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		request.Nama,
		request.Email,
		newPassword,
		request.NIK,
		request.Alamat,
		request.Telepon,
		request.Jenis_kelamin,
		request.Kepala_keluarga,
		request.Tempat_lahir,
		request.Tanggal_lahir,
		request.Jenis_usaha,
	)
	if err != nil {
		return response.NewResponseError(400, err.Error())
	}

	return nil
}

func UpdateUser(id string, request model.UpdateUserRequest, db *sql.DB) error {
	updates := []string{}
	values := []interface{}{}

	val := reflect.ValueOf(request)
	typ := reflect.TypeOf(request)

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		value := val.Field(i).Interface()

		jsonTag := field.Tag.Get("json")
		if jsonTag == "" {
			jsonTag = field.Name
		}

		if !isZero(value) {
			updates = append(updates, fmt.Sprintf("%s = ?", jsonTag))
			values = append(values, value)
		}
	}

	if len(updates) == 0 {
		return response.NewResponseError(400, "No valid fields to update")
	}

	query := fmt.Sprintf("UPDATE users SET %s WHERE id = ?", strings.Join(updates, ", "))
	values = append(values, id)

	stmt, err := db.Prepare(query)
	if err != nil {
		return response.NewResponseError(500, "Failed to prepare statement: "+err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(values...)
	if err != nil {
		return response.NewResponseError(400, "Failed to execute update query: "+err.Error())
	}

	return nil
}

// isZero memeriksa apakah nilai dari field adalah nilai default (kosong)
func isZero(value interface{}) bool {
	return reflect.DeepEqual(value, reflect.Zero(reflect.TypeOf(value)).Interface())
}

// GetSavedNews mengambil berita yang disimpan oleh pengguna berdasarkan id pengguna
func GetSavedNews(id string, db *sql.DB) (model.UserSavedNews, error) {
	var userSavedNews model.UserSavedNews
	var newsList []model.News

	query := `SELECT
				u.nama, u.email, n.id, n.gambar, n.judul, n.subjudul, n.isi, n.created_at
			FROM users u
			INNER JOIN saved_news sn ON u.id = sn.user_id
			INNER JOIN news n ON sn.news_id = n.id
			WHERE u.id = ?`

	stmt, err := db.Prepare(query)
	if err != nil {
		return userSavedNews, response.NewResponseError(500, "Failed to prepare statement")
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		return userSavedNews, response.NewResponseError(400, err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var news model.News
		if err := rows.Scan(&userSavedNews.Name, &userSavedNews.Email, &news.NewsId, &news.Gambar, &news.Judul, &news.Subjudul, &news.Isi, &news.CreatedAt); err != nil {
			return userSavedNews, response.NewResponseError(400, err.Error())
		}
		newsList = append(newsList, news)
	}

	if err := rows.Err(); err != nil {
		return userSavedNews, response.NewResponseError(500, err.Error())
	}

	userSavedNews.News = newsList
	return userSavedNews, nil
}

// GetUserSavedNewsComment mengambil berita yang disimpan dan komentar terkait untuk pengguna dari database.
func GetUserSavedNewsComment(id string, db *sql.DB) (model.UserSavedNewsComment, error) {
	var userSavedNewsComment model.UserSavedNewsComment
	var newsCommentsMap = make(map[string]*model.NewsComments)

	query := `SELECT
				users.nama, users.email, news.id, news.gambar, news.judul, news.subjudul, news.isi, news.created_at, comments.comment, comments.user_id, comments.created_at
			FROM users
			LEFT JOIN saved_news ON users.id = saved_news.user_id
			LEFT JOIN news ON saved_news.news_id = news.id
			LEFT JOIN comments ON news.id = comments.news_id
			WHERE users.id = ?`

	stmt, err := db.Prepare(query)
	if err != nil {
		return userSavedNewsComment, response.NewResponseError(400, err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		return userSavedNewsComment, response.NewResponseError(400, err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var newsComments model.NewsComments
		var comments model.Comments

		if err := rows.Scan(
			&userSavedNewsComment.Name,
			&userSavedNewsComment.Email,
			&newsComments.NewsId,
			&newsComments.Gambar,
			&newsComments.Judul,
			&newsComments.Subjudul,
			&newsComments.Isi,
			&newsComments.CreatedAt,
			&comments.Comment,
			&comments.User,
			&comments.CreatedAt,
		); err != nil {
			return userSavedNewsComment, response.NewResponseError(400, err.Error())
		}

		if existingNews, exists := newsCommentsMap[newsComments.NewsId]; exists {
			existingNews.Comments = append(existingNews.Comments, comments)
		} else {
			newsComments.Comments = append(newsComments.Comments, comments)
			newsCommentsMap[newsComments.NewsId] = &newsComments
		}
	}

	for _, news := range newsCommentsMap {
		userSavedNewsComment.NewsComments = append(userSavedNewsComment.NewsComments, *news)
	}

	return userSavedNewsComment, nil
}

// GetUserFacilities mengambil fasilitas pengguna (sertifikat, pelatihan, bantuan, alat) dari database.
func GetUserFacilities(id string, db *sql.DB) (model.UserFacilities, error) {
	var userFacilities model.UserFacilities
	helpMap := make(map[string]*model.Bantuan)
	sertifikatMap := make(map[string]bool)
	pelatihanMap := make(map[string]bool)

	query := `
		SELECT
			users.email, users.nama,
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
		WHERE users.id = ?
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		return userFacilities, response.NewResponseError(400, err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		return userFacilities, response.NewResponseError(400, err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var sertifikat model.Sertifikat
		var pelatihan model.Pelatihan
		var bantuan model.Bantuan
		var alat model.Alat
		var kuantitas sql.NullInt64

		err := rows.Scan(
			&userFacilities.Email, &userFacilities.Name,
			&sertifikat.Id, &sertifikat.Name, &sertifikat.No_sertifikat, &sertifikat.Tanggal_terbit, &sertifikat.Kadaluarsa, &sertifikat.Keterangan,
			&pelatihan.Id, &pelatihan.Name, &pelatihan.Penyelenggara, &pelatihan.Tanggal_pelaksanaan, &pelatihan.Tempat,
			&bantuan.Id, &bantuan.Name, &bantuan.Koordinator, &bantuan.Sumber_anggaran, &bantuan.Total_anggaran, &bantuan.Tahun_pemberian,
			&kuantitas, &alat.Id, &alat.Name, &alat.Harga, &alat.Deskripsi,
		)
		if err != nil {
			return userFacilities, response.NewResponseError(400, err.Error())
		}

		if sertifikat.Id.Valid {
			if _, exists := sertifikatMap[sertifikat.Id.String]; !exists {
				userFacilities.Sertifikat = append(userFacilities.Sertifikat, sertifikat)
				sertifikatMap[sertifikat.Id.String] = true
			}
		}

		if pelatihan.Id.Valid {
			if _, exists := pelatihanMap[pelatihan.Id.String]; !exists {
				userFacilities.Pelatihan = append(userFacilities.Pelatihan, pelatihan)
				pelatihanMap[pelatihan.Id.String] = true
			}
		}

		if bantuan.Id.Valid {
			bantuanIdStr := bantuan.Id.String
			if _, exists := helpMap[bantuanIdStr]; !exists {
				helpMap[bantuanIdStr] = &bantuan
			}

			if alat.Id.Valid {
				alat.Kuantitas.Valid = kuantitas.Valid
				alat.Kuantitas.Int64 = kuantitas.Int64

				if helpMap[bantuanIdStr].Alat == nil {
					helpMap[bantuanIdStr].Alat = []model.Alat{}
				}

				exists := false
				for _, existingAlat := range helpMap[bantuanIdStr].Alat {
					if existingAlat.Id == alat.Id {
						exists = true
						break
					}
				}
				if !exists {
					helpMap[bantuanIdStr].Alat = append(helpMap[bantuanIdStr].Alat, alat)
				}
			}
		}
	}

	for _, bantuan := range helpMap {
		if bantuan.Id.Valid {
			userFacilities.Bantuan = append(userFacilities.Bantuan, *bantuan)
		}
	}

	return userFacilities, nil
}
