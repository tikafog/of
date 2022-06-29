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
	cliB, err := openClient[*bootnode.Client](path, o)
	if err != nil {
		return nil, err
	}

	cliK, err := openClient[*kernel.Client](path, o)
	if err != nil {
		return nil, err
	}

	cliU, err := openClient[*upgrade.Client](path, o)
	if err != nil {
		return nil, err
	}

	cliM, err := openClient[*media.Client](path, o)
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
