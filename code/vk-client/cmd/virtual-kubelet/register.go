package main

import (
	docker_machine "github.com/virtual-kubelet/azure-aci/providers/docker-machine"
	"github.com/virtual-kubelet/azure-aci/providers/silly"
	"github.com/virtual-kubelet/virtual-kubelet/providers"
)

func registerACI(s *providers.Store) error {

	s.Register("docker-machine", func(cfg providers.InitConfig) (providers.Provider, error) {
		return docker_machine.NewDockerMachineProvider(cfg.ConfigPath,
			cfg.NodeName,
			cfg.OperatingSystem,
			cfg.InternalIP,
			cfg.DaemonPort)
	})
	return s.Register("silly", func(cfg providers.InitConfig) (providers.Provider, error) {
		return silly.NewSillyProvider(cfg.ConfigPath,
			cfg.NodeName,
			cfg.OperatingSystem,
			cfg.InternalIP,
			cfg.DaemonPort)
	})
}
