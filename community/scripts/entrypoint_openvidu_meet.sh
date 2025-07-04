#!/bin/bash

. /scripts/utils.sh

URL=$(getDeploymentUrl ws)
export LIVEKIT_URL="${URL}"

/usr/local/bin/entrypoint.sh
