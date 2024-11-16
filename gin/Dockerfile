# Menggunakan Go 1.22 Alpine
FROM golang:1.22-alpine

# Menetapkan direktori kerja
WORKDIR /app

# Menyalin go.mod dan go.sum untuk caching dependensi
COPY go.mod go.sum ./

# Mengunduh dependensi
RUN go mod download

# Menyalin seluruh source code ke dalam container
COPY . .

# Membuild binary aplikasi
RUN go build -o app src/main.go

# Mengekspos port aplikasi
EXPOSE 8080

# Perintah untuk menjalankan binary aplikasi
CMD ["./app"]