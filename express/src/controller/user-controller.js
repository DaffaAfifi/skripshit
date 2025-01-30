import userService from "../service/user-service.js";
import { response } from "../response/response.js";
import jwt from "jsonwebtoken";

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

// Controller untuk memverifikasi token JWT dan mengembalikan decoded data
const testToken = (req, res, next) => {
  try {
    const token = req.get("Authorization");
    const secret = process.env.JWT_SECRET;
    const decoded = jwt.verify(token, secret);
    res.status(200).json({
      data: decoded,
    });
  } catch (error) {
    next(error);
  }
};

// Controller untuk mendapatkan semua pengguna
const getUsers = async (req, res, next) => {
  try {
    const page = parseInt(req.query.page) || 1;
    const limit = parseInt(req.query.limit) || 100;

    const result = await userService.getUsers(page, limit);
    response(200, result, "Get users success", res);
  } catch (error) {
    next(error);
  }
};

// Controller untuk membuat pengguna baru
const createUser = async (req, res, next) => {
  try {
    const result = await userService.createUser(req.body);
    response(201, result, "Create user success", res);
  } catch (error) {
    next(error);
  }
};

// Controller untuk mendapatkan pengguna berdasarkan ID
const getUserById = async (req, res, next) => {
  try {
    const result = await userService.getUserById(req.params.id);
    response(200, result, "Get user success", res);
  } catch (error) {
    next(error);
  }
};

// Controller untuk mendapatkan berita yang disimpan oleh pengguna
const getSavedNews = async (req, res, next) => {
  try {
    const result = await userService.getSavedNews(req.params.id);
    response(200, result, "Get saved news success", res);
  } catch (error) {
    next(error);
  }
};

// Controller untuk mendapatkan fasilitas yang dimiliki pengguna
const getFacilities = async (req, res, next) => {
  try {
    const result = await userService.getFacilities(req.params.id);
    response(200, result, "Get facilities success", res);
  } catch (error) {
    next(error);
  }
};

// Controller untuk memperbarui data pengguna berdasarkan ID
const updateUser = async (req, res, next) => {
  try {
    const result = await userService.updateUser(req.params.id, req.body);
    response(200, result, "Update user success", res);
  } catch (error) {
    next(error);
  }
};

// Controller untuk mendapatkan komentar berita yang disimpan oleh pengguna
const getSavedNewsComment = async (req, res, next) => {
  try {
    const result = await userService.getSavedNewsComment(req.params.id);
    response(200, result, "Get saved news comment success", res);
  } catch (error) {
    next(error);
  }
};

export default {
  test,
  getUsers,
  createUser,
  getUserById,
  getSavedNews,
  getFacilities,
  updateUser,
  getSavedNewsComment,
  testToken,
};
