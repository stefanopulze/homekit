package relay

type statusResponse struct {
	Ison   bool   `json:"ison"`
	Source string `json:"source"`
}
