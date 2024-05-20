const express = require("express");
const mongoose = require("mongoose");
const session = require("express-session");
const redis = require("redis");
const RedisStore = require("connect-redis").default;

const {
  MONGO_USERNAME,
  MONGO_PASSWORD,
  MONGO_HOST,
  MONGO_PORT,
  PORT,
  MONGO_DATABASE,
  REDIS_URL,
  REDIS_PORT,
  SESSION_SECRET,
} = require("./config/config");

const app = express();

const MONGO_URL = `mongodb://${MONGO_USERNAME}:${MONGO_PASSWORD}@${MONGO_HOST}:${MONGO_PORT}/${MONGO_DATABASE}?authSource=admin`;

const connectToMongoRetry = () => {
  mongoose
    .connect(MONGO_URL)
    .then(() => console.log("Connected to MongoDB"))
    .catch((e) => {
      console.log("Error while connecting to MongoDB", e);
      setTimeout(connectToMongoRetry, 5000); // Retry after 5 seconds
    });
};

const redisClient = redis.createClient({
  url: `redis://${REDIS_URL}:${REDIS_PORT}`,
});

redisClient.on("connect", function () {
  console.log("Connected to RedisServer");
});
redisClient.on("error", function (error) {
  console.log("redis client error", error);
});
redisClient.connect().catch(console.error);

// Initialize store.
let redisStore = new RedisStore({
  client: redisClient,
  // prefix: "myapp:",
});

app.use(
  session({
    store: new RedisStore({ client: redisClient }),
    secret: SESSION_SECRET,
    resave: false,
    cookie: {
      secure: false,
      httpOnly: true,
      maxAge: 30000,
    },
  })
);

// connectToRedisRetry();
connectToMongoRetry();

app.use(express.json());

app.use("/", require("./routes/post.route"));
app.use("/", require("./routes/user.routes"));

app.get("/", (_req, res) => {
  res.send("<H1>HI There!!</H1>");
});

// Common/global error handling
app.use((err, req, res, next) => {
  console.error("Error:", err);
  const statusCode = err.status || 500;
  const errorMessage = err.message || "Internal server error";
  res.status(statusCode).json({
    success: false,
    message: errorMessage,
  });
});

app.listen(PORT, () => {
  console.log(`Listening on port ${PORT}`);
});
