#!/bin/sh

KUBECTL="/usr/local/bin/kubectl"

# Get only nodes which are not drained yet
NOT_READY_NODES=$($KUBECTL get nodes --watch | grep -P 'NotReady' | awk '{print $1}' | xargs echo)
# Get only nodes which are still drained

echo "Ready nodes: $READY_NODES"


for node in $NOT_READY_NODES; do
  echo "Node $node not drained yet, draining..."
  $KUBECTL drain --ignore-daemonsets --force $node
  echo "Done"
done;

for node in $READY_NODES; do
  echo "Node $node still drained, uncordoning..."
  $KUBECTL uncordon $node
  echo "Done"
done;
