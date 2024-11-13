from application.database import db_connection, get_cursor
from response.response_error import ResponseError
import bcrypt
from validation.validation import validate
from validation.user_validation import CreateUserValidation, UpdateUserValidation
from collections import defaultdict

def get_users():
    cursor = get_cursor()
    try:
        cursor.execute("SELECT nama, email, NIK, alamat, telepon, jenis_kelamin, kepala_keluarga, tempat_lahir, tanggal_lahir, jenis_usaha FROM users")
        users = cursor.fetchall()
        return users
    except Exception as e:
        raise ResponseError(500, str(e))
    
def get_user_by_id(id):
    cursor = get_cursor()
    try:
        cursor.execute("SELECT nama, email, NIK, alamat, telepon, jenis_kelamin, kepala_keluarga, tempat_lahir, tanggal_lahir,jenis_usaha FROM users WHERE id = %s", (id,))
        user = cursor.fetchone()

        if not user:
            raise ResponseError(404, "User not found")

        return user
    except ResponseError as e:
        raise e
    except Exception as e:
        raise ResponseError(500, str(e))
    
def create_user(req):
    cursor = get_cursor()
    try:
        data = validate(CreateUserValidation, req)

        nama = data.nama
        email = data.email
        password = bcrypt.hashpw(data.password.encode('utf-8'), bcrypt.gensalt())
        NIK = data.NIK
        alamat = data.alamat
        telepon = data.telepon
        jenis_kelamin = data.jenis_kelamin
        kepala_keluarga = data.kepala_keluarga
        tempat_lahir = data.tempat_lahir
        tanggal_lahir = data.tanggal_lahir
        jenis_usaha = data.jenis_usaha

        query = """
            INSERT INTO users (nama, email, password, NIK, alamat, telepon, jenis_kelamin, kepala_keluarga, tempat_lahir, tanggal_lahir, jenis_usaha)
            VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)
            ON DUPLICATE KEY UPDATE email = VALUES(email)
        """

        cursor.execute(query, (nama, email, password, NIK, alamat, telepon, jenis_kelamin, kepala_keluarga, tempat_lahir, tanggal_lahir, jenis_usaha))
        db_connection.commit()

        if cursor.rowcount == 0:
            raise ResponseError(400, "Email or NIK already exists")
        
        return {"user_id": cursor.lastrowid}
    except ResponseError as e:
        raise e
    except Exception as e:
        raise ResponseError(500, str(e))

def get_user_saved_news(id):
    cursor = get_cursor()
    try:
        query = """
            SELECT 
                users.id, users.nama, users.email, 
                news.id AS news_id, news.gambar, news.judul, news.subjudul, news.isi, 
                DATE_FORMAT(news.created_at, '%d/%m/%Y') AS formatted_created_at
            FROM users 
            LEFT JOIN saved_news ON users.id = saved_news.user_id 
            LEFT JOIN news ON saved_news.news_id = news.id 
            WHERE users.id = %s
        """

        cursor.execute(query, (id,))
        rows = cursor.fetchall()

        if not rows:
            raise ResponseError(404, "User not found")
        
        payload = {
            'id': rows[0]['id'],
            'nama': rows[0]['nama'],
            'email': rows[0]['email'],
            'berita_tersimpan': [
                {
                    'id': row['news_id'],
                    'gambar': row['gambar'],
                    'judul': row['judul'],
                    'subjudul': row['subjudul'],
                    'isi': row['isi'],
                    'created_at': row['formatted_created_at']
                }
                for row in rows if row['news_id'] is not None
            ]
        }

        return payload
    except ResponseError as e:
        raise e
    except Exception as e:
        raise ResponseError(500, str(e))

def get_user_saved_news_comments(id):
    cursor = get_cursor()
    try:
        query = """
            SELECT 
                users.id, users.nama, users.email, news.id AS news_id, news.gambar, news.judul, news.subjudul, news.isi, DATE_FORMAT(news.created_at, '%d/%m/%Y') AS formatted_created_at,
                comments.news_id AS comment_id, comments.comment, DATE_FORMAT(comments.created_at, '%d/%m/%Y') AS comment_created_at
            FROM users
            LEFT JOIN saved_news ON users.id = saved_news.user_id
            LEFT JOIN news ON saved_news.news_id = news.id
            LEFT JOIN comments ON news.id = comments.news_id
            WHERE users.id = %s
        """

        cursor.execute(query, (id,))
        rows = cursor.fetchall()

        if not rows:
            raise ResponseError(404, "User not found")
        
        payload = {
                'id': rows[0]['id'],
                'nama': rows[0]['nama'],
                'email': rows[0]['email'],
                'news': []
            }
        
        news_map = {}

        for row in rows:
            if row['news_id'] not in news_map:
                news_map[row['news_id']] = {
                    'id': row['news_id'],
                    'gambar': row['gambar'],
                    'judul': row['judul'],
                    'subjudul': row['subjudul'],
                    'isi': row['isi'],
                    'created_at': row['formatted_created_at'],
                    'comment': []
                }
            if row['comment_id']:
                news_map[row['news_id']]['comment'].append({
                    'comment': row['comment'],
                    'created_at': row['comment_created_at']
                })

        payload['news'] = list(news_map.values())

        return payload
    except ResponseError as e:
        raise e
    except Exception as e:
        raise ResponseError(500, str(e))
    
def get_user_facilities(id):
    cursor = get_cursor()
    try:
        query = """
            SELECT
                users.id, users.nama, users.email, 
                sertificates.id AS id_sertifikat, sertificates.nama AS nama_sertifikat, user_sertificates.no_sertifikat, 
                sertificates.tanggal_terbit, sertificates.kadaluarsa, sertificates.keterangan,
                trainings.id AS id_pelatihan, trainings.nama AS nama_pelatihan, trainings.penyelenggara, 
                trainings.tanggal_pelaksanaan, trainings.tempat,
                assistance.id AS id_bantuan, assistance.nama AS nama_bantuan, assistance.koordinator, 
                assistance.sumber_anggaran, assistance.total_anggaran, assistance.tahun_pemberian,
                assistance_tools.kuantitas,
                tools.id AS id_alat, tools.nama_item, tools.harga, tools.deskripsi
            FROM users
            LEFT JOIN user_sertificates ON users.id = user_sertificates.user_id
            LEFT JOIN sertificates ON user_sertificates.sertificates_id = sertificates.id
            LEFT JOIN user_trainings ON users.id = user_trainings.user_id
            LEFT JOIN trainings ON user_trainings.trainings_id = trainings.id
            LEFT JOIN assistance ON users.id = assistance.user_id
            LEFT JOIN assistance_tools ON assistance.id = assistance_tools.assistance_id
            LEFT JOIN tools ON assistance_tools.tools_id = tools.id
            WHERE users.id = %s
        """

        cursor.execute(query, (id,))
        rows = cursor.fetchall()

        if not rows:
            raise ResponseError(404, "User not found")
        
        payload = {
            'id': rows[0]['id'],
            'nama': rows[0]['nama'],
            'email': rows[0]['email'],
            'sertifikat': [],
            'pelatihan': [],
            'bantuan': []
        }

        sertifikat_map = {}
        pelatihan_map = {}
        bantuan_map = defaultdict(lambda: {"alat": []})

        for row in rows:
            if row['id_sertifikat'] and row['id_sertifikat'] not in sertifikat_map:
                sertifikat_map[row['id_sertifikat']] = {
                    'id': row['id_sertifikat'],
                    'nama': row['nama_sertifikat'],
                    'no_sertifikat': row['no_sertifikat'],
                    'tanggal_terbit': row['tanggal_terbit'],
                    'kadaluarsa': row['kadaluarsa'],
                    'keterangan': row['keterangan']
                }

            if row['id_pelatihan'] and row['id_pelatihan'] not in pelatihan_map:
                pelatihan_map[row['id_pelatihan']] = {
                    'id': row['id_pelatihan'],
                    'nama': row['nama_pelatihan'],
                    'penyelenggara': row['penyelenggara'],
                    'tanggal_pelaksanaan': row['tanggal_pelaksanaan'],
                    'tempat': row['tempat']
                }

            if row['id_bantuan']:
                bantuan_item = bantuan_map[row['id_bantuan']]
                bantuan_item.update({
                    'id': row['id_bantuan'],
                    'nama': row['nama_bantuan'],
                    'koordinator': row['koordinator'],
                    'sumber_anggaran': row['sumber_anggaran'],
                    'total_anggaran': row['total_anggaran'],
                    'tahun_pemberian': row['tahun_pemberian']
                })

                if row['id_alat'] and not any(tool['id'] == row['id_alat'] for tool in bantuan_item['alat']):
                    bantuan_item['alat'].append({
                        'id': row['id_alat'],
                        'nama': row['nama_item'],
                        'harga': row['harga'],
                        'kuantitas': row['kuantitas']
                    })

        payload['sertifikat'] = list(sertifikat_map.values())
        payload['pelatihan'] = list(pelatihan_map.values())
        payload['bantuan'] = list(bantuan_map.values())

        return payload
    except ResponseError as e:
        raise e
    except Exception as e:
        raise ResponseError(500, str(e))
    
def update_user(id, req):
    cursor = get_cursor()
    try:
        data = validate(UpdateUserValidation, req)
        data_dict = data.dict(exclude_unset=True)

        if not data_dict:
            raise ResponseError(400, "No valid fields to update")

        updates = ", ".join(f"{key} = %s" for key in data_dict)
        values = list(data_dict.values())
        values.append(id)

        query = f"UPDATE users SET {updates} WHERE id = %s"
        cursor.execute(query, values)
        db_connection.commit()

        if cursor.rowcount == 0:
            raise ResponseError(404, "User not found")

        return {"affected_rows": cursor.rowcount}

    except ResponseError as e:
        raise e
    except Exception as e:
        raise ResponseError(500, str(e))