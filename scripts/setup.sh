#!/bin/sh

if [ "$RUN_WITH_SCRIPT" = 'false' ]; then
  echo '------------------------'
  echo ''
  echo 'Do not run this docker-compose file via "docker compose up" directly.'
  echo 'Please run it via the provided scripts.'
  echo ''
  echo '    - Linux: ./openvidu_linux.sh'
  echo '    - MacOS: ./openvidu_macos.sh'
  echo '    - Windows: ./openvidu_windows.ps1'
  echo ''
  echo '------------------------'
  exit 1
fi
if [ "$PRIVATE_IP" = '?' ]; then
  echo 'PRIVATE_IP is required'
  echo 'Define it in the .env file'
  exit 1
fi
mkdir -p /minio/data &&
mkdir -p /mongo/data &&
mkdir -p /mongo/data/ &&
mkdir -p /egress/home/egress &&
chown 1001:1001 /minio /minio/data
chown 1001:1001 /mongo /mongo/data
chown 1001:1001 /egress
chown 1001:1001 /egress/home
chown 1001:1001 /egress/home/egress
