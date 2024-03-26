#!/bin/sh
set -e
. /scripts/utils.sh

URL=$(getDeploymentUrl)
export OPENVIDU_SHIM_URL="${URL}"
export LIVEKIT_URL="${URL}"
node dist/server.js
