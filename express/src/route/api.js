import express from "express";
import { authMiddleware } from "../middleware/auth-middleware.js";
import authController from "../controller/auth-controller.js";
import userController from "../controller/user-controller.js";
import newsController from "../controller/news-controller.js";
import assistanceController from "../controller/assistance-controller.js";

// Membuat router untuk API user
const userRouter = new express.Router();

// Menambahkan middleware autentikasi pada setiap route di bawah ini
userRouter.use(authMiddleware);

// Route untuk test token
userRouter.get("/api/test-token", userController.testToken);

// Users API
userRouter.get("/api/users", userController.getUsers); // Mendapatkan semua pengguna - low complexity
userRouter.post("/api/users", userController.createUser); // Membuat pengguna baru - low complexity
userRouter.put("/api/users/:id", userController.updateUser); // Memperbarui data pengguna - low complexity
userRouter.get("/api/users/:id", userController.getUserById); // Mendapatkan data pengguna berdasarkan ID - low complexity
userRouter.get("/api/users/saved-news/:id", userController.getSavedNews); // Mendapatkan berita yang disimpan oleh pengguna - medium complexity
userRouter.get("/api/users/facilities/:id", userController.getFacilities); // Mendapatkan fasilitas yang dimiliki pengguna (sertifikat, pelatihan, bantuan) - high complexity
userRouter.get(
  "/api/users/saved-news/comment/:id",
  userController.getSavedNewsComment
); // Mendapatkan komentar dari berita yang disimpan pengguna - high complexity

// News API
userRouter.get("/api/news/:id", newsController.getNewsCommentsById); // Mendapatkan berita dan komentar berdasarkan ID - medium complexity

// Assistance API
userRouter.get("/api/assistance/:id", assistanceController.getAssistanceById); // Mendapatkan bantuan & alat berdasarkan ID - medium complexity
userRouter.post(
  "/api/assistance-tools",
  assistanceController.createAssistanceTools
); // Membuat alat bantuan yang memicu pembaruan total harga - high complexity

// Logout
userRouter.post("/api/logout", authController.logout); // Logout pengguna - low complexity

export { userRouter };
