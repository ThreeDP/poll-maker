FROM golang:1.22.0-alpine3.18
COPY ./conf/run-node.sh .
RUN chmod +x run-node.sh

EXPOSE 5050
WORKDIR /api

ENTRYPOINT ["ash", "/run-golang.sh"]