from flask import request
from service.auth_service import login, logout
from response.response import response
from response.response_error import ResponseError

# Controller untuk menangani proses login pengguna
def login_controller():
    try:
        data = request.get_json()
        result = login(data)
        return response(200, result, "Login success")
    except ResponseError as e:
        raise e
    
# Controller untuk menangani proses logout pengguna
def logout_controller():
    try:
        token = request.headers.get("Authorization")
        result = logout(token)
        return response(200, result, "Logout success")
    except ResponseError as e:
        raise e