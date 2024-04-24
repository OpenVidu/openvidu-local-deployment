#!/bin/sh

getPrivateIp() {
    ip="$(ip route get 8.8.8.8 | sed -n '/src/{s/.*src *\([^ ]*\).*/\1/p;q}')"
    echo "$ip"
}

LAN_PRIVATE_IP=$(getPrivateIp)
if [ -z "$LAN_PRIVATE_IP" ]; then
    echo "No LAN private IP found"
    echo "Specify the LAN private IP in the .env file"
    exit 1
fi

# Replace the LAN_PRIVATE_IP in the .env file
sed -i "s/LAN_PRIVATE_IP=.*/LAN_PRIVATE_IP=$LAN_PRIVATE_IP/g" .env
