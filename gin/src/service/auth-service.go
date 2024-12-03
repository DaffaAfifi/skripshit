package service

import (
	"database/sql"
	"gin-project/src/model"
	"gin-project/src/response"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Login(request model.LoginUserRequest, db *sql.DB) (string, error) {
	var id, role int
	var hashedPassword, nama, email string

	query := "SELECT id, email, nama, role_id, password FROM users WHERE email = ?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return "", response.NewResponseError(500, err.Error())
	}
	defer stmt.Close()

	err = stmt.QueryRow(request.Email).Scan(&id, &email, &nama, &role, &hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", response.NewResponseError(400, "Email or password is incorrect")
		}
		return "", response.NewResponseError(500, err.Error())
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(request.Password))
	if err != nil {
		return "", response.NewResponseError(400, "Email or password is incorrect")
	}

	expirationTime := time.Now().Add(2 * time.Hour)
	claims := &model.Claims{
		Id:    id,
		Email: email,
		Nama:  nama,
		Role:  role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET")
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", response.NewResponseError(500, err.Error())
	}

	insertQuery := `INSERT INTO sessions (token, user_id, expiry, created_at, updated_at) VALUES (?, ?, ?, ?, ?)`
	stmtInsert, err := db.Prepare(insertQuery)
	if err != nil {
		return "", response.NewResponseError(500, err.Error())
	}
	defer stmtInsert.Close()

	_, err = stmtInsert.Exec(tokenString, id, expirationTime, time.Now(), time.Now())
	if err != nil {
		return "", response.NewResponseError(500, err.Error())
	}

	return tokenString, nil
}

func Logout(token string, db *sql.DB) error {
	query := `DELETE FROM sessions WHERE token = ?`
	stmt, err := db.Prepare(query)
	if err != nil {
		return response.NewResponseError(500, err.Error())
	}
	defer stmt.Close()

	result, err := stmt.Exec(token)
	if err != nil {
		return response.NewResponseError(400, err.Error())
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return response.NewResponseError(500, err.Error())
	}

	if rowsAffected == 0 {
		return response.NewResponseError(404, "Token not found")
	}

	return nil
}
