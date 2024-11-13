from flask import Flask
from middleware.error_middleware import error_middleware
from response.response_error import ResponseError
from route.api import router
from route.public_api import public_router

def create_app():
    app = Flask(__name__)

    app.register_blueprint(public_router, url_prefix='/api')
    app.register_blueprint(router, url_prefix='/api')

    app.register_error_handler(ResponseError, error_middleware)

    return app