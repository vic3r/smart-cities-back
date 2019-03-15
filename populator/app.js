require('dotenv').load();
const config = require('config');

const start = require('./src');

start(config);
