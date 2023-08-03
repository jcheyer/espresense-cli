package espresense

type Device struct {
	ID       string  `json:"id"`
	IDType   int     `json:"idType"`
	Rssi1M   int     `json:"rssi@1m"`
	Rssi     int     `json:"rssi"`
	Raw      float64 `json:"raw"`
	Distance float64 `json:"distance"`
	Speed    float64 `json:"speed"`
	Mac      string  `json:"mac"`
	Interval int     `json:"interval"`
	Name     string  `json:"name,omitempty"`
	Disc     string  `json:"disc,omitempty"`
}

type DevicesResponse struct {
	Room    string    `json:"room"`
	Devices []*Device `json:"devices"`
}
