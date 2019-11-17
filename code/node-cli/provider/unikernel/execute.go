package unikernel

import (
	"context"
	"github.com/virtual-kubelet/virtual-kubelet/log"
	v1 "k8s.io/api/core/v1"
	"os/exec"
	"reflect"
)

const defaultCommand = "ncat -l 8080"

func Execute(ctx context.Context, pod *v1.Pod) (context.CancelFunc, *exec.Cmd) {

	commandContext, cancel := context.WithCancel(context.Background())

	cmd := exec.CommandContext(commandContext, "sh", "-c", buildCommand(pod))

	log.G(ctx).Infof("starting Command", cmd)
	err := cmd.Start()
	if err != nil {
		log.G(ctx).Warnf("Couldn't start command", err)
		pod.Status.Phase = v1.PodFailed
	}
	go func() {

		err = cmd.Wait()
		if err != nil {
			log.G(ctx).Warnf("Waiting on cmd", err)
		}

	}()
	return cancel, cmd
}

func buildCommand(pod *v1.Pod) string {
	if true || reflect.DeepEqual(pod, &v1.Pod{}) { //Fixme: remove true
		return defaultCommand
	}

	return "x"
}
