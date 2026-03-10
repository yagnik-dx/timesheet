@echo off
setlocal

set "APP_NAME=timesheet"
set "BIN_DIR=%USERPROFILE%\bin"
set "REPO_URL=https://raw.githubusercontent.com/yagnik-dx/timesheet/main/timesheet.exe"
set "EXE_PATH=%BIN_DIR%\%APP_NAME%.exe"

echo Downloading %APP_NAME% from GitHub...
if not exist "%BIN_DIR%" mkdir "%BIN_DIR%"

:: Try curl (Windows 10+), then PowerShell
where curl >nul 2>nul
if %ERRORLEVEL% EQU 0 (
  curl -fsSL "%REPO_URL%" -o "%EXE_PATH%"
) else (
  powershell -NoProfile -Command "Invoke-WebRequest -Uri '%REPO_URL%' -OutFile '%EXE_PATH%' -UseBasicParsing"
)

if not exist "%EXE_PATH%" (
  echo Failed to download. Get timesheet.exe from https://github.com/yagnik-dx/timesheet
  exit /b 1
)

:: Add BIN_DIR to user PATH if not already there
echo %PATH% | findstr /I /C:"%BIN_DIR%" >nul 2>nul
if %ERRORLEVEL% NEQ 0 (
  echo Adding %BIN_DIR% to your user PATH...
  setx PATH "%PATH%;%BIN_DIR%" >nul
)

echo.
echo [SUCCESS] %APP_NAME% installed to %BIN_DIR%
echo Close and reopen CMD/PowerShell, then run: %APP_NAME%
echo.

endlocal
