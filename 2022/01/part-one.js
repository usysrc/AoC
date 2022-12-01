const fs = require("fs");

const data = fs.readFileSync("./input", "utf8");
const array = data.split("\n");

let max = 0;
let sum = 0;
for (let index = 0; index < array.length; index++) {
  const line = array[index];
  if (line == "") {
    if (sum > max) {
      max = sum;
    }
    sum = 0;
  } else {
    sum = sum + Number(line);
  }
}
if (sum > max) {
  max = sum;
}

console.log(max);
