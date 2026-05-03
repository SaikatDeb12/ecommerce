package handler

import (
	"net/http"

	"github.com/SaikatDeb12/ecommerce/internal/models"
	"github.com/SaikatDeb12/ecommerce/internal/utils"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	var req models.SignUpModel
	if err := utils.ParseBody(r.Body, &req); err != nil {
		utils.ResponseError(w, http.StatusBadRequest, err, "invalid payload")
		return
	}

	if err := utils.ValidateStruct(r.Body); err != nil {
		utils.ResponseError(w, http.StatusBadRequest, err, "payload validation error")
		return
	}

	// isEmailExists, err := dbhelper.CheckEmailALreadyExists(req.Email)
	// if err != nil {
	// 	utils.ResponseError(w, http.StatusUnauthorized, err, "email already exists")
	// 	return
	// }
}

func SignIn(w http.ResponseWriter, r *http.Request) {
}
