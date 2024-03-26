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

URL=$(getDeploymentUrl)

for i in $(seq 1 10); do
  echo 'Starting OpenVidu... Please be patient...'
  sleep 1
done;
echo ''
echo ''
echo '------------------------'
echo 'OpenVidu is ready!'
echo '------------------------'
echo ''
echo '🎉🎉🎉 Welcome Page: http://localhost:8090' 🎉🎉🎉
echo ''
echo '------------------------'
if [ "${USE_TLS}" = 'true' ]; then
  echo '========================'
  echo 'HTTPS services:'
  echo '========================'
  if [ "${LAN_MODE}" = 'true' ]; then
    echo 'NOTE: You can access all of these services on any device'
    echo 'connected to the same network as this machine'
  fi
  echo "- OpenVidu and LiveKit API: ${URL}"
  echo "- OpenVidu Dashboard: ${URL}/dashboard"
  echo "- Minio Console: ${URL}/minio-console"
  echo "- OpenVidu Call: ${URL}/openvidu-call"
  echo "- Your App*: ${URL}"
  echo " *: Any application deployed at port 5442 will be accessible through ${URL}"
fi
echo ''
echo '========================'
echo 'HTTP services:'
echo '========================'
echo "- OpenVidu and LiveKit API: http://localhost:8090"
echo "- OpenVidu Dashboard: http://localhost:8090/dashboard"
echo "- Minio Console: http://localhost:8090/minio-console"
if [ "${USE_TLS}" = 'false' ]; then
  echo '- OpenVidu Call: http://localhost:8090/openvidu-call'
fi
echo ''
echo '========================'
echo 'Credentials:'
echo '========================'
echo 'OpenVidu Basic Auth:'
echo '  - Username: OPENVIDUAPP'
echo "  - Password: ${OPENVIDU_SHIM_SECRET}" 
echo 'LiveKit API:'
echo "  - Username: ${LIVEKIT_API_KEY}"
echo "  - Password: ${LIVEKIT_API_SECRET}"
echo 'OpenVidu Dashboard:'
echo "  - Username: ${DASHBOARD_ADMIN_USERNAME}" 
echo "  - Password: ${DASHBOARD_ADMIN_PASSWORD}"
echo 'Minio:'
echo "  - Access Key: ${MINIO_ACCESS_KEY}"
echo "  - Secret Key: ${MINIO_SECRET_KEY}"
echo '------------------------'
echo ''
echo ''
