#!/bin/sh

if [ -z "$LAN_PRIVATE_IP" ]; then
  echo '------------------------'
  echo ''
  echo 'LAN_PRIVATE_IP is required in the .env file.'
  echo 'Depending on your OS, you can run the following command to set your LAN private IP:'
  echo ''
  echo '    - Linux: ./configure_lan_private_ip_linux.sh'
  echo '    - MacOS: ./configure_lan_private_ip_macos.sh'
  echo '    - Windows: .\configure_lan_private_ip_windows.bat'
  echo ''
  echo 'The script will automatically update the .env file with the LAN_PRIVATE_IP.'
  echo 'If it can'\''t be found, you can manually set it in the .env file'
  echo '------------------------'
  exit 1
fi

if [ "$LAN_MODE" = 'true' ] && [ "$USE_HTTPS" = 'false' ]; then
  echo 'LAN_MODE cannot be "true" if USE_HTTPS is "false"'
  exit 1
fi

# Prepare volumes
mkdir -p /minio/data
mkdir -p /mongo/data
mkdir -p /mongo/data/
mkdir -p /egress/home/egress
chown 1001:1001 /minio /minio/data
chown 1001:1001 /mongo /mongo/data
chown 1001:1001 /egress
chown 1001:1001 /egress/home
chown 1001:1001 /egress/home/egress
