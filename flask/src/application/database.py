import os
import mysql.connector
from mysql.connector import pooling
from dotenv import load_dotenv

load_dotenv()

POOL_SIZE = 32

# Fungsi untuk membuat pool koneksi MySQL
def create_pool():
    try:
        pool = pooling.MySQLConnectionPool(
            pool_name="my_pool",
            pool_size=POOL_SIZE,
            host=os.getenv("DB_HOST"),
            port=os.getenv("DB_PORT"),
            user=os.getenv("DB_USER"),
            password=os.getenv("DB_PASSWORD"),
            database=os.getenv("DB_NAME")
        )
        print("MySQL Connection Pool created successfully")
        return pool
    except mysql.connector.Error as err:
        print(f"Error creating connection pool: {err}")
        return None

# Membuat pool koneksi dan menyimpannya ke variabel db_pool
db_pool = create_pool()

# Fungsi untuk mendapatkan koneksi baru dari pool
def get_connection():
    try:
        if db_pool is None:
            raise Exception("Connection pool not initialized")
        connection = db_pool.get_connection()
        if connection.is_connected():
            return connection
        else:
            raise Exception("Connection is not valid")
    except Exception as e:
        print(f"Error getting connection from pool: {e}")
        return None