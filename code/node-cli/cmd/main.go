package main

import (
	"context"
	"github.com/virtual-kubelet/node-cli/provider/unikernel"

	"github.com/sirupsen/logrus"
	cli "github.com/virtual-kubelet/node-cli"
	logruscli "github.com/virtual-kubelet/node-cli/logrus"
	"github.com/virtual-kubelet/node-cli/provider"
	"github.com/virtual-kubelet/virtual-kubelet/log"
	logruslogger "github.com/virtual-kubelet/virtual-kubelet/log/logrus"
)

func main() {
	ctx := cli.ContextWithCancelOnSignal(context.Background())
	logger := logrus.StandardLogger()

	log.L = logruslogger.FromLogrus(logrus.NewEntry(logger))
	logConfig := &logruscli.Config{LogLevel: "info"}

	node, err := cli.New(ctx,
		cli.WithProvider("unikernel", func(cfg provider.InitConfig) (provider.Provider, error) {
			provider, err := unikernel.NewProvider(cfg.ConfigPath, cfg.NodeName, cfg.OperatingSystem, cfg.InternalIP, cfg.DaemonPort, cfg.PodCapacity)
			return provider, err
		}),
		// Adds flags and parsing for using logrus as the configured logger
		cli.WithPersistentFlags(logConfig.FlagSet()),
		cli.WithPersistentPreRunCallback(func() error {
			return logruscli.Configure(logConfig, logger)
		}),
	)

	if err != nil {
		panic(err)
	}

	// Notice that the context is not passed in here, this is due to limitations
	// of the underlying CLI library (cobra).
	// contexts get passed through from `cli.New()`
	//
	// Args can be specified here, or os.Args[1:] will be used.
	if err := node.Run(); err != nil {
		panic(err)
	}
}
