#!/bin/sh
set -e

if [ "$LAN_PRIVATE_IP" != "" ]; then
    export NODE_IP="$LAN_PRIVATE_IP"
fi

./livekit-server "$@"
