#!/bin/sh

if [ "$RUN_WITH_SCRIPT" = 'false' ]; then
  echo '------------------------'
  echo ''
  echo 'Do not run this docker-compose file via "docker compose up" directly.'
  echo 'Please run it via the provided scripts.'
  echo ''
  echo '    - Linux: ./openvidu_linux.sh start'
  echo '    - MacOS: ./openvidu_macos.sh start'
  echo '    - Windows: ./openvidu_windows.bat start'
  echo ''
  echo '------------------------'
  exit 1
fi

if [ -z "$LAN_PRIVATE_IP" ]; then
  echo 'LAN_PRIVATE_IP is required'
  echo 'Valid values are: "none", "auto" or a valid IP address'
  echo 'Define it in the .env file'
  exit 1
fi
if [ "$LAN_MODE" = 'true' ] && [ "$USE_TLS" = 'false' ]; then
  echo 'LAN_MODE cannot be "true" if USE_TLS is "false"'
  exit 1
fi

# Prepare volumes
mkdir -p /minio/data &&
mkdir -p /mongo/data &&
mkdir -p /mongo/data/ &&
mkdir -p /egress/home/egress &&
chown 1001:1001 /minio /minio/data
chown 1001:1001 /mongo /mongo/data
chown 1001:1001 /egress
chown 1001:1001 /egress/home
chown 1001:1001 /egress/home/egress
