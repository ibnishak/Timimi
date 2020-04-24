#!/usr/bin/env node

var nativeMessage = require('./index');

process.stdin
    .pipe(new nativeMessage.Input())
    .pipe(new nativeMessage.Debug())
    .pipe(process.stdout)
;
