OutFile "timimi-Installer.exe"
Name "Timimi for Tiddlywiki5" 
RequestExecutionLevel user
BrandingText " "
Icon "logo.ico"


!define APPNAME "Timimi"

PageEx readme
   LicenseData readme.rtf
PageExEnd

PageEx license
   LicenseData license.txt
	 LicenseForceSelection radiobuttons "Accept" "Decline"
PageExEnd
Page components
Page instfiles
UninstPage uninstConfirm
UninstPage instfiles

Section "Main Program"
SectionIn RO ; Read only, always installed
InstallDir "$APPDATA\${APPNAME}"
SetOutPath $INSTDIR
 
File "timimi.exe"
File "logo.ico"
File "scripts.ico"
File "timimi.json"

WriteRegStr HKCU "SOFTWARE\Mozilla\NativeMessagingHosts\timimi" "" "$\"$INSTDIR\timimi.json$\""

WriteUninstaller $INSTDIR\uninstaller.exe

createDirectory "$INSTDIR\Scripts"

SectionEnd

Section "Start Menu Shortcuts"
createDirectory "$SMPROGRAMS\${APPNAME}"
createShortCut "$SMPROGRAMS\${APPNAME}\Scripts.lnk" "$INSTDIR\Scripts" "" "$INSTDIR\scripts.ico"
createShortCut "$SMPROGRAMS\${APPNAME}\Uninstall.lnk" "$INSTDIR\uninstaller.exe"
SectionEnd


function un.onInit
	SetShellVarContext all
 
	MessageBox MB_OKCANCEL "Permanantly remove ${APPNAME}?" IDOK next
		Abort
functionEnd

Section "Uninstaller"
 
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

