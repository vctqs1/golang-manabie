package controllers

import (
	"net/http"
	"encoding/json"
	"context"

	"github.com/vctqs1/golang-manabie/pkg/services"
	u "github.com/vctqs1/golang-manabie/pkg/controllers/utils"
)

var BuyProducts = func(w http.ResponseWriter, r *http.Request) {
	resp := protov1.BuyProducts(context.Context, r.Body)
	u.Respond(w, resp)
}
