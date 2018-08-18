#!/bin/sh
# Copyright 2013 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license

set -e
DIR="$( cd "$( dirname "$0" )" && pwd )"
if [ "$(uname -s)" = "Darwin" ]; then
  if [ "$(whoami)" = "root" ]; then
    TARGET_DIR="/Library/Google/Chrome/NativeMessagingHosts"
  else
    TARGET_DIR="$HOME/Library/Application Support/Google/Chrome/NativeMessagingHosts"
  fi
else
  if [ "$(whoami)" = "root" ]; then
    TARGET_DIR="/etc/opt/chrome/native-messaging-hosts"
  else
    TARGET_DIR="$HOME/.config/google-chrome/NativeMessagingHosts"
  fi
fi
HOST_NAME=timimi
# Create directory to store native messaging host.
mkdir -p "$TARGET_DIR"
# Copy native messaging host manifest. CHANGE
cp "$DIR/$HOST_NAME.json" "$TARGET_DIR"
cp "$DIR/$HOST_NAME.py" "$TARGET_DIR"
# Update host path in the manifest.
HOST_PATH="$TARGET_DIR/$HOST_NAME.py"
ESCAPED_HOST_PATH=${HOST_PATH////\\/}
sed -i -e "s/HOST_PATH/$ESCAPED_HOST_PATH/" "$TARGET_DIR/$HOST_NAME.json"
# Set permissions for the manifest so that all users can read it.
chmod o+r "$TARGET_DIR/$HOST_NAME.json"
echo "Native messaging host $HOST_NAME has been installed."