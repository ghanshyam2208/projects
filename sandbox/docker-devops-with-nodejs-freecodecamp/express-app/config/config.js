module.exports = {
  PORT: process.env.PORT,
  MONGO_USERNAME: process.env.MONGO_USERNAME,
  MONGO_PASSWORD: process.env.MONGO_PASSWORD,
  MONGO_HOST: process.env.MONGO_HOST,
  MONGO_PORT: process.env.MONGO_PORT,
  MONGO_DATABASE: process.env.MONGO_DATABASE,
  REDIS_URL: process.env.REDIS_URL || "my-node-redis",
  REDIS_PORT: process.env.REDIS_PORT || 6379,
  SESSION_SECRET: process.env.SESSION_SECRET || "sample",
};
