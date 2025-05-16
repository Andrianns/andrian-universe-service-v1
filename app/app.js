function New() {
  if (process.env.NODE_ENV !== 'production') {
    require('dotenv').config(); //development
  }
  // console.log(process.env.SECRET_KEY);
  const cors = require('cors');
  const express = require('express');

  const app = express();
  const port = process.env.PORT || 3000;

  app.use(cors());
  app.use(express.urlencoded({ extended: true }));
  app.use(express.json());
  app.listen(port, () => {
    console.log(`Example app listening on port ${port}`);
  });

  app.get('/', (req, res) => {
    res.send('Hello World!');
  });
}

module.exports = New();
