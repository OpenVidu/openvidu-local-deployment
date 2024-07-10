#!/bin/sh
set -e

if [ "$LAN_PRIVATE_IP" != "none" ]; then
    export NODE_IP="$LAN_PRIVATE_IP"
fi

# Configure container private IP as node private IP
export LIVEKIT_OPENVIDU_NODE_PRIVATE_IP="$(hostname -i)"

./livekit-server "$@"
