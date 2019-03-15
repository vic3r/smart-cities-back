const serialize = require('./serializer');
const writeData = require('./writer');
const isNotEmpty = require('./validator');

const File = require('./File');
const Storage = require('./Storage');

function start (config) {
  if (!isNotEmpty(config)) return 'invalid configuration';
  const {
    host, port, password, filepath, inputfile, outputfile
  } = config;

  const file = new File(filepath, inputfile, outputfile);
  const storage = new Storage(host, port, password);
  return uploadData(file, storage.cli);
}

function uploadData (file, client) {
  serialize(file)
    .then(outputFile => writeData(outputFile, client))
    .then(() => `success`)
    .catch(err => `unexpected error ${err}`);
}

module.exports = start;
