from flask import request
from service.auth_service import login, logout
from response.response import response
from response.response_error import ResponseError

def login_controller():
    try:
        data = request.get_json()
        result = login(data)
        return response(200, result, "Login success")
    except ResponseError as e:
        raise e
    
def logout_controller():
    try:
        token = request.headers.get("Authorization")
        result = logout(token)
        return response(200, result, "Logout success")
    except ResponseError as e:
        raise e
    