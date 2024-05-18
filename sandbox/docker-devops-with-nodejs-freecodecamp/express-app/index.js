const express = require("express");
const mongoose = require("mongoose");
const {
  MONGO_USERNAME,
  MONGO_PASSWORD,
  MONGO_HOST,
  MONGO_PORT,
  APP_PORT,
} = require("./config/config");

const app = express();
const PORT = APP_PORT || 3000;
const MONGO_URL = `mongodb://${MONGO_USERNAME}:${MONGO_PASSWORD}@${MONGO_HOST}:${MONGO_PORT}/?authSource=admin`;

const connectToMongoRetry = async () => {
  mongoose
    .connect(MONGO_URL)
    .then(() => console.log("Connected to mongodb"))
    .catch(async (e) => {
      console.log("error while connecting to mongodb", e);
      await setTimeout(5000);
      connectToMongoRetry();
    });
};

connectToMongoRetry();

app.get("/", (_req, res) => {
  res.send("<H1>HI There!!</H1>");
});

app.listen(PORT, () => {
  console.log(`listening on port ${PORT}`);
});
