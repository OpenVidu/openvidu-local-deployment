@echo off
call :getPrivateIp
if "%ip%"=="" (
    echo No LAN private IP found
    echo Specify the LAN private IP in the .env file
    exit /b 1
)

:: Replace the LAN_PRIVATE_IP in the .env file
setlocal enabledelayedexpansion
set "tempFile=%temp%\temp_env_%random%.txt"
if exist "%tempFile%" del "%tempFile%"
(
    for /f "delims=" %%i in ('findstr /n "^" ".env"') do (
        set "line=%%i"
        set "line=!line:*:=!"
        if "!line:~0,15!"=="LAN_PRIVATE_IP=" (
            echo LAN_PRIVATE_IP=%ip%
        ) else (
            echo(!line!
        )
    )
) > "%tempFile%"

move /y "%tempFile%" ".env" >nul
endlocal
exit /b 0

:getPrivateIp
for /f "tokens=4" %%i in ('route print ^| findstr "\<0.0.0.0\>"') do (
    set ip=%%i
    goto :eof
)
goto :eof
