package dbc

import (
	"github.com/tikafog/of/dbc/bootnode"
	"github.com/tikafog/of/dbc/kernel"
	"github.com/tikafog/of/dbc/media"
	"github.com/tikafog/of/dbc/upgrade"
)

type DBC struct {
	opt *Option
	cs  [ClientTypeMax]any
}

func Open(path string, opts ...Opts) (*DBC, error) {
	o := parseOption(opts)
	debug = o.Debug()
	ignores := o.Ignores()
	dbc := &DBC{
		opt: o,
	}
	var err error
	if _, ok := ignores[ClientTypeBootnode]; !ok {
		dbc.cs[ClientTypeBootnode], err = openClient[bootnode.Client](path, o)
		if err != nil {
			return nil, merr.Wrap(err, "open bootnode failed")
		}
	}
	if _, ok := ignores[ClientTypeKernel]; !ok {
		dbc.cs[ClientTypeKernel], err = openClient[kernel.Client](path, o)
		if err != nil {
			return nil, merr.Wrap(err, "open kernel failed")
		}
	}
	if _, ok := ignores[ClientTypeUpgrade]; !ok {
		dbc.cs[ClientTypeUpgrade], err = openClient[upgrade.Client](path, o)
		if err != nil {
			return nil, merr.Wrap(err, "open upgrade failed")
		}
	}
	if _, ok := ignores[ClientTypeMedia]; !ok {
		dbc.cs[ClientTypeMedia], err = openClient[media.Client](path, o)
		if err != nil {
			return nil, merr.Wrap(err, "open media failed")
		}
	}
	return dbc, nil
}

func (d *DBC) BootNode() *Client[bootnode.Client] {
	return d.cs[ClientTypeBootnode].(*Client[bootnode.Client])
}

func (d *DBC) Kernel() *Client[kernel.Client] {
	return d.cs[ClientTypeKernel].(*Client[kernel.Client])
}

func (d *DBC) Upgrade() *Client[upgrade.Client] {
	return d.cs[ClientTypeUpgrade].(*Client[upgrade.Client])
}

func (d *DBC) Media() *Client[media.Client] {
	return d.cs[ClientTypeMedia].(*Client[media.Client])
}

func (d *DBC) Client(p ClientType) (any, bool) {
	if p >= ClientTypeMax {
		return nil, false
	}
	return d.cs[p], true
}
