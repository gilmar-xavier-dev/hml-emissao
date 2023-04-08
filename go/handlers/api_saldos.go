package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"pagai/teste12/go/models"

	"github.com/go-playground/validator/v10"
)

func CreateQuery(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var saldo models.Query
	json.Unmarshal(body, &saldo)

	a, err := json.Marshal(saldo)

	//	fmt.Println(&compras)

	db, err := initDB()
	if err != nil {
		var resposta models.RespostaSimples
		resposta.Rmessage = "Sistema indisponível. Erro ao acessar a base de dados"
		resposta.Rcode = 900
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(w).Encode(&resposta)
		return
	}
	defer db.Close()

	validate := validator.New()
	errval := validate.Struct(saldo)
	//	fmt.Println("[", errval, "]")
	//	string_uuid := (uuid.New()).String()

	if errval != nil {
		var resposta models.RespostaSimples
		b, err := json.Marshal(resposta)
		sqlStatement := `INSERT INTO wallet_trans (transaction_id, purchase_id, account_id, request_body, response_body) VALUES ($1,$2,$3,$4,$5)`
		res, err := db.Exec(sqlStatement, saldo.QueryId, "", saldo.AccountId, string(a), string(b))
		if err != nil {
			var resposta models.RespostaSimples
			resposta.Rmessage = "Sistema indisponível. Erro ao acessar a base de dados"
			resposta.Rcode = 900
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusServiceUnavailable)
			json.NewEncoder(w).Encode(&resposta)
			return
		}
		_, err = res.RowsAffected()
		if err != nil {
			panic(err)
		}

		resposta.Rmessage = fmt.Sprintf("Requisição inválida - %s", errval)
		resposta.Rcode = 999
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&resposta)

		return
	}

	var Actualbalance int64
	//	var OldAmount int64
	var Status int
	//	var Lasttransid int
	sqlStatement := `SELECT amount, status FROM wallet_balance WHERE account_id=$1;`
	row := db.QueryRow(sqlStatement, saldo.AccountId)
	err = row.Scan(&Actualbalance, &Status)
	switch err {
	case sql.ErrNoRows:
		{
			var resposta models.RespostaSimples
			b, err := json.Marshal(resposta)
			sqlStatement := `INSERT INTO wallet_trans (transaction_id, purchase_id, account_id, request_body, response_body) VALUES ($1,$2,$3,$4,$5)`
			res, err := db.Exec(sqlStatement, saldo.QueryId, "", saldo.AccountId, string(a), string(b))
			if err != nil {
				var resposta models.RespostaSimples
				resposta.Rmessage = "Sistema indisponível. Erro ao acessar a base de dados"
				resposta.Rcode = 900
				w.Header().Add("Content-Type", "application/json")
				w.WriteHeader(http.StatusServiceUnavailable)
				json.NewEncoder(w).Encode(&resposta)
				return
			}
			_, err = res.RowsAffected()
			if err != nil {
				var resposta models.RespostaSimples
				resposta.Rmessage = "Sistema indisponível. Erro ao acessar a base de dados"
				resposta.Rcode = 900
				w.Header().Add("Content-Type", "application/json")
				w.WriteHeader(http.StatusServiceUnavailable)
				json.NewEncoder(w).Encode(&resposta)
				return
			}

			resposta.Rmessage = "Conta não encontrada"
			resposta.Rcode = 111
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(&resposta)

			return
		}
	}

	var resposta models.RespostaTaxa
	b, err := json.Marshal(resposta)
	//	fmt.Println("STRUCT:", string(b))

	sqlStatement = `INSERT INTO wallet_trans (transaction_id, purchase_id, account_id, request_body, response_body) VALUES ($1,$2,$3,$4,$5)`
	res, err := db.Exec(sqlStatement, saldo.QueryId, "", saldo.AccountId, string(a), string(b))

	if err != nil {
		var resposta models.RespostaSimples
		resposta.Rmessage = "Sistema indisponível. Erro ao acessar a base de dados"
		resposta.Rcode = 900
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(w).Encode(&resposta)
		return
	}
	_, err = res.RowsAffected()
	if err != nil {
		var resposta models.RespostaSimples
		resposta.Rmessage = "Sistema indisponível. Erro ao acessar a base de dados"
		resposta.Rcode = 900
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(w).Encode(&resposta)
		return
	}

	resposta.Rbalance.Ramount = int(Actualbalance)
	resposta.Rbalance.Rcurrency_code = 986
	resposta.Rfee.Ramount = 0
	resposta.Rfee.Rcurrency_code = 986
	resposta.RautorizationId = 123456
	resposta.Rmessage = "Operacao realizada com sucesso"
	resposta.Rcode = 0
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&resposta)

	return

}

func CreateWithdrawalQuery(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "CreateWithdrawalQuery: [%s]", r.Body)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
