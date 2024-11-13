from loguru import logger
import sys

logger.remove()
logger.add(sys.stderr, level="INFO", format="{time} {level} {message}")

if __name__ == "__main__":
    logger.info("Logger is ready")