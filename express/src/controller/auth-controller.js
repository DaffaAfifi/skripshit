import authService from "../service/auth-service.js";
import { response } from "../response/response.js";

// Controller untuk login user
const login = async (req, res, next) => {
  try {
    const result = await authService.login(req.body);
    response(200, result, "Login success", res);
  } catch (error) {
    next(error);
  }
};

// Controller untuk logout user
const logout = async (req, res, next) => {
  try {
    const token = req.get("Authorization");
    const result = await authService.logout(token);
    response(200, result, "Logout success", res);
  } catch (error) {
    next(error);
  }
};

export default {
  login,
  logout,
};
