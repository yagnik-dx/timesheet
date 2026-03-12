@echo off
setlocal

set APP=timesheet
set BIN_DIR=%USERPROFILE%\bin
set SRC=bin\%APP%.exe

if not exist "%SRC%" (
    echo [x] %SRC% not found. Please ensure the binary exists in the bin folder.
    exit /b 1
)

if not exist "%BIN_DIR%" (
    echo [*] Creating %BIN_DIR%...
    mkdir "%BIN_DIR%"
)

echo [*] Installing %APP% to %BIN_DIR%...
copy /Y "%SRC%" "%BIN_DIR%\%APP%.exe" >nul

echo [*] Adding %BIN_DIR% to PATH...
powershell -NoProfile -Command ^
    "$p = [Environment]::GetEnvironmentVariable('Path', 'User');" ^
    "if ($p -notlike '*%BIN_DIR%*') {" ^
    "  [Environment]::SetEnvironmentVariable('Path', $p.TrimEnd(';') + ';%BIN_DIR%', 'User');" ^
    "  Write-Host 'PATH updated.'" ^
    "} else { Write-Host 'PATH already contains %BIN_DIR%.' }"

echo.
echo [OK] Done! Open a new terminal and run: %APP%
echo.

endlocal
