@echo off

go build

set INSTALL_DIR=%USERPROFILE%\pin

if not exist %INSTALL_DIR% mkdir %INSTALL_DIR%

copy /y pin.exe "%INSTALL_DIR%\"

set PATH="%PATH%;%INSTALL_DIR%"
