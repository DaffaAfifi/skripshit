import os
import jwt
from functools import wraps
from flask import request, jsonify
from application.database import get_cursor
from response.response_error import ResponseError
from application.logging import logger

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

            cursor = get_cursor()
            query = "SELECT * FROM sessions WHERE token = %s"
            cursor.execute(query, (token,))
            rows = cursor.fetchall()

            if not rows:
                raise ResponseError(401, "Unauthorized")

            request.user = decoded_token
            return f(*args, **kwargs)
        except ResponseError as e:
            raise e
        except Exception as e:
            raise ResponseError(401, str(e))

    return decorated_function