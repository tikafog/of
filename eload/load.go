package eload

import (
	"encoding/json"
	"errors"
	"github.com/cornelk/hashmap"
	"github.com/tikafog/of"
	"github.com/tikafog/of/dbc"
	"github.com/tikafog/of/option"
	"sync"
)

type injectFunc func(module of.Module) error

// RegisterAny ...
type RegisterAny interface {
	RegisterAny(v any) error
}

type loader struct {
	rw      sync.RWMutex
	names   *hashmap.Map[uint64, of.Name]
	modules *hashmap.Map[uint64, of.Module]
	//dbc     *dbc.DBC
	//event   of.Event
	//api     of.API
	//option  option.Option
}

// LoadConfig ...
// @receiver *loader
// @param json.RawMessage
// @return json.RawMessage
// @return error
func (i *loader) LoadConfig(message json.RawMessage) (json.RawMessage, error) {
	cfgs := make(map[string]json.RawMessage, i.modules.Len())
	err := json.Unmarshal(message, &cfgs)
	if err != nil {
		return nil, err
	}
	i.modules.Range(func(u uint64, m of.Module) bool {
		cfg, ok := cfgs[m.Name().String()]
		if !ok {
			return true
		}
		mm, ok := m.(ConfigLoader)
		if !ok {
			return true
		}
		cfg, err = mm.LoadConfig(cfg)
		if err != nil {
			return true
		}
		cfgs[m.Name().String()] = cfg
		return true
	})
	return json.MarshalIndent(cfgs, "", " ")
}

func newLoader() Loader {
	return &loader{
		names:   hashmap.New[uint64, of.Name](),
		modules: hashmap.New[uint64, of.Module](),
	}
}

func (i *loader) Names() []of.Name {
	rets := make([]of.Name, 0, i.names.Len())
	i.names.Range(func(k uint64, n of.Name) bool {
		rets = append(rets, n)
		return true
	})
	return rets
}

func (i *loader) WithOption(o option.Option) of.ModuleStarter {
	return nil
}

func (i *loader) Register(id uint64, name of.Name) error {
	if !i.names.Insert(id, name) {
		return errors.New("module is already loaded")
	}
	return nil
}

func (i *loader) Add(id uint64, m of.Module) error {
	err := i.Register(id, m.Name())
	if err != nil {
		return err
	}
	if !i.modules.Insert(id, m) {
		return errors.New("module is already loaded")
	}
	return nil
}

func (i *loader) Get(id uint64) (of.Module, bool) {
	return i.modules.Get(id)
}

func (i *loader) GetByName(name of.Name) (of.Module, bool) {
	return i.modules.Get(name.ID())
}

func (i *loader) Inject(v any) error {
	var fn injectFunc
	switch d := v.(type) {
	case *dbc.DBC:
		fn = i.injectDBC(d)
	case of.Event:
		fn = i.injectEvent(d)
	case of.API:
		fn = i.injectAPI(d)
	case option.Option:
		fn = i.injectOption(d)
	default:
		fn = i.injectAny(d)
	}

	if fn != nil {
		i.modules.Range(func(_ uint64, module of.Module) bool {
			err := fn(module)
			if err != nil {
				return false
			}
			return true
		})
	}
	return nil
}

func (i *loader) injectDBC(d *dbc.DBC) injectFunc {
	return func(module of.Module) error {
		if v, ok := module.(DatabaseRegister); ok {
			err := v.RegisterDBC(d)
			if err != nil {
				return err
			}
		}
		return nil
	}
}

func (i *loader) injectEvent(event of.Event) injectFunc {
	return func(module of.Module) error {
		if v, ok := module.(EventRegister); ok {
			err := v.RegisterEvent(event)
			if err != nil {
				return err
			}
		}
		return nil
	}
}

func (i *loader) injectAPI(api of.API) injectFunc {
	return func(module of.Module) error {
		if v, ok := module.(APIRegister); ok {
			err := v.RegisterAPI(api)
			if err != nil {
				return err
			}
		}
		return nil
	}
}

func (i *loader) injectOption(op option.Option) injectFunc {
	return func(module of.Module) error {
		if v, ok := module.(OptionRegister); ok {
			err := v.RegisterOption(op)
			if err != nil {
				return err
			}
		}
		return nil
	}
}

func (i *loader) injectAny(d any) injectFunc {
	return func(module of.Module) error {
		if v, ok := module.(RegisterAny); ok {
			err := v.RegisterAny(d)
			if err != nil {
				return err
			}
		}
		return nil
	}
}
