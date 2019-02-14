#! /bin/sh -ui

gcloud auth activate-service-account --key-file ${GCS_KEY}

MONGODB_HOST=${MONGODB_HOST:-localhost}
MONGODB_PORT=${MONGODB_PORT:-27017}
MONGODB_DB=${MONGODB_DB:-}
MONGODB_USER=${MONGODB_USER:-}
MONGODB_PASSWORD=${MONGODB_PASSWORD:-}

login=""
db=""

DATE=$(date -u "+%F-%H%M%S")
BACKUP_FILE="backup-${DATE}.tar.gz"

if [ ! -z ${MONGODB_USER} ] && [ ! -z ${MONGODB_PASSWORD} ]
then
  echo "setting login credentials for mongodb"
  login="--username=\"${MONGODB_USER}\" --password=\"${MONGODB_PASSWORD}\" "
fi

if [ ! -z ${MONGODB_DB} ]
then
  echo "specifying a specific database"
  db="--db=\"$MONGODB_DB\" "
fi

if [ -z ${MONGODB_HOST} ] || [ -z ${BUCKET_PATH} ]
then
    echo "missing required env variables"
    exit 1
fi

echo "backing up host: ${MONGODB_HOST}:${MONGODB_PORT}"
CMD="mongodump --host=\"$MONGODB_HOST\" --port=\"$MONGODB_PORT\" $login$db --gzip --archive=$BACKUP_FILE"
eval "${CMD}"

shortDate=`date +%Y-%m-%d`
gsutil cp ${BACKUP_FILE} gs://${BUCKET_PATH}/${shortDate}/


