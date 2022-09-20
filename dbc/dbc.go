package dbc

import (
	"github.com/tikafog/of/dbc/bootnode"
	"github.com/tikafog/of/dbc/kernel"
	"github.com/tikafog/of/dbc/media"
	"github.com/tikafog/of/dbc/upgrade"
)

type Client interface {
	*kernel.Client | *bootnode.Client | *upgrade.Client | *media.Client
}

type DBC2[C Client] struct {
	opt *Option
}

type DBC struct {
	opt     *Option
	clients map[any]any
	clik    *kernel.Client
	clib    *bootnode.Client
	cliu    *upgrade.Client
	clim    *media.Client
}

var (
	BootnodeClient = new(bootnode.Client)
	KernelClient   = new(kernel.Client)
	UpgradeClient  = new(upgrade.Client)
	MediaClient    = new(media.Client)
)

func Open(path string, opts ...Opts) (*DBC, error) {
	o := parseOption(opts)
	debug = o.debug

	clients := make(map[any]any)
	cliB, err := openClient[*bootnode.Client](path, o)
	if err != nil {
		return nil, err
	}
	clients[BootnodeClient] = cliB

	cliK, err := openClient[*kernel.Client](path, o)
	if err != nil {
		return nil, err
	}
	clients[KernelClient] = cliK
	cliU, err := openClient[*upgrade.Client](path, o)
	if err != nil {
		return nil, err
	}
	clients[UpgradeClient] = cliU

	cliM, err := openClient[*media.Client](path, o)
	if err != nil {
		return nil, err
	}
	clients[MediaClient] = cliM
	//todo: check dbc supported
	return &DBC{
		opt:     o,
		clients: clients,
	}, nil
}

func (d *DBC) BootNode() *bootnode.Client {
	return d.clients[BootnodeClient].(*bootnode.Client)
}

func (d *DBC) Kernel() *kernel.Client {
	return d.clients[KernelClient].(*kernel.Client)
}

func (d *DBC) Upgrade() *upgrade.Client {
	return d.clients[UpgradeClient].(*upgrade.Client)
}

func (d *DBC) Media() *media.Client {
	return d.clients[MediaClient].(*media.Client)
}

func (d *DBC) Client(v any) (any, bool) {
	c, ok := d.clients[v]
	return c, ok
}
