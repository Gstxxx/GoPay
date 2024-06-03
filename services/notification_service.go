package services

import (
	"GoPay/models"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func Notify(user *models.User, message string) error {
	notification := struct {
		User    string `json:"user"`
		Message string `json:"message"`
	}{
		User:    user.Email,
		Message: message,
	}

	jsonData, err := json.Marshal(notification)
	if err != nil {
		return err
	}

	resp, err := http.Post("https://util.devi.tools/api/v1/notify", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("Error sending notification")
	}

	return nil
}
