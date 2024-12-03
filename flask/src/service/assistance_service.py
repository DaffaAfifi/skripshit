from response.response_error import ResponseError
from validation.validation import validate
from validation.assistance_validation import CreateAssistanceToolsValidation
from application.database import get_cursor, db_connection

def get_assistance_tools(id):
    cursor = get_cursor()
    if cursor is None:
        raise ResponseError(500, "Error obtaining database connection")
    
    try:
        query = """
            SELECT 
                assistance.id, assistance.nama, assistance.koordinator, 
                assistance.sumber_anggaran, assistance.total_anggaran, 
                assistance.tahun_pemberian, 
                assistance_tools.kuantitas, 
                tools.nama_item, tools.harga, tools.deskripsi
            FROM assistance
            LEFT JOIN assistance_tools ON assistance.id = assistance_tools.assistance_id
            LEFT JOIN tools ON assistance_tools.tools_id = tools.id
            WHERE assistance.id = %s
        """

        cursor.execute(query, (id,))
        rows = cursor.fetchall()

        if not rows:
            raise ResponseError(404, "Assistance not found")
        
        assistance = rows[0]
        payload = {
            'id': assistance['id'],
            'nama': assistance['nama'],
            'koordinator': assistance['koordinator'],
            'sumber_anggaran': assistance['sumber_anggaran'],
            'total_anggaran': assistance['total_anggaran'],
            'tahun_pemberian': assistance['tahun_pemberian'],
            'tools': []
        }

        for row in rows:
            if row['nama_item']:
                payload['tools'].append({
                    'kuantitas': row['kuantitas'],
                    'nama_item': row['nama_item'],
                    'harga': row['harga'],
                    'deskripsi': row['deskripsi']
                })

        return payload
    except ResponseError as e:
        raise e
    except Exception as e:
        raise ResponseError(500, str(e))
    finally:
        if cursor:
            cursor.close()
    
def create_assistance_tools(req):
    cursor = get_cursor()
    if cursor is None:
        raise ResponseError(500, "Error obtaining database connection")
    
    try:
        data = validate(CreateAssistanceToolsValidation, req)

        assistance_id = data.assistance_id
        tools_id = data.tools_id
        kuantitas = data.kuantitas

        query = """
            INSERT INTO assistance_tools (assistance_id, tools_id, kuantitas)
            VALUES (%s, %s, %s)
        """

        cursor.execute(query, (assistance_id, tools_id, kuantitas))
        db_connection.commit()

        return "oke"
    except ResponseError as e:
        raise e
    except Exception as e:
        raise ResponseError(500, str(e))
    finally:
        if cursor:
            cursor.close()