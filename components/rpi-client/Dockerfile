FROM alpine

RUN apk update && apk add curl util-linux

COPY run.sh run.sh

CMD ["sh","run.sh"]