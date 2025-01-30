import express from "express";
import authController from "../controller/auth-controller.js";
import userController from "../controller/user-controller.js";

// Membuat router untuk API publik
const publicRouter = new express.Router();

publicRouter.get("/api/test-connection", userController.test); // Uji koneksi - low complexity
publicRouter.post("/api/login", authController.login); // Login pengguna - medium complexity

export { publicRouter };
