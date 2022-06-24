package dbc

import (
	"github.com/tikafog/of/dbc/bootnode"
	"github.com/tikafog/of/dbc/kernel"
)

type DBC struct {
	opt              *Option
	Bootstrap        *bootnode.BootstrapClient
	BootstrapVersion *bootnode.VersionClient
	Instruct         *kernel.InstructClient
	Pin              *kernel.PinClient
	KernelVersion    *kernel.VersionClient
}

func Open(path string, opts ...Opts) (*DBC, error) {
	o := opt(opts)
	cliB, err := open[*bootnode.Client]("bootnode", path, o)
	if err != nil {
		return nil, err
	}
	//d.Bootstrap = cliB.Bootstrap
	//d.BootstrapVersion = cliB.Version

	cliK, err := open[*kernel.Client]("kernel", path, o)
	if err != nil {
		return nil, err
	}
	//d.Instruct = cliK.Instruct
	//d.Pin = cliK.Pin
	//d.KernelVersion = cliK.Version
	return &DBC{
		opt:              o,
		Bootstrap:        cliB.Bootstrap,
		BootstrapVersion: cliB.Version,
		Instruct:         cliK.Instruct,
		Pin:              cliK.Pin,
		KernelVersion:    cliK.Version,
	}, nil
}
