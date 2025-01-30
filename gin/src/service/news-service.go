package service

import (
	"database/sql"
	"gin-project/src/model"
	"gin-project/src/response"
)

// GetNewsComments mengambil komentar-komentar terkait sebuah berita berdasarkan id berita
func GetNewsComments(id string, db *sql.DB) (model.NewsComments, error) {
	var newsComments model.NewsComments

	// Query untuk mendapatkan berita beserta komentar yang terkait
	query := `SELECT 
          news.id, news.gambar, news.judul, news.subjudul, news.isi, news.created_at, 
          comments.user_id AS user_id, comments.comment, comments.created_at AS comment_created_at
        FROM news 
        LEFT JOIN comments ON news.id = comments.news_id 
        WHERE news.id = ?`

	// Menyiapkan statement SQL
	stmt, err := db.Prepare(query)
	if err != nil {
		// Mengembalikan error jika terjadi kesalahan saat menyiapkan query
		return newsComments, response.NewResponseError(400, err.Error())
	}
	defer stmt.Close()

	// Mengeksekusi query dengan id berita
	rows, err := stmt.Query(id)
	if err != nil {
		// Mengembalikan error jika terjadi kesalahan saat menjalankan query
		return newsComments, response.NewResponseError(400, err.Error())
	}
	defer rows.Close()

	// Menyimpan hasil query ke dalam newsComments
	for rows.Next() {
		var comments model.Comments

		// Memindahkan hasil query ke dalam struktur data newsComments dan comments
		if err := rows.Scan(
			&newsComments.NewsId,
			&newsComments.Gambar,
			&newsComments.Judul,
			&newsComments.Subjudul,
			&newsComments.Isi,
			&newsComments.CreatedAt,
			&comments.User,
			&comments.Comment,
			&comments.CreatedAt,
		); err != nil {
			// Mengembalikan error jika terjadi kesalahan saat memindahkan data
			return newsComments, response.NewResponseError(400, err.Error())
		}

		// Menambahkan komentar ke dalam daftar komentar berita
		newsComments.Comments = append(newsComments.Comments, comments)
	}

	// Mengembalikan data berita beserta komentar yang terkait
	return newsComments, nil
}
