#!/bin/sh
set -e

if [ "$LAN_PRIVATE_IP" != "" ] && [ "$LAN_MODE" = 'true' ]; then
    echo "Using as NODE_IP: $LAN_PRIVATE_IP"
    export NODE_IP="$LAN_PRIVATE_IP"
fi

# Configure container private IP as node private IP
LIVEKIT_OPENVIDU_NODE_PRIVATE_IP="$(hostname -i)"
export LIVEKIT_OPENVIDU_NODE_PRIVATE_IP

./livekit-server "$@"
