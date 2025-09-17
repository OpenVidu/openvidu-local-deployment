#!/bin/sh

getDeploymentUrl() {
    schema="${1:-http}"
    service="${2:-default}"

    if [ "$service" = "meet" ]; then
        URL="$schema://localhost:9080"
        if [ "${USE_HTTPS}" = 'true' ]; then
            URL="${schema}s://localhost:9443"
        fi
    else
        URL="$schema://localhost:7880"
        if [ "${USE_HTTPS}" = 'true' ]; then
            URL="${schema}s://localhost:7443"
        fi
    fi

    if [ "${LAN_MODE}" = 'true' ]; then
        LAN_DOMAIN=${LAN_DOMAIN:-"openvidu-local.dev"}
        if [ "$LAN_PRIVATE_IP" != 'none' ] && [ "${LAN_DOMAIN}" = 'openvidu-local.dev' ]; then
            # Replace dots with dashes
            LAN_DOMAIN="$(echo "$LAN_PRIVATE_IP" | sed 's/\./-/g').openvidu-local.dev"
        fi
        if [ "$service" = "meet" ]; then
            URL="${schema}s://${LAN_DOMAIN}:9443"
        else
            URL="${schema}s://${LAN_DOMAIN}:7443"
        fi
    fi
    echo "$URL"
}
