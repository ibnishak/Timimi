:: Copyright 2014 The Chromium Authors. All rights reserved.
:: Use of this source code is governed by a BSD-style license

REG ADD "HKEY_CURRENT_USER\SOFTWARE\Mozilla\NativeMessagingHosts\timimi" /ve /t REG_SZ /d "%~dp0timimi.json" /f