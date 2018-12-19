#!/usr/bin/env bash

tr -dc A-Za-z0-9 </dev/urandom | head -c 1024 > ~/.timimi/Timimi-First-Test-Success.txt
cat ~/.timimi/Timimi-First-Test-Success.txt
