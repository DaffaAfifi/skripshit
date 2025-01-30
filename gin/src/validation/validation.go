package validation

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

// Variabel untuk menyimpan instance dari validator
var validate *validator.Validate

// InitValidator menginisialisasi objek validator
func InitValidator() {
	// Membuat instance validator baru jika belum ada
	validate = validator.New()
}

// FormatValidationError memformat error validasi menjadi pesan yang lebih deskriptif
func FormatValidationError(err error) string {
	// Menyimpan daftar pesan error
	var errors []string
	// Melakukan type assertion untuk mendapatkan daftar error validasi
	for _, err := range err.(validator.ValidationErrors) {
		var message string
		// Mengecek tag validasi untuk menghasilkan pesan error yang sesuai
		switch err.Tag() {
		case "required":
			// Pesan jika field tidak boleh kosong
			message = fmt.Sprintf("%s harus diisi", err.Field())
		case "email":
			// Pesan jika field harus berisi email yang valid
			message = fmt.Sprintf("%s harus valid email", err.Field())
		case "max":
			// Pesan jika nilai field melebihi batas maksimum
			message = fmt.Sprintf("%s maksimal %s", err.Field(), err.Param())
		case "min":
			// Pesan jika nilai field kurang dari batas minimum
			message = fmt.Sprintf("%s minimal %s", err.Field(), err.Param())
		case "len":
			// Pesan jika panjang karakter tidak sesuai
			message = fmt.Sprintf("%s harus berisi %s karakter", err.Field(), err.Param())
		case "number":
			// Pesan jika nilai field harus berupa angka
			message = fmt.Sprintf("%s harus berupa angka", err.Field())
		default:
			// Pesan default jika tag error tidak teridentifikasi
			message = fmt.Sprintf("%s tidak valid", err.Field())
		}
		// Menambahkan pesan ke daftar error
		errors = append(errors, message)
	}
	// Menggabungkan semua pesan error menjadi satu string, dipisahkan dengan koma
	return strings.Join(errors, ", ")
}
