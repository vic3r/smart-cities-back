const fs = require('fs');
const path = require('path');

function writeData (outputfile, client) {
  const outputFilePath = path.resolve(__dirname, '../../', outputfile);
  const jsonFile = fs.readFileSync(outputFilePath.toString());
  const stations = JSON.parse(jsonFile);

  const promises = stations.every(async (station) => {
    const strStation = JSON.stringify(station);
    await client.lpush(station.zone, strStation);
  });

  return Promise.all(promises);
}

module.exports = writeData;
