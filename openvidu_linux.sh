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
    ip="$(ip route get 8.8.8.8 | sed -n '/src/{s/.*src *\([^ ]*\).*/\1/p;q}')"
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

    if [ "$PRIVATE_IP" = "auto" ]; then
        PRIVATE_IP="$(getPrivateIp)"
        if [ -z "$PRIVATE_IP" ]; then
            PRIVATE_IP=none
        fi
        export PRIVATE_IP
    fi

    echo "Starting OpenVidu..."
    export RUN_WITH_SCRIPT=true
    docker compose down
    docker compose up
fi

if [ "$STOP" = "true" ]; then
    echo "Stopping OpenVidu"
    docker compose down
fi
