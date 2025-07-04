#!/bin/bash

. /scripts/utils.sh

URL=$(getDeploymentUrl)
export LIVEKIT_URL="${URL}"

/usr/local/bin/entrypoint.sh
