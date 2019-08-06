FROM alpine

RUN apk update && apk add curl bash  && \
curl -LO https://storage.googleapis.com/kubernetes-release/release/`curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt`/bin/linux/amd64/kubectl \
&&chmod +x ./kubectl \
&& mv ./kubectl /usr/local/bin/kubectl


COPY shareconfig-server-linux shareconfig-server
COPY shareconfig.txt shareconfig.txt
EXPOSE 8080

CMD ["./shareconfig-server"]
