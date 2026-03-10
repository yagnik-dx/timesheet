@echo off
setlocal

set APP_NAME=timesheet
set BIN_DIR=%USERPROFILE%\bin

if not exist "%APP_NAME%.exe" (
  echo [x] %APP_NAME%.exe not found! 
  echo Please make sure you have the pre-built %APP_NAME%.exe in the same folder as this script.
  exit /b 1
)

echo [*] Creating bin directory at %BIN_DIR%...
if not exist "%BIN_DIR%" mkdir "%BIN_DIR%"

echo [*] Installing %APP_NAME% to %BIN_DIR%...
copy /Y %APP_NAME%.exe "%BIN_DIR%\%APP_NAME%.exe"

echo [*] Adding %BIN_DIR% to PATH (user scope)...
setx PATH "%PATH%;%BIN_DIR%" >nul

echo.
echo [SUCCESS] Installation complete!
echo [+] Close and reopen CMD/PowerShell
echo [+] Run: %APP_NAME%
echo.

endlocal
