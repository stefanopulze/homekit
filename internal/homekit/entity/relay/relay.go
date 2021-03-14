package relay

import (
	"encoding/json"
	"fmt"
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/service"
	"github.com/sirupsen/logrus"
	"homekit-server/internal/homekit/entity"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type relay struct {
	*accessory.Accessory
	Switch  *service.Switch
	Name    string
	options Options
}

func New(entity entity.BaseEntity) *relay {
	// workaround
	opBytes, err := json.Marshal(entity.Options)
	if err != nil {
		logrus.Error(err)
	}

	r := relay{
		Name: entity.Name,
	}
	json.Unmarshal(opBytes, &r.options)

	info := accessory.Info{Name: entity.Name}

	if len(entity.SerialNumber) > 0 {
		info.SerialNumber = entity.SerialNumber
	}

	if len(entity.FirmwareVersion) > 0 {
		info.FirmwareRevision = entity.FirmwareVersion
	}

	r.Accessory = accessory.New(info, accessory.TypeSwitch)
	r.Switch = service.NewSwitch()
	r.AddService(r.Switch.Service)

	r.Switch.On.OnValueRemoteGet(func() bool {
		status, err := r.getStatus()

		if err != nil {
			return false
		}

		return status.Ison
	})

	r.Switch.On.OnValueRemoteUpdate(func(on bool) {
		uri := buildUri(r.options)
		turn := "on"

		if !on {
			turn = "off"
		}

		data := url.Values{}
		data.Set("turn", turn)

		logrus.Debugf("Turning %v sensor %s", on, entity.Name)

		resp, err := http.Post(uri, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
		if err != nil {
			logrus.Warn(err)
			return
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logrus.Warn(err)
			return
		}

		var response statusResponse
		if err := json.Unmarshal(body, &response); err != nil {
			logrus.Warn(err)
			return
		}

		logrus.Warn(fmt.Sprintf("Relay %+v", response))

		if r.options.ReloadTimeout > 0 {
			go func() {
				time.Sleep(time.Duration(r.options.ReloadTimeout) * time.Millisecond)
				status, err := r.getStatus()
				if err != nil {
					r.Switch.On.SetValue(false)
				}

				r.Switch.On.SetValue(status.Ison)
			}()
		}
	})

	return &r
}

func buildUri(options Options) string {
	return fmt.Sprintf("http://%s/relay/%d", options.IP, options.RelayId)
}

func (r *relay) getStatus() (*statusResponse, error) {
	uri := buildUri(r.options)
	logrus.Debugf("Requesting status: %s", uri)
	resp, err := http.Get(uri)
	if err != nil {
		logrus.Warn(err)
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Warn(err)
		return nil, err
	}

	var response statusResponse
	if err := json.Unmarshal(body, &response); err != nil {
		logrus.Warnf("Error reading %s value: %v", r.Name, err)
		return nil, err
	}

	return &response, nil
}
