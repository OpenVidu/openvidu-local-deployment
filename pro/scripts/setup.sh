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
else
  # Check if the LAN_PRIVATE_IP is reachable
  if ! ping -c 1 -W 1 "$LAN_PRIVATE_IP" > /dev/null; then
    echo "ERROR: LAN_PRIVATE_IP $LAN_PRIVATE_IP is not reachable"
    echo "    Maybe you changed your network or the IP is wrong"
    echo "    Please update the LAN_PRIVATE_IP in the .env file or"
    echo "    run the configure_lan_private_ip script again:"
    echo ""
    echo "        - Linux: ./configure_lan_private_ip_linux.sh"
    echo "        - MacOS: ./configure_lan_private_ip_macos.sh"
    echo "        - Windows: .\configure_lan_private_ip_windows.bat"
    echo ""
    exit 1
  fi
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
