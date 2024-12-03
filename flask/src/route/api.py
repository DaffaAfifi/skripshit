from flask import Blueprint
from controller.user_controller import get_users_controller, get_user_by_id_controller, create_user_controller, get_user_saved_news_controller, get_user_saved_news_comments_controller, get_user_facilities_controller, update_user_controller
from controller.news_controller import get_news_comments_controller
from controller.assistance_controller import get_assistance_tools_controller, create_assistance_tools_controller
from controller.auth_controller import logout_controller
from middleware.auth_middleware import auth_middleware

router = Blueprint('user', __name__)

# get all users
@router.route('/users', methods=['GET'])
@auth_middleware
def get_users():
    return get_users_controller()

# get user by id
@router.route('/users/<id>', methods=['GET'])
@auth_middleware
def get_user_by_id(id):
    return get_user_by_id_controller(id)

# create user
@router.route('/users', methods=['POST'])
@auth_middleware
def create_user():
    return create_user_controller()

# get user saved news
@router.route('/users/saved-news/<id>', methods=['GET'])
@auth_middleware
def get_user_saved_news(id):
    return get_user_saved_news_controller(id)

# get user saved news comments
@router.route('/users/saved-news/comment/<id>', methods=['GET'])
@auth_middleware
def get_user_saved_news_comment(id):
    return get_user_saved_news_comments_controller(id)

# get user facilities
@router.route('/users/facilities/<id>', methods=['GET'])
@auth_middleware
def get_user_facilities(id):
    return get_user_facilities_controller(id)

# update user
@router.route('/users/<id>', methods=['PUT'])
@auth_middleware
def update_user(id):
    return update_user_controller(id)

# get news comments by news id
@router.route('/news/<id>', methods=['GET'])
@auth_middleware
def get_news_comments(id):
    return get_news_comments_controller(id)

# get assistance tools by assistance id
@router.route('/assistance/<id>', methods=['GET'])
@auth_middleware
def get_assistance_tools(id):
    return get_assistance_tools_controller(id)

# create assistance tools
@router.route('/assistance-tools', methods=['POST'])
@auth_middleware
def create_assistance_tools():
    return create_assistance_tools_controller()

# logout
@router.route('/logout', methods=['POST'])
@auth_middleware
def logout():
    return logout_controller()