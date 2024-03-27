#!/bin/sh

getPrivateIp() {
    ip=$(ipconfig getifaddr $(route -n get default | grep interface | awk '{print $2}'))
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

# If sillicon mac, enable EXPERIMENTAL_DOCKER_DESKTOP_FORCE_QEMU flag
if [ "$(uname -m)" = "arm64" ]; then
    if ! grep -q "EXPERIMENTAL_DOCKER_DESKTOP_FORCE_QEMU" .env; then
        echo "# Enable this flag to run Docker Desktop on Apple Silicon Macs" >> .env
        echo "EXPERIMENTAL_DOCKER_DESKTOP_FORCE_QEMU=1" >> .env
    fi
    sed -i "s/EXPERIMENTAL_DOCKER_DESKTOP_FORCE_QEMU=.*/EXPERIMENTAL_DOCKER_DESKTOP_FORCE_QEMU=1/g" .env
fi
