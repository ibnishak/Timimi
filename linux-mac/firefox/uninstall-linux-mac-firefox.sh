#!/bin/sh
# Copyright 2013 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license

set -e
if [ "$(uname -s)" = "Darwin" ]; then
  if [ "$(whoami)" = "root" ]; then
    TARGET_DIR="/Library/Application Support/Mozilla/NativeMessagingHosts"
  else
    TARGET_DIR="$HOME/Library/Application Support/Mozilla/NativeMessagingHosts"
  fi
else
  if [ "$(whoami)" = "root" ]; then
    TARGET_DIR="/usr/lib/mozilla/native-messaging-hosts"
  else
    TARGET_DIR="$HOME/.mozilla/native-messaging-hosts"
  fi
fi
HOST_NAME=timimi
rm "$TARGET_DIR/$HOST_NAME.json"
rm "$TARGET_DIR/$HOST_NAME.py"
echo "Native messaging host $HOST_NAME has been uninstalled."