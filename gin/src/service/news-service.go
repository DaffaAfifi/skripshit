package service

import (
	"database/sql"
	"gin-project/src/model"
	"gin-project/src/response"
)

func GetNewsComments(id string, db *sql.DB) (model.NewsComments, error) {
	var newsComments model.NewsComments

	query := `SELECT 
          news.id, news.gambar, news.judul, news.subjudul, news.isi, news.created_at, 
          comments.user_id AS user_id, comments.comment, comments.created_at AS comment_created_at
        FROM news 
        LEFT JOIN comments ON news.id = comments.news_id 
        WHERE news.id = ?`

	stmt, err := db.Prepare(query)
	if err != nil {
		return newsComments, response.NewResponseError(400, err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		return newsComments, response.NewResponseError(400, err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var comments model.Comments

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
			return newsComments, response.NewResponseError(400, err.Error())
		}

		newsComments.Comments = append(newsComments.Comments, comments)
	}

	return newsComments, nil
}
