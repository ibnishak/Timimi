#!/usr/bin/env node

var fs = require('fs');
var readline = require('readline');

var nativeMessage = require('./index');

var output = new nativeMessage.Output();

readline.createInterface({
    input: process.stdin,
    output: output,
    terminal: false
}).on('line', function(line) {
    output.write(JSON.parse(line));
});

output.pipe(process.stdout);
