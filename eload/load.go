package eload

import (
	"errors"
	"github.com/cornelk/hashmap"
	"github.com/tikafog/of"
	"github.com/tikafog/of/dbc"
	"github.com/tikafog/of/option"
	"sync"
)

type loader struct {
	rw      sync.RWMutex
	names   *hashmap.Map[uint64, of.Name]
	modules *hashmap.Map[uint64, of.Module]
	dbc     *dbc.DBC
	event   of.Event
	api     of.API
	option  option.Option
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

func newLoader() Loader {
	return &loader{
		names:   hashmap.New[uint64, of.Name](),
		modules: hashmap.New[uint64, of.Module](),
	}
}

func (i *loader) AddSource(v any) error {
	switch d := v.(type) {
	case *dbc.DBC:
		i.dbc = d
	case of.Event:
		i.event = d
	case of.API:
		i.api = d
	case option.Option:
		i.option = d
	default:
		return errors.New("unknown add type")
	}
	return nil
}

func (i *loader) Inject() error {
	var err error
	if i.option != nil {
		i.modules.Range(func(idx uint64, module of.Module) bool {
			err = i.InjectOption(module)
			if err != nil {
				return false
			}
			return true
		})
	}
	if i.dbc != nil {
		i.modules.Range(func(idx uint64, module of.Module) bool {
			err = i.InjectDBC(module)
			if err != nil {
				return false
			}
			return true
		})
	}
	if i.event != nil {
		i.modules.Range(func(idx uint64, module of.Module) bool {
			err = i.InjectEvent(module)
			if err != nil {
				return false
			}
			return true
		})
	}
	if i.api != nil {
		i.modules.Range(func(idx uint64, module of.Module) bool {
			err = i.InjectAPI(module)
			if err != nil {
				return false
			}
			return true
		})
	}
	return nil
}

func (i *loader) InjectDBC(module of.Module) error {
	if v, ok := module.(dbc.DatabaseRegister); ok {
		err := v.RegisterDBC(i.dbc)
		if err != nil {
			return err
		}
	}
	return nil
}

func (i *loader) InjectEvent(module of.Module) error {
	if v, ok := module.(of.EventRegister); ok {
		err := v.RegisterEvent(i.event)
		if err != nil {
			return err
		}
	}
	return nil
}

func (i *loader) InjectAPI(module of.Module) error {
	if v, ok := module.(of.APIRegister); ok {
		err := v.RegisterAPI(i.api)
		if err != nil {
			return err
		}
	}
	return nil
}

func (i *loader) InjectOption(module of.Module) error {
	if v, ok := module.(option.Register); ok {
		err := v.RegisterOption(i.option)
		if err != nil {
			return err
		}
	}
	return nil
}
