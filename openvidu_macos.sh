#!/bin/sh

showHelp() {
    echo ""
    echo "Run OpenVidu Local Deployment in Linux"
    echo ""
    echo "-------"
    echo " Usage "
    echo "-------"
    echo "  $0 <command>"
    echo ""
    echo "----------"
    echo " Commands "
    echo "----------"
    echo "  start               - Start OpenVidu"
    echo "  stop                - Stop OpenVidu"
    echo "  help                - Show this help"
    echo ""
}

getPrivateIp() {
    ip=$(ipconfig getifaddr $(route -n get default | grep interface | awk '{print $2}'))
    echo "$ip"
}

# Flags
START=false
STOP=false

if [ -n "${1:-}" ]; then
    while :; do
        case "${1:-}" in
        start)
            START=true
            shift 1
            break
            ;;
        stop)
            STOP=true
            shift 1
            break
            ;;
        help)
            showHelp
            exit 0
            ;;
        *)
            echo "Not a valid command. For usage information: \"$0 help\""
            exit 1
            ;;
        esac
    done
else
    showHelp
    exit
fi

if [ "$START" = "true" ]; then
    # Load environment variables
    if [ -f .env ]; then
        export $(grep -v '^#' .env | xargs)
    fi

    if [ "$LAN_PRIVATE_IP" = "auto" ]; then
        LAN_PRIVATE_IP="$(getPrivateIp)"
        if [ -z "$LAN_PRIVATE_IP" ]; then
            LAN_PRIVATE_IP=none
        fi
        export LAN_PRIVATE_IP
    fi

    echo "Starting OpenVidu..."
    export RUN_WITH_SCRIPT=true
    docker compose down --volumes
    docker compose up
fi

if [ "$STOP" = "true" ]; then
    echo "Stopping OpenVidu"
    docker compose down --volumes
fi
