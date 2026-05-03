package handler

import (
	"net/http"

	"github.com/SaikatDeb12/ecommerce/internal/utils"
)

func CheckHealth(w http.ResponseWriter, r *http.Request) {
	utils.RespondJSON(w, 200, map[string]string{
		"status": "server is running",
	})
}
