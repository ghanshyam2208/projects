const crypto = require("crypto");
const User = require("../models/user.model");

const signUp = async (req, res) => {
  const algorithm = "aes256";
  const key = "password";
  const cipher = crypto.createCipher(algorithm, key);

  req.body.password =
    cipher.update(req.body.password, "utf8", "hex") + cipher.final("hex");

  const user = await User.create(req.body);

  req.session.user = user;
  res.json({
    success: true,
    users: user,
  });
};

const login = async (req, res, next) => {
  try {
    const algorithm = "aes256";
    const key = "password";
    const decipher = crypto.createDecipher(algorithm, key);

    const user = await User.findOne({ username: req.body.username });
    if (!user) {
      throw new Error(`User not found with ${req.body.username}`);
    }

    let decrypted =
      decipher.update(user.password, "hex", "utf8") + decipher.final("utf8");

    console.log(decrypted);

    if (decrypted !== req.body.password) {
      throw new Error(`Password did not match!`);
    }

    req.session.user = user;

    res.json({
      success: true,
      msg: "Successfully logged in!",
    });
  } catch (e) {
    next(e);
  }
};

module.exports = {
  signUp,
  login,
};
