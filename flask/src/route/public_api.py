from flask import Blueprint
from controller.auth_controller import login_controller

public_router = Blueprint('public', __name__)

# login
@public_router.route('/login', methods=['POST'])
def login():
    return login_controller()