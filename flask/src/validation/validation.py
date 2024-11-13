from pydantic import ValidationError
from response.response_error import ResponseError

def validate(schema_class, data):
    try:
        return schema_class(**data)
    except ValidationError as e:
        for error in e.errors():
            message = error['msg']
        
        raise ResponseError(400, message=message)