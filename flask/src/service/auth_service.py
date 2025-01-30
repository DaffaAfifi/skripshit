import bcrypt
import jwt
import os
from datetime import datetime, timedelta
from application.database import get_connection
from response.response_error import ResponseError
from validation.user_validation import LoginUserValidation
from validation.validation import validate

# Fungsi untuk login user, memverifikasi kredensial, dan menghasilkan token
def login(req):
    connection = get_connection()
    if connection is None:
        raise ResponseError(500, "Error obtaining database connection")
    cursor = connection.cursor(dictionary=True)
    
    try:
        data = validate(LoginUserValidation, req)

        email = data.email
        password = data.password

        print(email)

        cursor.execute(
            "SELECT id, email, nama, role_id, password FROM users WHERE email = %s", 
            (email,)
        )
        user = cursor.fetchone()

        print(user)

        if not user:
            raise ResponseError(400, "Username or password wrong hehe")

        if not bcrypt.checkpw(password.encode('utf-8'), user['password'].encode('utf-8')):
            raise ResponseError(400, "Username or password wrong haha")

        exp_time = (datetime.now() + timedelta(hours=2)).timestamp()

        token = jwt.encode(
            {
                'id': user['id'],
                'email': user['email'],
                'nama': user['nama'],
                'role': user['role_id'],
                'exp': exp_time
            },
            os.getenv('JWT_SECRET'),
            algorithm='HS256'
        )

        cursor.execute(
            """
            INSERT INTO sessions (token, user_id, expiry, created_at, updated_at)
            VALUES (%s, %s, %s, %s, %s)
            """, 
            (
                token,
                user['id'],
                datetime.utcnow() + timedelta(hours=2),
                datetime.utcnow(),
                datetime.utcnow()
            )
        )
        connection.commit()

        return token
    except ResponseError as e:
        raise e
    except Exception as e:
        raise ResponseError(500, str(e))
    finally:
        if cursor:
            cursor.close()
        if connection:
            connection.close()
    
    
# Fungsi untuk logout user dengan menghapus sesi token
def logout(token):
    connection = get_connection()
    if connection is None:
        raise ResponseError(500, "Error obtaining database connection")
    cursor = connection.cursor(dictionary=True)
    
    try:
        cursor.execute(
            "DELETE FROM sessions WHERE token = %s", 
            (token,)
        )
        connection.commit()

        return "oke"
    except ResponseError as e:
        raise e 
    except Exception as e:
        raise ResponseError(500, str(e))
    finally:
        if cursor:
            cursor.close()
        if connection:
            connection.close()
        