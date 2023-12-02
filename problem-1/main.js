const fs = require('fs');

const input = fs.readFileSync('./input.txt').toString();
const parts = input.split('\n');

const nameToNumber = {
    'zero': 0,
    'one': 1,
    'two': 2,
    'three': 3,
    'four': 4,
    'five': 5,
    'six': 6,
    'seven': 7,
    'eight': 8,
    'nine': 9,
    '0': 0,
    '1': 1,
    '2': 2,
    '3': 3,
    '4': 4,
    '5': 5,
    '6': 6,
    '7': 7,
    '8': 8,
    '9': 9
};

const seekSymbols = Object.keys(nameToNumber);

let sum = 0;

for (const row of parts) {
    let first = null;
    let last = null;
    
    let firstIndex = row.length;
    let lastIndex = 0;

    for (const symbol of seekSymbols) {
        const f = row.indexOf(symbol);
        const l = row.lastIndexOf(symbol);
        if (f > -1 && f < firstIndex) {
            firstIndex = f;
            first = nameToNumber[symbol];
        }

        if (l >= lastIndex) {
            lastIndex = l;
            last = nameToNumber[symbol];
        }
    }

    if (first === null || last === null) {
        console.log(row);
    }
    sum += first*10 + last;
}

console.log(sum);