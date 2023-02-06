package handler

import (
	"crypto/tls"
	"encoding/json"
	"go-ms/apis"
	"io/ioutil"
	"net/http"
)

type ShoppingCartHandler struct {
}

func (h ShoppingCartHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if http.MethodGet == r.Method {
		handleGet(w, r)
	} else if http.MethodPost == r.Method {
		handlePost(w, r)
	} else {
		http.Error(w, "Not Implemented", http.StatusNotImplemented)
		return
	}

	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	client.Get("")

}

func handlePost(w http.ResponseWriter, r *http.Request) {

	r.URL.Query().Get("")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {

	}
	cart := &apis.ShoppingCart{}
	err = json.Unmarshal(body, cart)

	if err != nil {

	}
	json.NewEncoder(w).Encode(cart)

}

func handleGet(w http.ResponseWriter, r *http.Request) {
	cart := &apis.ShoppingCart{
		CartItems: []string{
			"Iphone",
			"Samsung",
		},
	}

	//fmt.Fprintf(writer, "responding from server %+v", cart)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cart)
}
