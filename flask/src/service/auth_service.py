import bcrypt
import jwt
import os
from datetime import datetime, timedelta
from application.database import get_cursor,db_connection
from response.response_error import ResponseError
from validation.user_validation import LoginUserValidation
from validation.validation import validate

def login(req):
    cursor = get_cursor()
    try:
        data = validate(LoginUserValidation, req)

        email = data.email
        password = data.password

        cursor.execute(
            "SELECT id, email, nama, role_id, password FROM users WHERE email = %s", 
            (email,)
        )
        user = cursor.fetchone()

        if not user:
            raise ResponseError(400, "Username or password wrong")

        if not bcrypt.checkpw(password.encode('utf-8'), user['password'].encode('utf-8')):
            raise ResponseError(400, "Username or password wrong")

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
            INSERT INTO sessions (token, email, expiry, created_at, updated_at)
            VALUES (%s, %s, %s, %s, %s)
            """, 
            (
                token,
                user['email'],
                datetime.utcnow() + timedelta(hours=2),
                datetime.utcnow(),
                datetime.utcnow()
            )
        )
        db_connection.commit()

        return token
    except ResponseError as e:
        raise e
    except Exception as e:
        raise ResponseError(500, str(e))
    
def logout(token):
    cursor = get_cursor()
    try:
        cursor.execute(
            "DELETE FROM sessions WHERE token = %s", 
            (token,)
        )
        db_connection.commit()

        return "oke"
    except ResponseError as e:
        raise e 
    except Exception as e:
        raise ResponseError(500, str(e))