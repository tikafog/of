package dbc

import (
	"io"

	"github.com/tikafog/of/dbc/bootnode"
	"github.com/tikafog/of/dbc/kernel"
	"github.com/tikafog/of/dbc/media"
	"github.com/tikafog/of/dbc/upgrade"
	"github.com/tikafog/of/feature/query"
)

type DBC struct {
	opt *Option
	cs  [query.ClientTypeMax]io.Closer
}

func Open(path string, opts ...Opts) (*DBC, error) {
	var dbc DBC
	dbc.opt = parseOption(opts)
	debug = dbc.opt.Debug()
	ignores := dbc.opt.Ignores()

	var err error
	if _, ok := ignores[query.ClientTypeBootnode]; !ok {
		dbc.cs[query.ClientTypeBootnode], err = openClient[bootnode.Client](path, dbc.opt)
		if err != nil {
			return nil, merr.Wrap(err, "open bootnode failed")
		}
	}
	if _, ok := ignores[query.ClientTypeKernel]; !ok {
		dbc.cs[query.ClientTypeKernel], err = openClient[kernel.Client](path, dbc.opt)
		if err != nil {
			return nil, merr.Wrap(err, "open kernel failed")
		}
	}
	if _, ok := ignores[query.ClientTypeUpgrade]; !ok {
		dbc.cs[query.ClientTypeUpgrade], err = openClient[upgrade.Client](path, dbc.opt)
		if err != nil {
			return nil, merr.Wrap(err, "open upgrade failed")
		}
	}
	if _, ok := ignores[query.ClientTypeMedia]; !ok {
		dbc.cs[query.ClientTypeMedia], err = openClient[media.Client](path, dbc.opt)
		if err != nil {
			return nil, merr.Wrap(err, "open media failed")
		}
	}
	return &dbc, nil
}

func (d *DBC) BootNode() *Client[bootnode.Client] {
	return d.cs[query.ClientTypeBootnode].(*Client[bootnode.Client])
}

func (d *DBC) Kernel() *Client[kernel.Client] {
	return d.cs[query.ClientTypeKernel].(*Client[kernel.Client])
}

func (d *DBC) Upgrade() *Client[upgrade.Client] {
	return d.cs[query.ClientTypeUpgrade].(*Client[upgrade.Client])
}

func (d *DBC) Media() *Client[media.Client] {
	return d.cs[query.ClientTypeMedia].(*Client[media.Client])
}

func (d *DBC) Close() error {
	for i := range d.cs {
		if d.cs[i] == nil {
			continue
		}
		if err := d.cs[i].Close(); err != nil {
			return err
		}
	}
	return nil
}

func (d *DBC) Client(p query.ClientType) (any, bool) {
	if p >= query.ClientTypeMax {
		return nil, false
	}
	return d.cs[p], true
}
