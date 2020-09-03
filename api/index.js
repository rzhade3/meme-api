const { readFileSync } = require('fs');
const { join } = require('path');
const filepath = join(__dirname, '..', 'db', 'script.txt');
const file = readFileSync(filepath, 'utf8');

module.exports = (req, res) => {
  res.send(file)
}
