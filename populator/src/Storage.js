const redis = require('redis');

class Storage {
  constructor (host, port, password) {
    this.client = redis.createClient({
      port,
      host,
      password
    });
  }

  get cli () {
    return this.client;
  }
}

module.exports = Storage;
