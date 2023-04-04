package models

type RespostaOK struct {
	Rmessage        string          `json:"message"`
	Rcode           int             `json:"code"`
	RautorizationId int             `json:"autorizationId"`
	Rbalance        Respostabalance `json:"balance"`
}

type RespostaCompleta struct {
	Rmessage string          `json:"message"`
	Rbalance Respostabalance `json:"balance"`
	Rcode    int             `json:"code"`
}

type RespostaSimples struct {
	Rmessage string `json:"message"`
	Rcode    int    `json:"code"`
}

type Respostabalance struct {
	Ramount        int `json:"amount"`
	Rcurrency_code int `json:"currency_code"`
}
