const cvsToJson = require('convert-csv-to-json');

async function serialize (file) {
  cvsToJson.fieldDelimiter(',');
  await cvsToJson.generateJsonFileFromCsv(file.inputFile, file.outputFile);

  return file.outputFile;
}

module.exports = serialize;
