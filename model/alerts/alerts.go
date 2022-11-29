package alerts

type GetAlertsModel struct {
	Jsonrpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  GetAlertsOpts `json:"params"`
	Auth    string        `json:"auth"`
	ID      float64       `json:"id"`
}
type GetAlertsOpts struct {
	Output    string `json:"output"`
	Actionids string `json:"actionids"`
}

type AlertsResult struct {
	Jsonrpc string              `json:"jsonrpc"`
	Result  []AlertsResultParam `json:"result"`
	ID      float64             `json:"id"`
}
type AlertsResultParam struct {
	Alertid       string `json:"alertid"`
	Actionid      string `json:"actionid"`
	Eventid       string `json:"eventid"`
	Userid        string `json:"userid"`
	Clock         string `json:"clock"`
	Mediatypeid   string `json:"mediatypeid"`
	Sendto        string `json:"sendto"`
	Subject       string `json:"subject"`
	Message       string `json:"message"`
	Status        string `json:"status"`
	Retries       string `json:"retries"`
	Error         string `json:"error"`
	EscStep       string `json:"esc_step"`
	Alerttype     string `json:"alerttype"`
	PEventid      string `json:"p_eventid"`
	Acknowledgeid string `json:"acknowledgeid"`
}
