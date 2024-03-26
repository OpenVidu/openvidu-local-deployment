#!/bin/sh

. /scripts/utils.sh

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
echo "Open $URL in your browser"
echo '------------------------'
echo ''
echo ''
