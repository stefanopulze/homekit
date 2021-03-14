package relay

type Options struct {
	IP            string `json:"ip"`
	RelayId       uint8  `json:"relayId"`
	ReloadTimeout uint16 `json:"reloadTimeout"`
}
