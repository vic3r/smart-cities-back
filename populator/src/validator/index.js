
function isNotEmpty (config) {
  const values = Object.values(config);
  return values.every(existsValue);
}

function existsValue (value) {
  return (
    (typeof value !== 'undefined') &&
    (value !== null) &&
    (value !== false) &&
    (value.length !== 0) &&
    (value !== '')
  );
}

module.exports = isNotEmpty;
