// src/routes/user_router.js
const express = require('express');

module.exports = function (userController) {
  const router = express.Router();
  router.get('/', userController.getAccounts.bind(userController));

  return router;
};
