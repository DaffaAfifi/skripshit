from flask import jsonify

# Fungsi untuk membentuk response API yang konsisten dengan status, data, dan pesan
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