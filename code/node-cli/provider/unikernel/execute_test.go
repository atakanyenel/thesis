package unikernel

import (
	"context"
	"fmt"
	v1 "k8s.io/api/core/v1"
	"os/exec"
	"testing"
	"time"
)

func TestExecute(t *testing.T) {

	p := &v1.Pod{}
	Execute(context.Background(), p)

	fmt.Printf("%v", p.Status.Phase)

}

func Test_buildCommand(t *testing.T) {
	type args struct {
		pod *v1.Pod
	}
	p := v1.Pod{}
	p.Name = "custom"
	p.Spec.Containers = []v1.Container{{Image: "my-unikernel-image", Env: []v1.EnvVar{{Name: "sensor", Value: "Temperature"}}}}
	k := p.DeepCopy()
	k.Spec.Containers[0].Image = "ls"
	tests := []struct {
		name string
		args args
		want string
	}{
		{"empty",
			args{pod: &v1.Pod{}},
			defaultCommand,
		},
		{
			"custom",
			args{&p},
			defaultCommand, //Todo: change to error
		}, {
			"inpath",
			args{k},
			"/bin/ls --sensor=Temperature",
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildCommand(tt.args.pod); got != tt.want {
				t.Errorf("buildCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRunCommand(t *testing.T) {
	commandContext, cancel := context.WithCancel(context.Background())

	cmd := exec.CommandContext(commandContext, "sh", "-c", "program")
	err := cmd.Start()
	if err != nil {
		t.Errorf("%s", err)
	}
	go func() {

		err = cmd.Wait()
		if err != nil {
			t.Errorf("%s", err)
		}

	}()

	time.Sleep(10 * time.Second)
	cancel() //cancel running command
	fmt.Println("asd")
}
