const Post = require("../models/post.model");

const getAllPosts = async (req, res) => {
  console.log("getting posts");
  const posts = await Post.find();
  res.json({
    success: true,
    posts: posts,
  });
};

const createPostController = async (req, res) => {
  try {
    const post = await Post.create(req.body);
    res.json({
      success: true,
      posts: post,
    });
  } catch (error) {
    console.log(error);
  }
};

module.exports = { createPostController, getAllPosts };
