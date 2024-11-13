from flask import request
from response.response_error import ResponseError
from response.response import response
from service.assistance_service import get_assistance_tools, create_assistance_tools

def get_assistance_tools_controller(id):
    try:
        result = get_assistance_tools(id)
        return response(200, result, "Get assistance tools success")
    except ResponseError as e:
        raise e
    
def create_assistance_tools_controller():
    try:
        data = request.get_json()
        result = create_assistance_tools(data)
        return response(200, result, "Create assistance tools success")
    except ResponseError as e:
        raise e