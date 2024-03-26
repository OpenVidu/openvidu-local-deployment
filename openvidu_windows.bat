@echo off
setlocal enabledelayedexpansion

:: Function to show help
call :showHelp || exit /b 1
goto :main

:showHelp
echo.
echo Run OpenVidu Local Deployment in Windows
echo.
echo ------- 
echo  Usage  
echo ------- 
echo   %~nx0 ^<command^>
echo.
echo ---------- 
echo  Commands  
echo ---------- 
echo   start               - Start OpenVidu
echo   stop                - Stop OpenVidu
echo   help                - Show this help
echo.
exit /b 0

:getPrivateIp
for /f "tokens=4" %%i in ('route print ^| findstr "\<0.0.0.0\>"') do (
    set ip=%%i
    goto :eof
)
goto :eof

:main
if "%1"=="" (
    call :showHelp
    goto :eof
)
if /i "%1"=="start" (
    call :startOpenVidu
    goto :eof
)
if /i "%1"=="stop" (
    call :stopOpenVidu
    goto :eof
)
if /i "%1"=="help" (
    call :showHelp
    goto :eof
)
echo Not a valid command. For usage information: ".\%~nx0 help"
exit /b 0


:startOpenVidu
if exist .env (
    for /f "usebackq tokens=*" %%a in (".env") do (
        echo %%a | findstr /v /b /c:"#" >nul
        if not errorlevel 1 set %%a
    )
)

if "%LAN_PRIVATE_IP%"=="auto" (
    call :getPrivateIp || exit /b 1
    if defined ip (
        set LAN_PRIVATE_IP=!ip!
    ) else (
        set LAN_PRIVATE_IP=none
    )
)

echo Starting OpenVidu...
set RUN_WITH_SCRIPT=true
docker-compose down --volumes || exit /b 1
docker-compose up || exit /b 1
exit /b 0

:stopOpenVidu
echo Stopping OpenVidu...
docker-compose down --volumes || exit /b 1
exit /b 0
