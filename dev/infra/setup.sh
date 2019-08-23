
INSTANCE=microk8s-test

gcloud compute instances create ${INSTANCE} \
--machine-type=custom-2-8192 \
--image-family=ubuntu-minimal-1904 \
--image-project=ubuntu-os-cloud \
--metadata=startup-script='
!# /bin/bash
sudo snap install microk8s - classic
&& sudo snap alias microk8s.kubectl k
'

gcloud compute scp \
${INSTANCE}:/var/snap/microk8s/current/credentials/client.config ~/.kube/config \



gcloud compute ssh ${INSTANCE} \
--ssh-flag="-L 16443:localhost:16443" 

