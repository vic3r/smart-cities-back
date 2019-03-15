const path = require('path');

class CsvFile {
  constructor (filepath = '', inputfile, outputfile) {
    this.file = path.join('./', filepath, inputfile);
    this.outputfile = path.join('./', filepath, outputfile);
  }

  get inputFile () {
    return this.file;
  }

  get outputFile () {
    return this.outputfile;
  }
}

module.exports = CsvFile;
