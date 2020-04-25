FROM alpine

RUN apk update && apk add curl util-linux bash 

COPY run.sh run.sh
COPY virtual-kubelet virtual-kubelet
RUN chmod +x virtual-kubelet
CMD ["sh","run.sh"]