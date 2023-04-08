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

func CreateChargeback(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var estorno models.Chargeback
	json.Unmarshal(body, &estorno)

	a, err := json.Marshal(estorno)

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
	errval := validate.Struct(estorno)
	//	fmt.Println("[", errval, "]")
	//	string_uuid := (uuid.New()).String()

	if errval != nil {
		var resposta models.RespostaSimples
		b, err := json.Marshal(resposta)
		sqlStatement := `INSERT INTO wallet_trans (transaction_id, purchase_id, account_id, request_body, response_body) VALUES ($1,$2,$3,$4,$5)`
		res, err := db.Exec(sqlStatement, estorno.ChargebackId, estorno.OriginalPurchaseId, estorno.AccountId, string(a), string(b))
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
	var OldAmount int64
	var Status int
	var Lasttransid int
	sqlStatement := `SELECT amount, status FROM wallet_balance WHERE account_id=$1;`
	row := db.QueryRow(sqlStatement, estorno.AccountId)
	err = row.Scan(&Actualbalance, &Status)
	switch err {
	case sql.ErrNoRows:
		{
			var resposta models.RespostaSimples
			b, err := json.Marshal(resposta)
			sqlStatement := `INSERT INTO wallet_trans (transaction_id, purchase_id, account_id, request_body, response_body) VALUES ($1,$2,$3,$4,$5)`
			res, err := db.Exec(sqlStatement, estorno.ChargebackId, estorno.OriginalPurchaseId, estorno.AccountId, string(a), string(b))
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

	fmt.Println("SITUACAO:", Actualbalance, estorno.OriginalPurchaseId, estorno.AccountId)

	sqlStatement = `SELECT amount, status FROM wallet_purchase WHERE purchase_id = $1 and account_id = $2;`
	row = db.QueryRow(sqlStatement, estorno.OriginalPurchaseId, estorno.AccountId)
	err = row.Scan(&OldAmount, &Status)

	if err == sql.ErrNoRows {
		var resposta models.RespostaSimples
		b, err := json.Marshal(resposta)
		sqlStatement := `INSERT INTO wallet_trans (transaction_id, purchase_id, account_id, request_body, response_body) VALUES ($1,$2,$3,$4,$5)`
		res, err := db.Exec(sqlStatement, estorno.ChargebackId, estorno.OriginalPurchaseId, estorno.AccountId, string(a), string(b))
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

		resposta.Rmessage = "Compra não encontrada"
		resposta.Rcode = 111
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(&resposta)

		return
	}

	if Status == 1 {
		var resposta models.RespostaSimples
		b, err := json.Marshal(resposta)
		sqlStatement := `INSERT INTO wallet_trans (transaction_id, purchase_id, account_id, request_body, response_body) VALUES ($1,$2,$3,$4,$5)`
		res, err := db.Exec(sqlStatement, estorno.ChargebackId, estorno.OriginalPurchaseId, estorno.AccountId, string(a), string(b))
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

		resposta.Rmessage = "Compra já cancelada"
		resposta.Rcode = 144
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusPreconditionFailed)
		json.NewEncoder(w).Encode(&resposta)

		return
	}

	if estorno.ChargebackAmount.Amount != OldAmount {
		fmt.Println("VALORES DIFERENTES:", estorno.ChargebackAmount.Amount, OldAmount)
		var resposta models.RespostaCompleta
		b, err := json.Marshal(resposta)
		sqlStatement := `INSERT INTO wallet_trans (transaction_id, purchase_id, account_id, request_body, response_body) VALUES ($1,$2,$3,$4,$5)`
		res, err := db.Exec(sqlStatement, estorno.ChargebackId, estorno.OriginalPurchaseId, estorno.AccountId, string(a), string(b))
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

		resposta.Rbalance.Ramount = int(estorno.OriginalAmount.Amount)
		resposta.Rbalance.Rcurrency_code = 986
		resposta.Rmessage = "Valor da compra diferente do cancelamento"
		resposta.Rcode = 530
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusPreconditionFailed)
		json.NewEncoder(w).Encode(&resposta)

		return
	}

	NewAmount := Actualbalance + estorno.ChargebackAmount.Amount
	NewTransid := Lasttransid + 1
	if NewTransid > 999998 {
		NewTransid = 1
	}

	fmt.Println("NOVO SALDO:", "CANC:", estorno.ChargebackAmount.Amount, "ANTER:", Actualbalance, "NOVO:", NewAmount)

	sqlStatement = `UPDATE wallet_balance SET amount = $2, status = 0, remarks = $3 WHERE account_id = $1;`
	res, err := db.Exec(sqlStatement, estorno.AccountId, NewAmount, estorno.OriginalPurchaseId)
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

	sqlStatement = `UPDATE wallet_purchase SET status = 1, balance = $3 WHERE account_id = $1 and purchase_id = $2;`
	res, err = db.Exec(sqlStatement, estorno.AccountId, estorno.OriginalPurchaseId, NewAmount)
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
	//	string_uuid := (uuid.New()).String()

	var resposta models.RespostaOK
	b, err := json.Marshal(resposta)
	//	fmt.Println("STRUCT:", string(b))

	sqlStatement = `INSERT INTO wallet_trans (transaction_id, purchase_id, account_id, request_body, response_body) VALUES ($1,$2,$3,$4,$5)`
	res, err = db.Exec(sqlStatement, estorno.ChargebackId, estorno.OriginalPurchaseId, estorno.AccountId, string(a), string(b))

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

	resposta.Rbalance.Ramount = int(estorno.ChargebackAmount.Amount)
	resposta.Rbalance.Rcurrency_code = 986
	resposta.RautorizationId = NewTransid
	resposta.Rmessage = "Operacao realizada com sucesso"
	resposta.Rcode = 0
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&resposta)

	return

}

func CreateChargebackCancellation(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "CreateChargebackCancellation: [%s]", r.Body)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

//Gilmar
