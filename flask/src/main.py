from application.web import create_app
from application.logging import logger

app = create_app()

port = 5000
if __name__ == "__main__":
    app.run(port=port, debug=True)
    logger.info(f"Server started on port {port}")