import newsService from "../service/news-service.js";
import { response } from "../response/response.js";

// Controller untuk mendapatkan komentar berita berdasarkan ID
const getNewsCommentsById = async (req, res, next) => {
  try {
    const result = await newsService.getNewsCommentsById(req.params.id);
    response(200, result, "Get news comment success", res);
  } catch (error) {
    next(error);
  }
};

// Controller untuk test, hanya mengembalikan data statis
const test = async (req, res, next) => {
  try {
    res.status(200).json({
      data: "test",
    });
  } catch (error) {
    next(error);
  }
};

export default {
  getNewsCommentsById,
  test,
};
