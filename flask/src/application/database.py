import os
import mysql.connector
from mysql.connector import pooling
from dotenv import load_dotenv

load_dotenv()

POOL_SIZE = 32

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

db_pool = create_pool()

db_connection = db_pool.get_connection()

def get_cursor():
    return db_connection.cursor(dictionary=True)