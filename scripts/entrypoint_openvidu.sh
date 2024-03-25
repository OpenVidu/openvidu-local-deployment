#!/bin/sh
set -e
CONFIG_FILE_TMP="/tmp/livekit.yaml"
CONFIG_FILE="/etc/livekit.yaml"
PRIVATE_IP="${PRIVATE_IP:-}"

cp ${CONFIG_FILE_TMP} ${CONFIG_FILE}

if [ "$PRIVATE_IP" != "none" ]; then
    if ! grep -q "^[[:space:]]*node_ip:.*" "$CONFIG_FILE"; then
        if grep -q "^rtc:" "$CONFIG_FILE"; then
            sed -i "/^rtc:/a \    node_ip: $PRIVATE_IP" "$CONFIG_FILE"
        else
            echo "rtc:" >> "$CONFIG_FILE"
            echo "    node_ip: $PRIVATE_IP" >> "$CONFIG_FILE"
        fi
    fi
fi

./livekit-server "$@"
