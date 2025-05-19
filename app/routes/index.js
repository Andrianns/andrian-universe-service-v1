// src/frameworks/webserver/createExpressApp.js
const express = require('express');
const createUserRoutes = require('./user_router');

function createExpressApp(controllers) {
  const app = express();
  app.use(express.json());

  app.use('/user', createUserRoutes(controllers.UserController));

  return app;
}

module.exports = createExpressApp;
