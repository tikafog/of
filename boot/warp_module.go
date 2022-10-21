package boot

import (
	"context"
	"encoding/json"

	"github.com/tikafog/of"
)

type wrapModule struct {
	of.Module
	event of.EventRegister
	api   of.APIRegister
	core  of.CoreLoader
	run   of.ModuleRunner
	cfg   of.ConfigLoader
}

func (w wrapModule) LoadConfig(message json.RawMessage) (json.RawMessage, error) {
	if w.cfg != nil {
		return w.cfg.LoadConfig(message)
	}
	return nil, nil
}

func (w wrapModule) Init() error {
	if w.run != nil {
		return w.run.Init()
	}
	return nil
}

func (w wrapModule) Run(ctx context.Context) error {
	if w.run != nil {
		return w.run.Run(ctx)
	}
	return nil
}

func (w wrapModule) Destroy() {
	if w.run != nil {
		w.run.Destroy()
	}
}

func (w wrapModule) PreloadCore(core of.Core) error {
	if w.core != nil {
		return w.core.PreloadCore(core)
	}
	return nil
}

func (w wrapModule) RegisterAPI(api of.API) error {
	if w.api != nil {
		return w.api.RegisterAPI(api)
	}
	return nil
}

func (w wrapModule) RegisterEvent(event of.Event) error {
	if w.event != nil {
		return w.event.RegisterEvent(event)
	}
	return nil
}

func WarpLoader(m of.Module) of.ModuleStarter {
	wm := &wrapModule{Module: m}

	event, ok := m.(of.EventRegister)
	if ok {
		wm.event = event
	}
	api, ok := m.(of.APIRegister)
	if ok {
		wm.api = api
	}
	core, ok := m.(of.CoreLoader)
	if ok {
		wm.core = core
	}
	run, ok := m.(of.ModuleRunner)
	if ok {
		wm.run = run
	}
	cfg, ok := m.(of.ConfigLoader)
	if ok {
		wm.cfg = cfg
	}

	return wm
}
