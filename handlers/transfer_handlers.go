package handlers

import (
	"GoPay/models"
	"GoPay/services"
	"GoPay/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

type TransferRequest struct {
	Value float64 `json:"value"`
	Payer int     `json:"payer"`
	Payee int     `json:"payee"`
}

func Transfer(w http.ResponseWriter, r *http.Request) {
	var transferReq TransferRequest
	_ = json.NewDecoder(r.Body).Decode(&transferReq)

	payer, payee, err := validateTransferRequest(transferReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = processTransfer(payer, payee, transferReq.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Transfer successful"})
}

func validateTransferRequest(req TransferRequest) (*models.User, *models.User, error) {
	var payer, payee *models.User

	for i, user := range models.Users {
		if user.ID == req.Payer {
			payer = &models.Users[i]
		}
		if user.ID == req.Payee {
			payee = &models.Users[i]
		}
	}

	if payer == nil || payee == nil {
		return nil, nil, fmt.Errorf("Invalid payer or payee")
	}

	if payer.IsMerchant {
		return nil, nil, fmt.Errorf("Merchants cannot initiate transfers")
	}

	if payer.Balance < req.Value {
		return nil, nil, fmt.Errorf("Insufficient balance")
	}

	return payer, payee, nil
}

func processTransfer(payer *models.User, payee *models.User, value float64) error {
	err := utils.RunInTransaction(func() error {
		if authorized, err := services.Authorize(); !authorized || err != nil {
			return fmt.Errorf("Authorization failed")
		}

		payer.Balance -= value

		payee.Balance += value

		if err := services.Notify(payer, "Successfully created a transaction"); err != nil {
			return fmt.Errorf("Notification failed")
		}
		if err := services.Notify(payee, "Successfully received a transaction"); err != nil {
			return fmt.Errorf("Notification failed")
		}

		return nil
	})

	return err
}
