// src/app/config/index.js
require('dotenv').config();

const config = {
  appName: process.env.SERVICE_NAME || 'andrian-universe',
  env: process.env.NODE_ENV || 'development',
  port: process.env.PORT || 3000,

  db: {
    development: {
      username: 'postgres',
      password: 'postgres',
      database: 'andrian_universe_service',
      host: '127.0.0.1',
      dialect: 'postgres',
    },
    test: {
      username: process.env.DB_USERNAME,
      password: process.env.DB_PASSWORD,
      database: process.env.DB_NAME,
      host: process.env.DB_HOST,
      dialect: process.env.DB_DIALECT || 'postgres',
      dialectOptions: {
        ssl: {
          require: true,
          rejectUnauthorized: false,
        },
      },
    },
    production: {
      username: process.env.DB_USERNAME,
      password: process.env.DB_PASSWORD,
      database: process.env.DB_NAME,
      host: process.env.DB_HOST,
      dialect: process.env.DB_DIALECT || 'postgres',
      dialectOptions: {
        ssl: {
          require: true,
          rejectUnauthorized: false,
        },
      },
    },
  },

  services: {
    inquiry: {
      baseURL: process.env.INQUIRY_SERVICE_URL,
    },
  },
};

module.exports = config;
