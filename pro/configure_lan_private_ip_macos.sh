#!/bin/sh

getPrivateIp() {
    interface=$(route -n get default | grep interface | awk '{print $2}')
    ip=$(ipconfig getifaddr "$interface")
    echo "$ip"
}

LAN_PRIVATE_IP=$(getPrivateIp)
if [ -z "$LAN_PRIVATE_IP" ]; then
    echo "No LAN private IP found"
    echo "Specify the LAN private IP in the .env file"
    exit 1
fi

# Replace the LAN_PRIVATE_IP in the .env file
sed -i'' -e "s/LAN_PRIVATE_IP=.*/LAN_PRIVATE_IP=$LAN_PRIVATE_IP/g" .env