from flask import request
from response.response import response
from response.response_error import ResponseError
from service.user_service import get_users, get_user_by_id, create_user, get_user_saved_news, get_user_saved_news_comments, get_user_facilities, update_user

# Controller untuk mengambil semua data pengguna
def get_users_controller():
    try:
        page = request.args.get('page', default=1, type=int)
        limit = request.args.get('limit', default=100, type=int)

        result = get_users(page, limit)
        return response(200, result, "Get users success")
    except ResponseError as e:
        raise e
        
# Controller untuk mengambil data pengguna berdasarkan ID
def get_user_by_id_controller(id):
    try:
        result = get_user_by_id(id)
        return response(200, result, "Get user by id success")
    except ResponseError as e:
        raise e

# Controller untuk membuat pengguna baru
def create_user_controller():
    try:
        data = request.get_json()
        result = create_user(data)
        return response(200, result, "Create user success")
    except ResponseError as e:
        raise e
    
# Controller untuk mengambil berita yang disimpan oleh pengguna berdasarkan ID
def get_user_saved_news_controller(id):
    try:
        result = get_user_saved_news(id)
        return response(200, result, "Get user saved news success")
    except ResponseError as e:
        raise e
    
# Controller untuk mengambil komentar berita yang disimpan oleh pengguna berdasarkan ID
def get_user_saved_news_comments_controller(id):
    try:
        result = get_user_saved_news_comments(id)
        return response(200, result, "Get user saved news comments success")
    except ResponseError as e:
        raise e
    
# Controller untuk mengambil fasilitas yang dimiliki pengguna berdasarkan ID
def get_user_facilities_controller(id):
    try:
        result = get_user_facilities(id)
        return response(200, result, "Get user facilities success")
    except ResponseError as e:
        raise e
    
# Controller untuk memperbarui data pengguna berdasarkan ID
def update_user_controller(id):
    try:
        data = request.get_json()
        result = update_user(id, data)
        return response(200, result, "Update user success")
    except ResponseError as e:
        raise e