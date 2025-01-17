//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by iocli, run 'iocli gen' to re-generate

package main

import (
	autowire "github.com/alibaba/ioc-golang/autowire"
	normal "github.com/alibaba/ioc-golang/autowire/normal"
	singleton "github.com/alibaba/ioc-golang/autowire/singleton"
	util "github.com/alibaba/ioc-golang/autowire/util"
)

func init() {
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &app_{}
		},
	})
	appStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &App{}
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
	}
	singleton.RegisterStructDescriptor(appStructDescriptor)
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &otherMux_{}
		},
	})
	otherMuxStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &OtherMux{}
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
	}
	singleton.RegisterStructDescriptor(otherMuxStructDescriptor)
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &mux_{}
		},
	})
	muxStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &Mux{}
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
	}
	singleton.RegisterStructDescriptor(muxStructDescriptor)
}

type app_ struct {
}

type otherMux_ struct {
}

type mux_ struct {
}

type AppIOCInterface interface {
}

type OtherMuxIOCInterface interface {
}

type MuxIOCInterface interface {
}

var _appSDID string

func GetAppSingleton() (*App, error) {
	if _appSDID == "" {
		_appSDID = util.GetSDIDByStructPtr(new(App))
	}
	i, err := singleton.GetImpl(_appSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(*App)
	return impl, nil
}

func GetAppIOCInterfaceSingleton() (AppIOCInterface, error) {
	if _appSDID == "" {
		_appSDID = util.GetSDIDByStructPtr(new(App))
	}
	i, err := singleton.GetImplWithProxy(_appSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(AppIOCInterface)
	return impl, nil
}

var _otherMuxSDID string

func GetOtherMuxSingleton() (*OtherMux, error) {
	if _otherMuxSDID == "" {
		_otherMuxSDID = util.GetSDIDByStructPtr(new(OtherMux))
	}
	i, err := singleton.GetImpl(_otherMuxSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(*OtherMux)
	return impl, nil
}

func GetOtherMuxIOCInterfaceSingleton() (OtherMuxIOCInterface, error) {
	if _otherMuxSDID == "" {
		_otherMuxSDID = util.GetSDIDByStructPtr(new(OtherMux))
	}
	i, err := singleton.GetImplWithProxy(_otherMuxSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(OtherMuxIOCInterface)
	return impl, nil
}

var _muxSDID string

func GetMuxSingleton() (*Mux, error) {
	if _muxSDID == "" {
		_muxSDID = util.GetSDIDByStructPtr(new(Mux))
	}
	i, err := singleton.GetImpl(_muxSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(*Mux)
	return impl, nil
}

func GetMuxIOCInterfaceSingleton() (MuxIOCInterface, error) {
	if _muxSDID == "" {
		_muxSDID = util.GetSDIDByStructPtr(new(Mux))
	}
	i, err := singleton.GetImplWithProxy(_muxSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(MuxIOCInterface)
	return impl, nil
}
