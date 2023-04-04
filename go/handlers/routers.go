package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello World! [%s]", r.Body)
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	Route{
		"CreatePurchase",
		strings.ToUpper("Post"),
		"/purchases",
		CreatePurchase,
	},

	Route{
		"CreatePurchaseCancellation",
		strings.ToUpper("Post"),
		"/purchases/cancel",
		CreatePurchaseCancellation,
	},

	Route{
		"CreateChargeback",
		strings.ToUpper("Post"),
		"/chargebacks",
		CreateChargeback,
	},

	Route{
		"CreateChargebackCancellation",
		strings.ToUpper("Post"),
		"/chargebacks/cancel",
		CreateChargebackCancellation,
	},

	Route{
		"CreateQuery",
		strings.ToUpper("Post"),
		"/queries",
		CreateQuery,
	},

	Route{
		"CreateWithdrawalQuery",
		strings.ToUpper("Post"),
		"/withdrawalQueries",
		CreateWithdrawalQuery,
	},

	Route{
		"CreateWithdrawal",
		strings.ToUpper("Post"),
		"/withdrawals",
		CreateWithdrawal,
	},

	Route{
		"CreateWithdrawalCancellation",
		strings.ToUpper("Post"),
		"/withdrawals/cancel",
		CreateWithdrawalCancellation,
	},

	Route{
		"GetStatus",
		strings.ToUpper("Get"),
		"/status",
		GetStatus,
	},

	Route{
		"CreateTransfer",
		strings.ToUpper("Post"),
		"/transfers",
		CreateTransfer,
	},

	Route{
		"CreateTransferCancellation",
		strings.ToUpper("Post"),
		"/transfers/cancel",
		CreateTransferCancellation,
	},
}
