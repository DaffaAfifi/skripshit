from flask import jsonify

def response(status_code, data, message):
    return jsonify({
        "payload": data,
        "message": message,
        "metadata": {
            "prev": "",
            "next": "",
            "current": "",
        }
    }), status_code