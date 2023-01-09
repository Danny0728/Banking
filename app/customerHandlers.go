package app

import (
	"encoding/json"
	"net/http"

	"github.com/Danny0728/BankAPI/service"
	"github.com/gorilla/mux"
)

type CustomerHandlers struct {
	service service.CustomerService
}

// rest adapters should be dependent on primary port(service)CustomerService
func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	var status *string
	if len(values) > 0 {
		status = &values["status"][0]
	}
	customers, err := ch.service.GetAllCustomer(status)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customers)
	}

}

func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]
	customer, err := ch.service.GetCustomer(id)

	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customer)
	}

}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}

}

// func (ch *CustomerHandlers) addCustomer(w http.ResponseWriter, r *http.Request) {
// 	customers, _ := ch.service.AddCustomer()

// 	if r.Header.Get("Content-Type") == "application/xml" {
// 		w.Header().Add("Content-Type", "application/xml")
// 		xml.NewEncoder(w).Encode(customers)
// 	} else {
// 		w.Header().Add("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(customers)
// 	}
// }
