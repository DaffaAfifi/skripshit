from application.database import get_connection
from response.response_error import ResponseError

# Fungsi untuk mendapatkan komentar dari berita berdasarkan ID berita
def get_news_comments(id):
    connection = get_connection()
    if connection is None:
        raise ResponseError(500, "Error obtaining database connection")
    cursor = connection.cursor(dictionary=True)
    try:
        query = """
            SELECT 
                news.id, news.gambar, news.judul, news.subjudul, news.isi, news.created_at, 
                comments.user_id AS user_id, comments.comment, comments.created_at AS comment_created_at
            FROM news 
            LEFT JOIN comments ON news.id = comments.news_id 
            WHERE news.id = %s
        """

        cursor.execute(query, (id,))
        rows = cursor.fetchall()

        if not rows:
            raise ResponseError(404, "News not found")
        
        news = rows[0]
        payload = {
            'id': news['id'],
            'gambar': news['gambar'],
            'judul': news['judul'],
            'subjudul': news['subjudul'],
            'isi': news['isi'],
            'created_at': news['created_at'],
            'comments': []
        }

        for row in rows:
            if row['comment']:
                payload['comments'].append({
                    'user_id': row['user_id'],
                    'comment': row['comment'],
                    'comment_created_at': row['comment_created_at']
                })

        return payload
    except ResponseError as e:
        raise e
    except Exception as e:
        raise ResponseError(500, str(e))
    finally:
        if cursor:
            cursor.close()
        if connection:
            connection.close()