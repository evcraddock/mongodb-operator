FROM google/cloud-sdk:alpine

RUN apk add --update --no-cache mongodb-tools

ADD ./backup.sh /mongodb-backup/backup.sh
WORKDIR /mongodb-backup

ENTRYPOINT ["/mongodb-backup/backup.sh"]