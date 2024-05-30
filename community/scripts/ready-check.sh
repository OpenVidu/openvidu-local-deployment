#!/bin/sh

. /scripts/utils.sh

trap 'handle_sigint' SIGINT

handle_sigint() {
  echo "SIGINT signal received, exiting..."
  exit 1
}

wait_for_service() {
  SERVICE_NAME=$1
  SERVICE_URL=$2
  shift 2
  EXTRA=$@
  if [ -n "$EXTRA" ]; then
    until curl $EXTRA $SERVICE_URL > /dev/null; do
      echo "Waiting for $SERVICE_NAME to start...";
      sleep 1;
    done;
  else
    until curl --silent --head --fail $SERVICE_URL > /dev/null; do
      echo "Waiting for $SERVICE_NAME to start...";
      sleep 1;
    done;
  fi;
}

wait_for_service 'OpenVidu' 'http://openvidu:7880'
wait_for_service 'Ingress' 'http://ingress:9091'
wait_for_service 'Egress' 'http://egress:9091'
wait_for_service 'Dashboard' 'http://dashboard:5000'
wait_for_service 'Minio' 'http://minio:9000/minio/health/live'
wait_for_service 'Minio Console' 'http://minio:9001/minio-console'
wait_for_service 'Mongo' 'http://mongo:27017' --connect-timeout 10 --silent

LAN_HTTP_URL=$(getDeploymentUrl http)
LAN_WS_URL=$(getDeploymentUrl ws)

for i in $(seq 1 10); do
  echo 'Starting OpenVidu... Please be patient...'
  sleep 1
done;
echo ''
echo ''
echo '========================================='
echo '🎉 OpenVidu is ready! 🎉'
echo '========================================='
echo ''
echo 'OpenVidu Server & LiveKit Server URLs:'
echo ''
echo '    - From this machine:'
echo ''
echo '        - http://localhost:7880'
echo '        - ws://localhost:7880'
echo ''
echo '    - From other devices in your LAN:'
echo ''
echo "        - $LAN_HTTP_URL"
echo "        - $LAN_WS_URL"
echo ''
echo '========================================='
echo ''
echo 'OpenVidu Developer UI (services and passwords):'
echo ''
echo '    - http://localhost:7880'
echo "    - $LAN_HTTP_URL"
echo ''
echo '========================================='
