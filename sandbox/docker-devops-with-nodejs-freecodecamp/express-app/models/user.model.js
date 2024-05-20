const mongoose = require("mongoose");

const userSchema = new mongoose.Schema({
  username: {
    type: String,
    require: [true, "Username is required"],
  },
  password: {
    type: String,
    require: [true, "Password is required"],
  },
});

const User = mongoose.model("User", userSchema);
module.exports = User;
