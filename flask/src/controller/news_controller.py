from response.response import response
from response.response_error import ResponseError
from service.news_service import get_news_comments

def get_news_comments_controller(id):
    try:
        result = get_news_comments(id)
        return response(200, result, "Get news comments success")
    except ResponseError as e:
        raise e