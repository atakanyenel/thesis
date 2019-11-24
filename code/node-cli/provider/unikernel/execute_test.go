package unikernel

import (
	"context"
	"fmt"
	v1 "k8s.io/api/core/v1"
	"testing"
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
