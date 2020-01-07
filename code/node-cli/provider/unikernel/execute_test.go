package unikernel

import (
	"context"
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
	"time"
)

func TestExecute(t *testing.T) {

	fmt.Println("hello")
	p := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{Name: "Custom"},
		Spec: v1.PodSpec{
			Containers: []v1.Container{{Image: "mac-demo", Args: []string{"YELLOW"}}},
		},
	}
	cancel, _ := Execute(context.Background(), p)
	time.Sleep(2 * time.Second)
	cancel()

	time.Sleep(4 * time.Second)
	p.Spec.Containers[0].Args = []string{"GREEN"}
	cancel, _ = Execute(context.Background(), p)
	time.Sleep(2 * time.Second)
	cancel()
	p.Spec.Containers[0].Args = []string{"new year"}
	cancel, _ = Execute(context.Background(), p)
	time.Sleep(2 * time.Second)
	cancel()
}

func Test_buildCommand(t *testing.T) {
	type args struct {
		pod *v1.Pod
	}
	p := v1.Pod{}
	p.Name = "custom"
	p.Spec.Containers = []v1.Container{{Image: "mac-demo", Args: []string{"YELLOW"}}}
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
			"/Users/atakanyenel/Desktop/Computer_Science/go/bin/mac-demo YELLOW",
		}, {
			"inpath",
			args{k},
			"/bin/ls YELLOW",
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildCommand(tt.args.pod); got != tt.want {
				t.Errorf("buildCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
