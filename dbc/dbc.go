package dbc

import (
	"github.com/tikafog/of/dbc/bootnode"
	"github.com/tikafog/of/dbc/kernel"
	"github.com/tikafog/of/dbc/media"
	"github.com/tikafog/of/dbc/upgrade"
)

type DBC struct {
	opt  *Option
	clik *kernel.Client
	clib *bootnode.Client
	cliu *upgrade.Client
	clim *media.Client
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

	cliU, err := open[*upgrade.Client]("upgrade", path, o)
	if err != nil {
		return nil, err
	}

	cliM, err := open[*media.Client]("media", path, o)
	if err != nil {
		return nil, err
	}

	return &DBC{
		opt:  o,
		clib: cliB,
		clik: cliK,
		cliu: cliU,
		clim: cliM,
	}, nil
}

func (d *DBC) BootNode() *bootnode.Client {
	return d.clib
}

func (d *DBC) Kernel() *kernel.Client {
	return d.clik
}

func (d *DBC) Upgrade() *upgrade.Client {
	return d.cliu
}

func (d *DBC) Media() *media.Client {
	return d.clim
}
