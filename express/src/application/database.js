import "dotenv/config";
import mysql from "mysql2";

// Membuat pool koneksi database dengan konfigurasi dari file .env
const pool = mysql.createPool({
  host: process.env.DB_HOST,
  port: process.env.DB_PORT,
  user: process.env.DB_USER,
  password: process.env.DB_PASSWORD,
  database: process.env.DB_NAME,
  waitForConnections: true,
  connectionLimit: 50,
  queueLimit: 0,
});

// Mengekspor pool agar dapat digunakan di file lain
export default pool;
