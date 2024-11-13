class ResponseError(Exception):
    def __init__(self, status, message):
        super().__init__(message)
        self.status = status

if __name__ == "__main__":
    error = ResponseError(404, "Not found")
    print(f"Error: {error}, Status: {error.status}")