package handler

import (
	"julo-mini-wallet/util"
	"julo-mini-wallet/wallet"
	"net/http"
	"strconv"
	"strings"

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
	router.POST("/api/v1/wallet", r.EnableWallet)
	router.POST("/api/v1/wallet/deposits", r.DepositMoney)
	router.POST("/api/v1/wallet/withdrawals", r.WithdrawMoney)
	router.GET("/api/v1/wallet", r.GetWallet)
	router.PATCH("/api/v1/wallet", r.DisableWallet)
	router.GET("/api/v1/wallet/transactions", r.GetWalletTransactions)

}

func (rh *restHandler) WithdrawMoney(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	responseObj := util.NewRestResponse()
	userID, errMessage := validateAndGetUserIDFromAuthHeader(r)
	if errMessage != "" {
		responseObj.WriteError(w, http.StatusUnauthorized, errMessage)
		return
	}
	amountString := r.FormValue("amount")
	amount, err := strconv.ParseFloat(amountString, 64)
	if err != nil {
		responseObj.WriteError(w, http.StatusBadRequest, "Invalid value : amount is not a number")
		return
	}
	refID := r.FormValue("reference_id")

	withdrawData, err := rh.walletUsecase.Withdraw(userID, amount, refID)
	if err != nil {
		responseObj.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	responseObj.Data = withdrawData
	responseObj.Status = "success"
	responseObj.Write(w, http.StatusOK)
	return
}

func (rh *restHandler) DepositMoney(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	responseObj := util.NewRestResponse()
	userID, errMessage := validateAndGetUserIDFromAuthHeader(r)
	if errMessage != "" {
		responseObj.WriteError(w, http.StatusUnauthorized, errMessage)
		return
	}
	amountString := r.FormValue("amount")
	amount, err := strconv.ParseFloat(amountString, 64)
	if err != nil {
		responseObj.WriteError(w, http.StatusBadRequest, "Invalid value : amount is not a number")
		return
	}
	refID := r.FormValue("reference_id")

	depositData, err := rh.walletUsecase.Deposit(userID, amount, refID)
	if err != nil {
		responseObj.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	responseObj.Data = depositData
	responseObj.Status = "success"
	responseObj.Write(w, http.StatusOK)
	return
}

func (rh *restHandler) GetWalletTransactions(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	responseObj := util.NewRestResponse()
	userID, errMessage := validateAndGetUserIDFromAuthHeader(r)
	if errMessage != "" {
		responseObj.WriteError(w, http.StatusUnauthorized, errMessage)
		return
	}

	transactionsData, err := rh.walletUsecase.GetWalletTransactions(userID)
	if err != nil {
		responseObj.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	responseObj.Data = transactionsData
	responseObj.Status = "success"
	responseObj.Write(w, http.StatusOK)
	return
}

func (rh *restHandler) GetWallet(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	responseObj := util.NewRestResponse()
	userID, errMessage := validateAndGetUserIDFromAuthHeader(r)
	if errMessage != "" {
		responseObj.WriteError(w, http.StatusUnauthorized, errMessage)
		return
	}

	tokenResp, err := rh.walletUsecase.GetWalletBalance(userID)
	if err != nil {
		responseObj.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	responseObj.Data = tokenResp
	responseObj.Status = "success"
	responseObj.Write(w, http.StatusOK)
	return
}

func (rh *restHandler) DisableWallet(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	responseObj := util.NewRestResponse()
	userID, errMessage := validateAndGetUserIDFromAuthHeader(r)
	if errMessage != "" {
		responseObj.WriteError(w, http.StatusUnauthorized, errMessage)
		return
	}
	var isDisabledInString = r.FormValue("is_disabled")
	if isDisabledInString == "" {
		responseObj.WriteError(w, http.StatusBadRequest, "Missing data for required field : is_disabled ")
		return
	}
	isDisabled, err := strconv.ParseBool(isDisabledInString)
	if err != nil {
		responseObj.WriteError(w, http.StatusBadRequest, "Invalid value : is_disabled is not boolean")
		return
	}
	if !isDisabled {
		responseObj.WriteError(w, http.StatusOK, "No Action : is_disabled is false")
		return
	}
	tokenResp, err := rh.walletUsecase.DisableWallet(userID)
	if err != nil {
		responseObj.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	responseObj.Data = tokenResp
	responseObj.Status = "success"
	responseObj.Write(w, http.StatusOK)
	return
}

func (rh *restHandler) EnableWallet(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	responseObj := util.NewRestResponse()
	userID, errMessage := validateAndGetUserIDFromAuthHeader(r)
	if errMessage != "" {
		responseObj.WriteError(w, http.StatusUnauthorized, errMessage)
		return
	}

	tokenResp, err := rh.walletUsecase.EnableWallet(userID)
	if err != nil {
		responseObj.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	responseObj.Data = tokenResp
	responseObj.Status = "success"
	responseObj.Write(w, http.StatusOK)
	return
}

func (rh *restHandler) CreateNewAccount(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	responseObj := util.NewRestResponse()

	var userID = r.FormValue("customer_xid")
	if userID == "" {
		responseObj.WriteError(w, http.StatusBadRequest, "Missing data for required field : customer_xid ")
		return
	}
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

func validateAndGetUserIDFromAuthHeader(r *http.Request) (userID string, errMessage string) {
	authorization := r.Header.Get("Authorization")
	if authorization == "" {
		errMessage = "Missing Authorization header"
		return
	}
	if !strings.HasPrefix(authorization, "Token") {
		errMessage = "Invalid Authorization header"
		return
	}
	authHeader := strings.Split(authorization, " ")
	if len(authHeader) < 2 {
		errMessage = "Invalid Authorization header"
		return
	}

	token := authHeader[1]
	userID, err := util.GetUserIDFromToken(token)
	if err != nil {
		errMessage = "Invalid Token"
		return
	}
	return
}
