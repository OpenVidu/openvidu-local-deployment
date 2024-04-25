#!/bin/sh
set -e
. /scripts/utils.sh

URL=$(getDeploymentUrl)
export V2COMPAT_OPENVIDU_SHIM_URL="${URL}"
export V2COMPAT_LIVEKIT_URL="${URL}"
/bin/server
