package main

import (

	"github.com/virtual-kubelet/azure-aci/providers/unikernel"
	"github.com/virtual-kubelet/virtual-kubelet/providers"
)

func registerACI(s *providers.Store) error {


	return s.Register("unikernel", func(cfg providers.InitConfig) (providers.Provider, error) {
		return unikernel.NewProvider(cfg.ConfigPath,
			cfg.NodeName,
			cfg.OperatingSystem,
			cfg.InternalIP,
			cfg.DaemonPort)
	})
}
