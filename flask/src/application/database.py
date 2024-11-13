import os
import mysql.connector
from dotenv import load_dotenv

load_dotenv()

def create_connection():
    connection = None
    try:
        connection = mysql.connector.connect(
            host=os.getenv("DB_HOST"),
            port=os.getenv("DB_PORT"),
            user=os.getenv("DB_USER"),
            password=os.getenv("DB_PASSWORD"),
            database=os.getenv("DB_NAME")
        )
        print("Connection to MySQL DB successful")
    except mysql.connector.Error as err:
        print(f"Error connecting to MySQL DB: {err}")

    return connection

db_connection = create_connection()

def get_cursor():
    return db_connection.cursor(dictionary=True)