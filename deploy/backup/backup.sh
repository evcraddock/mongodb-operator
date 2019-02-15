#! /bin/bash -ui

if [[ -z ${GCS_KEY} ]] || [[ -z ${BUCKET_PATH} ]] || [[ -z ${MONGODB_URI} ]]
then
    echo "GCS_KEY, BUCKET_PATH, and MONGODB_URI are required!"
    exit 1
fi

gcloud auth activate-service-account --key-file "${GCS_KEY}"

DATE=$(date -u "+%F-%H%M%S")
BACKUP_FILE="backup-${DATE}.tar.gz"


echo "backing up ${MONGODB_URI}"
CMD="mongodump --uri=\"$MONGODB_URI\" --oplog --gzip --archive=$BACKUP_FILE"
eval "${CMD}"

echo "finished creating backup"

shortDate=$(date +%Y-%m-%d)
gsutil cp "${BACKUP_FILE}" gs://"${BUCKET_PATH}"/"${shortDate}"/


