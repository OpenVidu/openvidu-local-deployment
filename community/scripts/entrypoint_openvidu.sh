#!/bin/sh
set -e

if [ "$LAN_PRIVATE_IP" != "" ] && [ "$LAN_MODE" = 'true' ]; then
    echo "Using as NODE_IP: $LAN_PRIVATE_IP"
    export NODE_IP="$LAN_PRIVATE_IP"
fi

./livekit-server "$@"
