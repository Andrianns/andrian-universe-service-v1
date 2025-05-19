const { NewUserRepository } = require('./repository/user_repository');
const { NewUserController } = require('./controllers/user_controller');
const models = require('./models/index');
const createExpressApp = require('./routes/index');
const express = require('express');
const cors = require('cors');
const port = process.env.PORT || 3000;

function New() {
  if (process.env.NODE_ENV !== 'production') {
    require('dotenv').config(); //development
  }

  const UserRepository = NewUserRepository(models);
  const UserController = NewUserController(null, UserRepository);

  const app = createExpressApp({
    UserController,
  });
  app.use(cors());
  app.use(express.urlencoded({ extended: true }));
  app.use(express.json());

  app.listen(port, () => {
    console.log(`Example app listening on port ${port}`);
  });
  app.get('/', (req, res) => {
    res.send('Hello World!');
    console.log(UserRepository);
  });
}

module.exports = New();
