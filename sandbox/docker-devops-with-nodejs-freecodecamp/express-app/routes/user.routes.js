const { Router } = require("express");
const { signUp, login } = require("../controllers/user.controller");
const router = Router();

router.post("/signup", signUp);
router.post("/login", login);

module.exports = router;
