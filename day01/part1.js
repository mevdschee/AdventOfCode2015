const fs = require('fs')
const input = fs.readFileSync(0).toString();
const output = input.split('').reduce((r,v)=>r+(v=='('?1:(v==')'?-1:0)),0)
console.log(output)
