echo "Starting registry script"

echo "Copying virtual kubelet binary"

cp ../../code/vk-client/bin/virtual-kubelet .
cp ../../code/vk-client/bin/linux-virtual-kubelet .

echo "generating kubeconfig"
kubectl config view --flatten --minify > shareconfig.txt

echo "Starting server"

python3 -m http.server -b 0.0.0.0