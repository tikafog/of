package dbc

import (
	"io"

	"github.com/tikafog/of/dbc/bootnode"
	"github.com/tikafog/of/dbc/kernel"
	"github.com/tikafog/of/dbc/media"
	"github.com/tikafog/of/dbc/upgrade"
	"github.com/tikafog/of/feature/source"
)

type DBC struct {
	opt *Option
	cs  [source.TypeMax]io.Closer
}

func Open(path string, opts ...Opts) (*DBC, error) {
	var dbc DBC
	dbc.opt = parseOption(opts)
	debug = dbc.opt.Debug()
	ignores := dbc.opt.Ignores()

	var err error
	if _, ok := ignores[source.TypeBootnode]; !ok {
		dbc.cs[source.TypeBootnode], err = openClient[bootnode.Client](path, dbc.opt)
		if err != nil {
			return nil, merr.Wrap(err, "open bootnode failed")
		}
	}
	if _, ok := ignores[source.TypeKernel]; !ok {
		dbc.cs[source.TypeKernel], err = openClient[kernel.Client](path, dbc.opt)
		if err != nil {
			return nil, merr.Wrap(err, "open kernel failed")
		}
	}
	if _, ok := ignores[source.TypeUpgrade]; !ok {
		dbc.cs[source.TypeUpgrade], err = openClient[upgrade.Client](path, dbc.opt)
		if err != nil {
			return nil, merr.Wrap(err, "open upgrade failed")
		}
	}
	if _, ok := ignores[source.TypeMedia]; !ok {
		dbc.cs[source.TypeMedia], err = openClient[media.Client](path, dbc.opt)
		if err != nil {
			return nil, merr.Wrap(err, "open media failed")
		}
	}
	return &dbc, nil
}

func (d *DBC) BootNode() *Client[bootnode.Client] {
	return d.cs[source.TypeBootnode].(*Client[bootnode.Client])
}

func (d *DBC) Kernel() *Client[kernel.Client] {
	return d.cs[source.TypeKernel].(*Client[kernel.Client])
}

func (d *DBC) Upgrade() *Client[upgrade.Client] {
	return d.cs[source.TypeUpgrade].(*Client[upgrade.Client])
}

func (d *DBC) Media() *Client[media.Client] {
	return d.cs[source.TypeMedia].(*Client[media.Client])
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

func (d *DBC) Client(p source.Type) (any, bool) {
	if p >= source.TypeMax {
		return nil, false
	}
	return d.cs[p], true
}
