package handler

import (
	"julo-mini-wallet/util"
	"julo-mini-wallet/wallet"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type (
	restHandler struct {
		walletUsecase wallet.Usecase
	}
)

func RegisterRoute(router *httprouter.Router, walletUC wallet.Usecase) {
	r := &restHandler{
		walletUsecase: walletUC,
	}
	router.POST("/api/v1/init", r.CreateNewAccount)

}

func (rh *restHandler) CreateNewAccount(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	responseObj := util.NewRestResponse()

	var userID = r.FormValue("customer_xid")
	tokenResp, err := rh.walletUsecase.CreateAccount(userID)
	if err != nil {
		responseObj.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	responseObj.Data = tokenResp
	responseObj.Status = "success"
	responseObj.Write(w, http.StatusOK)
	return
}
