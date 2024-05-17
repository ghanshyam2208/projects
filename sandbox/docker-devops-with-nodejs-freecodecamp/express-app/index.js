// const express = require("express");
import express from "express";
import mongoose, { mongo } from "mongoose";

const app = express();
const PORT = process.env.PORT || 3000;

mongoose
  .connect("mongodb://root:root@my-node-app-mongo:27017/?authSource=admin")
  .then(() => console.log("Connected to mongodb"))
  .catch((e) => console.log("error while connecting to mongodb", e));

app.get("/", (_req, res) => {
  //   res.json({
  //     success: true,
  //     message: "successful response",
  //   });
  //   return res.status(200).json({
  //     success: true,
  //     message: "successful response",
  //   });

  //   res.status(200).send();

  res.send("<H1>HI There!!</H1>");
});

app.listen(PORT, () => {
  console.log(`listening on port ${PORT}`);
});
