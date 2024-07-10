#!/bin/sh
set -e

if [ "$LAN_PRIVATE_IP" != "none" ]; then
    export NODE_IP="$LAN_PRIVATE_IP"
fi

./livekit-server "$@"
