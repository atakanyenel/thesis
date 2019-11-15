#curl $REGISTRY_SERVER/shareconfig.txt > kubeconfig.yaml
#curl $REGISTRY_SERVER/linux-virtual-kubelet > virtual-kubelet

ID=rpi-$(uuidgen | tr -d - | tr -d '\n' | tr '[:upper:]' '[:lower:]')
echo    $ID

./virtual-kubelet --nodename $ID --kubeconfig /tmp/kubeconfig.yaml --labels $LABELS --podcapacity 20