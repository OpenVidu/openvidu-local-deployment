#!/bin/bash

. /scripts/utils.sh

LIVEKIT_URL=$(getDeploymentUrl ws)
MEET_BASE_URL=$(getDeploymentUrl http meet)
export LIVEKIT_URL="${LIVEKIT_URL}"
export MEET_BASE_URL="${MEET_BASE_URL}"

echo "LIVEKIT_URL is set to ${LIVEKIT_URL}"
echo "MEET_BASE_URL is set to ${MEET_BASE_URL}"

/usr/local/bin/entrypoint.sh
