from pydantic import ValidationError
from response.response_error import ResponseError

# Fungsi untuk melakukan validasi data menggunakan schema_class
def validate(schema_class, data):
    try:
        # Mencoba untuk memvalidasi data sesuai dengan schema_class
        return schema_class(**data)
    except ValidationError as e:
        # Jika terjadi error validasi, mengambil pesan error dari setiap masalah yang ditemukan
        for error in e.errors():
            message = error['msg']
        
        # Menaikkan ResponseError jika terjadi error validasi
        raise ResponseError(400, message=message)