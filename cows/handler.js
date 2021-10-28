'use strict'

const cows = require('cows');
const all = cows();

module.exports = async (event, context) => {
  let randomCow = Math.floor((Math.random() * all.length) + 1);

  return context
    .status(200)
    .headers({"content-type": "text/plain"})
    .succeed(all[randomCow]+"\n")
}
