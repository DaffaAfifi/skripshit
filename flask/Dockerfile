# Menggunakan Python 3.11 sebagai base image
FROM python:3.11-alpine

# Menetapkan direktori kerja
WORKDIR /app

# Menyalin file requirements.txt untuk mengelola dependensi
COPY requirements.txt .

# Menginstal dependensi aplikasi
RUN pip install --no-cache-dir -r requirements.txt

# Menyalin seluruh source code dari folder src ke dalam folder /app di container
COPY src/ /app/

# Mengekspos port Flask (default 5000)
EXPOSE 5000

# Menentukan perintah untuk menjalankan aplikasi
CMD ["python", "/app/main.py"]