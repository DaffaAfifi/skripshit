import assistanceService from "../service/assistance-service.js";
import { response } from "../response/response.js";

// Controller untuk mengambil bantuan berdasarkan ID
const getAssistanceById = async (req, res, next) => {
  try {
    const result = await assistanceService.getAssistanceById(req.params.id);
    response(200, result, "Get assistance success", res);
  } catch (error) {
    next(error);
  }
};

// Controller untuk membuat alat bantuan baru
const createAssistanceTools = async (req, res, next) => {
  try {
    const result = await assistanceService.createAssistanceTools(req.body);
    response(201, result, "Create assistance tools success", res);
  } catch (error) {
    next(error);
  }
};

export default { getAssistanceById, createAssistanceTools };
