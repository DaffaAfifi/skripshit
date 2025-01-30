import os
import jwt
from functools import wraps
from flask import request, jsonify
from application.database import get_connection
from response.response_error import ResponseError
from application.logging import logger

# Middleware untuk autentikasi
def auth_middleware(f):
    @wraps(f)
    def decorated_function(*args, **kwargs):
        token = request.headers.get("Authorization")
        logger.info(token)

        if not token:
            raise ResponseError(401, "Unauthorized")

        try:
            jwt_secret = os.getenv("JWT_SECRET")
            decoded_token = jwt.decode(token, jwt_secret, algorithms=["HS256"])

            connection = get_connection()
            if connection is None:
                raise ResponseError(500, "Error obtaining database connection")
            
            cursor = connection.cursor(dictionary=True)
            query = "SELECT * FROM sessions WHERE token = %s"
            cursor.execute(query, (token,))
            session = cursor.fetchone()

            if not session:
                raise ResponseError(401, "Unauthorized")
            
            request.user = decoded_token
            return f(*args, **kwargs)
        except jwt.ExpiredSignatureError:
            raise ResponseError(401, "Token has expired")
        except jwt.InvalidTokenError:
            raise ResponseError(401, "Invalid token")
        except ResponseError as e:
            raise e
        except Exception as e:
            raise ResponseError(500, str(e))
        finally:
            if cursor:
                cursor.close()
            if connection:
                connection.close()
    return decorated_function