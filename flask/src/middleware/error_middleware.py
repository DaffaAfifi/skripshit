from flask import Flask, jsonify
from response.response_error import ResponseError

app = Flask(__name__)

# Error handler untuk menangani exception yang terjadi di aplikasi
@app.errorhandler(Exception)
def error_middleware(error):
    if isinstance(error, ResponseError):
        response = jsonify({"errors": str(error)})
        response.status_code = error.status
    else:
        response = jsonify({"errors": str(error)})
        response.status_code = 500
    return response