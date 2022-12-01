const fs = require("fs");

const data = fs.readFileSync("./input", "utf8");
const array = data.split("\n");

let max = [];
let sum = 0;
for (let i = 0; i < array.length; i++) {
  const line = array[i];
  if (line == "") {
    if (sum > 0) {
      max.push(sum);
    }
    sum = 0;
  } else {
    sum = sum + Number(line);
  }
}
if (sum > 0) {
  max.push(sum);
}

max.sort((a, b) => a - b);
console.log(max[max.length - 1] + max[max.length - 2] + max[max.length - 3]);
