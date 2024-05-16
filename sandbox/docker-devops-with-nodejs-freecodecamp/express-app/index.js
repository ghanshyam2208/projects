// const express = require("express");
import express from "express";
const app = express();
const PORT = process.env.PORT || 3000;

app.get("/", (_req, res) => {
  //   res.json({
  //     success: true,
  //     message: "successful response",
  //   });
  //   return res.status(200).json({
  //     success: true,
  //     message: "successful response",
  //   });

  res.status(200).send();
});

app.listen(PORT, () => {
  console.log(`listening on port ${PORT}`);
});
