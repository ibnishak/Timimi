#!/bin/sh
# Copyright 2013 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license

set -e
if [ "$(uname -s)" = "Darwin" ]; then
  if [ "$(whoami)" = "root" ]; then
   TARGET_DIR="/Library/Application Support/Chromium/NativeMessagingHosts"
  else
    TARGET_DIR="$HOME/Library/Application Support/Chromium/NativeMessagingHosts"
  fi
else
  if [ "$(whoami)" = "root" ]; then
    TARGET_DIR="/etc/chromium/native-messaging-hosts"
  else
    TARGET_DIR="$HOME/.config/chromium/NativeMessagingHosts"
  fi
fi
HOST_NAME=timimi
rm "$TARGET_DIR/$HOST_NAME.json"
rm "$TARGET_DIR/$HOST_NAME.py"
echo "Native messaging host $HOST_NAME has been uninstalled."