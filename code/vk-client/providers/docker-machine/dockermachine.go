package docker_machine

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/virtual-kubelet/virtual-kubelet/errdefs"
	"github.com/virtual-kubelet/virtual-kubelet/log"
	"github.com/virtual-kubelet/virtual-kubelet/node/api"
	"github.com/virtual-kubelet/virtual-kubelet/providers"
	"github.com/virtual-kubelet/virtual-kubelet/trace"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	defaultCPUCapacity    = "20"
	defaultMemoryCapacity = "100Gi"
	defaultPodCapacity    = "20"
)

// DockerMachineProvider implements the virtual-kubelet provider interface and runs locally, pretending to run pods
type DockerMachineProvider struct {
	nodeName           string
	operatingSystem    string
	internalIP         string
	daemonEndpointPort int32
	pods               map[string]*v1.Pod
	config             MockConfig
	startTime          time.Time
	notifier           func(*v1.Pod)
}
type MockConfig struct {
	CPU    string `json:"cpu,omitempty"`
	Memory string `json:"memory,omitempty"`
	Pods   string `json:"pods,omitempty"`
}

func NewDockerMachineProvider(providerConfig, nodeName, operatingSystem string, internalIP string, daemonEndpointPort int32) (*DockerMachineProvider, error) {

	provider := DockerMachineProvider{
		nodeName:           nodeName,
		operatingSystem:    operatingSystem,
		internalIP:         internalIP,
		daemonEndpointPort: daemonEndpointPort,
		pods:               make(map[string]*v1.Pod),
		config:             MockConfig{CPU: defaultCPUCapacity, Memory: defaultMemoryCapacity, Pods: defaultPodCapacity},
		startTime:          time.Now(),
		// By default notifier is set to a function which is a no-op. In the event we've implemented the PodNotifier interface,
		// it will be set, and then we'll call a real underlying implementation.
		// This makes it easier in the sense we don't need to wrap each method.
		notifier: func(*v1.Pod) {},
	}

	return &provider, nil
}

func (s *DockerMachineProvider) CreatePod(ctx context.Context, pod *v1.Pod) error {
	ctx, span := trace.StartSpan(ctx, "CreatePod")
	defer span.End()

	ctx = addAttributes(ctx, span, "namespace", pod.Namespace, "name", pod.Name)

	now := metav1.NewTime(time.Now())
	pod.Status = v1.PodStatus{
		Phase:     v1.PodRunning,
		HostIP:    "1.2.3.4",
		PodIP:     "5.6.7.8",
		StartTime: &now,
		Conditions: []v1.PodCondition{
			{
				Type:   v1.PodInitialized,
				Status: v1.ConditionTrue,
			},
			{
				Type:   v1.PodReady,
				Status: v1.ConditionTrue,
			},
			{
				Type:   v1.PodScheduled,
				Status: v1.ConditionTrue,
			},
		},
	}
	for _, container := range pod.Spec.Containers {
		pod.Status.ContainerStatuses = append(pod.Status.ContainerStatuses, v1.ContainerStatus{
			Name:         container.Name,
			Image:        container.Image,
			Ready:        true,
			RestartCount: 0,
			State: v1.ContainerState{
				Running: &v1.ContainerStateRunning{
					StartedAt: now,
				},
			},
		})
	}
	s.pods[pod.Name] = pod
	s.notifier(pod)
	return nil
}

func (s *DockerMachineProvider) UpdatePod(ctx context.Context, pod *v1.Pod) error {
	ctx, span := trace.StartSpan(ctx, "UpdatePod")
	defer span.End()

	// Add the pod's coordinates to the current span.
	ctx = addAttributes(ctx, span, "namespace", pod.Namespace, "name", pod.Name)

	log.G(ctx).Infof("receive UpdatePod %q", pod.Name)

	key, err := buildKey(pod)
	if err != nil {
		return err
	}

	s.pods[key] = pod
	s.notifier(pod)
	return nil
}

func (*DockerMachineProvider) DeletePod(ctx context.Context, pod *v1.Pod) error {
	panic("Delete pod")
}

func (s *DockerMachineProvider) GetPod(ctx context.Context, namespace, name string) (*v1.Pod, error) {
	var err error
	ctx, span := trace.StartSpan(ctx, "GetPod")
	defer func() {
		span.SetStatus(err)
		span.End()
	}()

	ctx = addAttributes(ctx, span, "namespace", namespace, "name", name)

	log.G(ctx).Infof("receive GetPod %q", name)

	key, err := fmt.Sprintf("%s-%s", namespace, name), nil
	if err != nil {
		return nil, err
	}

	if pod, ok := s.pods[key]; ok {
		return pod, nil
	}
	return nil, errdefs.NotFoundf("pod \"%s/%s\" is not known to the provider", namespace, name)
}

func (s *DockerMachineProvider) GetPodStatus(ctx context.Context, namespace, name string) (*v1.PodStatus, error) {
	pod, ok := s.pods[name]
	if !ok {
		return nil, nil
	}

	return &pod.Status, nil
}

func (s *DockerMachineProvider) GetPods(ctx context.Context) ([]*v1.Pod, error) {
	ctx, span := trace.StartSpan(ctx, "GetPods")
	defer span.End()

	log.G(ctx).Info("receive GetPods")
	var pods []*v1.Pod
	for _, pod := range s.pods {
		pods = append(pods, pod)
	}
	return pods, nil
}

func (*DockerMachineProvider) GetContainerLogs(ctx context.Context, namespace, podName, containerName string, opts api.ContainerLogOpts) (io.ReadCloser, error) {
	panic("Get Container Logs")
}

func (*DockerMachineProvider) RunInContainer(ctx context.Context, namespace, podName, containerName string, cmd []string, attach api.AttachIO) error {
	panic("RunInContainer")
}

func (*DockerMachineProvider) Capacity(ctx context.Context) v1.ResourceList {

	ctx, span := trace.StartSpan(ctx, "Capacity")
	defer span.End()

	return v1.ResourceList{
		"cpu":    resource.MustParse(defaultCPUCapacity),
		"memory": resource.MustParse(defaultMemoryCapacity),
		"pods":   resource.MustParse(defaultPodCapacity),
	}
}

func (*DockerMachineProvider) NodeConditions(ctx context.Context) []v1.NodeCondition {
	//	panic("NodeConditions")
	ctx, span := trace.StartSpan(ctx, "NodeConditions")
	defer span.End()

	return []v1.NodeCondition{{
		Type:               "Ready",
		Status:             v1.ConditionTrue,
		LastHeartbeatTime:  metav1.Now(),
		LastTransitionTime: metav1.Now(),
		Reason:             "KubeletReady",
		Message:            "kubelet is ready.",
	},
	}
}

func (s *DockerMachineProvider) NodeAddresses(ctx context.Context) []v1.NodeAddress {
	ctx, span := trace.StartSpan(ctx, "NodeAddresses")
	defer span.End()

	return []v1.NodeAddress{
		{
			Type:    "InternalIP",
			Address: s.internalIP,
		},
	}
}

func (s *DockerMachineProvider) NodeDaemonEndpoints(ctx context.Context) *v1.NodeDaemonEndpoints {
	ctx, span := trace.StartSpan(ctx, "NodeDaemonEndpoints")
	defer span.End()

	return &v1.NodeDaemonEndpoints{
		KubeletEndpoint: v1.DaemonEndpoint{
			Port: s.daemonEndpointPort,
		},
	}
}

func (s *DockerMachineProvider) OperatingSystem() string {
	return providers.OperatingSystemLinux
}

func addAttributes(ctx context.Context, span trace.Span, attrs ...string) context.Context {
	if len(attrs)%2 == 1 {
		return ctx
	}
	for i := 0; i < len(attrs); i += 2 {
		ctx = span.WithField(ctx, attrs[i], attrs[i+1])
	}
	return ctx
}

func buildKey(pod *v1.Pod) (string, error) {
	if pod.ObjectMeta.Namespace == "" {
		return "", fmt.Errorf("pod namespace not found")
	}

	if pod.ObjectMeta.Name == "" {
		return "", fmt.Errorf("pod name not found")
	}

	return fmt.Sprintf("%s-%s", pod.ObjectMeta.Namespace, pod.ObjectMeta.Name), nil
}

func (s *DockerMachineProvider) NotifyPods(ctx context.Context, notifier func(*v1.Pod)) {
	s.notifier = notifier
}
