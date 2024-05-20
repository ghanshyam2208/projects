const { Router } = require("express");
const {
  createPostController,
  getAllPosts,
} = require("../controllers/post.controller");
const {
  protectedRoute,
} = require("../helpers/middlewares/protected.middleware");

const router = Router();

router.post("/posts", protectedRoute, createPostController);
router.get("/posts", protectedRoute, getAllPosts);

module.exports = router;
