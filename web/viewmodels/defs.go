package viewmodels

type Response struct {
	Code   int         `json:"code"`
	Info   string      `json:"info"`
	Result interface{} `json:"result"`
	Err    interface{} `json:"err"`
}
