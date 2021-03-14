package homekit

import (
	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	"github.com/sirupsen/logrus"
	"homekit-server/internal/homekit/entity"
	"homekit-server/internal/homekit/entity/relay"
	"log"
)

type server struct {
	opts     *ConfigOpts
	termFunc hc.TermFunc
}

func New(opts *ConfigOpts) *server {
	return &server{
		opts: opts,
	}
}

func (s *server) Start(entities entity.BaseEntities) {
	accessories := make([]*accessory.Accessory, 0)

	for _, e := range entities {
		logrus.Debugf("Entity: %+v", e)

		if e.Type == "relay" {
			ac := relay.New(e)
			if ac != nil {
				accessories = append(accessories, ac.Accessory)
			}
		}
	}

	//// create an accessory
	//info := accessory.Info{Name: s.opts.Name}
	//ac := accessory.NewSwitch(info)
	//ac.Switch.On.OnValueRemoteUpdate(func(open bool) {
	//	if !open {
	//		return
	//	}
	//
	//	http.Get("http://192.168.1.124/open")
	//
	//	go func() {
	//		time.Sleep(1 * time.Second)
	//		ac.Switch.On.SetValue(false)
	//	}()
	//})
	//
	//ac.Switch.On.OnValueRemoteGet(func() bool {
	//	return false
	//})

	info := accessory.Info{Name: s.opts.Name}
	bridge := accessory.NewBridge(info)

	// configure the ip transport
	config := hc.Config{
		Pin:         s.opts.Pin,
		StoragePath: s.opts.StoragePath,
	}
	t, err := hc.NewIPTransport(config, bridge.Accessory, accessories...)
	if err != nil {
		log.Panic(err)
	}

	s.termFunc = func() {
		<-t.Stop()
	}

	//hc.OnTermination(func() {
	//	<-t.Stop()
	//})

	t.Start()
}

func (s *server) Shutdown() {
	logrus.Debug("Homekit server shutting down")
	s.termFunc()
	logrus.Info("Homekit server shutdown")
}
