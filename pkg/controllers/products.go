package controllers;

import (
	"net/http"

	"github.com/vctqs1/golang-manabie/pkg/services"
	u "github.com/vctqs1/golang-manabie/pkg/utils"
)

var Buy = func(w http.ResponseWriter, r *http.Request) {
	resp := protov1.BuyProduct(
		r.Context(),
		r.Body,
	)
	u.Respond(w, resp)
}
