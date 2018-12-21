OutFile "timimi-Installer.exe"
Name "Timimi for Tiddlywiki5" 
RequestExecutionLevel user
BrandingText " "
Icon "logo.ico"
LicenseData license.txt

!define APPNAME "Timimi"
InstallDir "$APPDATA\${APPNAME}"

Page license
Page components
Page instfiles
UninstPage uninstConfirm
UninstPage instfiles

Section "Main Program"
SectionIn RO ; Read only, always installed
SetOutPath $INSTDIR
 
File "timimi.exe"
File "timimi.json"
File "logo.ico"
File "scripts.ico"
File "license.txt"
File "readme.txt"

WriteRegStr HKCU "SOFTWARE\Mozilla\NativeMessagingHosts\timimi" "" "$INSTDIR\timimi.json"

WriteUninstaller $INSTDIR\uninstaller.exe

createDirectory "$INSTDIR\Scripts"

SectionEnd

Section "Start Menu Shortcuts"
createDirectory "$SMPROGRAMS\${APPNAME}"
createShortCut "$SMPROGRAMS\${APPNAME}\Scripts.lnk" "$INSTDIR\Scripts" "" "$INSTDIR\scripts.ico"
createShortCut "$SMPROGRAMS\${APPNAME}\Uninstall.lnk" "$INSTDIR\uninstaller.exe"
SectionEnd




Section "Uninstall"
 
# Always delete uninstaller first
Delete $INSTDIR\uninstaller.exe
 
# now delete installed file
Delete $INSTDIR\timimi.exe
Delete $INSTDIR\logo.ico
Delete $INSTDIR\scripts.ico
Delete $SMPROGRAMS\${APPNAME}\Scripts.lnk
Delete $SMPROGRAMS\${APPNAME}\Uninstall.lnk
DeleteRegKey HKCU "SOFTWARE\Mozilla\NativeMessagingHosts\timimi"
rmDir "$SMPROGRAMS\${APPNAME}"
rmDir $INSTDIR
SectionEnd

