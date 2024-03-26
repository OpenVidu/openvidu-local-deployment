#!/bin/sh

getDeploymentUrl() {
    URL="http://localhost:8090"
    if [ "${USE_TLS}" = 'true' ]; then
        URL="https://localhost:4443"
    fi
    if [ "${LAN_MODE}" = 'true' ]; then
        LAN_DOMAIN=${LAN_DOMAIN:-"openvidu-local.dev"}
        if [ "$LAN_PRIVATE_IP" != 'none' ] && [ "${LAN_DOMAIN}" = 'openvidu-local.dev' ]; then
            # Replace dots with dashes
            LAN_DOMAIN="$(echo "$LAN_PRIVATE_IP" | sed 's/\./-/g').openvidu-local.dev"
        fi
        URL="https://${LAN_DOMAIN}:4443"
    fi
    echo "$URL"
}
