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
	tests := []struct {
		name string
		args args
		want string
	}{
		{"empty",
			args{pod: &v1.Pod{}},
			"ncat -l 8080", // TODO: Add test cases.
		},
		{
			"custom",
			args{&p},
			"x",
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildCommand(tt.args.pod); got != tt.want {
				t.Errorf("buildCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
