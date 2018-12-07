package viewmodels

type Response struct {
	Code   int         `json:"code"`
	Info   string      `json:"info"`
	Result interface{} `json:"result"`
	Err    interface{} `json:"err"`
}

type Star struct {
	Id     int    `json:"id"`
	NameZh string `json:"name_zh"`
	NameEn string `json:"name_en"`
	Avatar string `json:"avatar"`
}
