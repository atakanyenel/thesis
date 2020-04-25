setup:
	cp ../../code/node-cli/bin/linux-node-cli virtual-kubelet
	kubectl config view --minify --flatten > kubeconfig.yaml

rpi:
	cp ../../code/node-cli/bin/rpi-node-cli virtual-kubelet
	kubectl config view --minify --flatten > kubeconfig.yaml
	upx virtual-kubelet
	scp kubeconfig.yaml pi@raspberrypi:.
	scp virtual-kubelet pi@raspberrypi:.
