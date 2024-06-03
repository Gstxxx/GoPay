package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"GoPay/models"
	"GoPay/utils"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:sua_senha_secreta@tcp(localhost:3306)/gopay_mysql")
	if err != nil {
		panic(err.Error())
	}
}

type RegisterRequest struct {
	Name       string `json:"name"`
	CPF        string `json:"cpf"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	IsMerchant bool   `json:"is_merchant"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = validateRegisterRequest(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	stmt, err := db.Prepare("INSERT INTO users (Name, CPF, Email, Password, IsMerchant, Balance) VALUES (?, ?, ?, ?, ?, 0.0)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	result, err := stmt.Exec(req.Name, req.CPF, req.Email, req.Password, req.IsMerchant)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user := models.User{
		ID:         int(id),
		Name:       req.Name,
		CPF:        req.CPF,
		Email:      req.Email,
		Password:   req.Password,
		IsMerchant: req.IsMerchant,
		Balance:    0.0,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func validateRegisterRequest(req RegisterRequest) error {
	for _, user := range models.Users {
		if user.CPF == req.CPF {
			return fmt.Errorf("CPF already registered")
		}
	}

	stmt, err := db.Prepare("SELECT COUNT(*) FROM users WHERE email = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	var count int
	err = stmt.QueryRow(req.Email).Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		return fmt.Errorf("Email already registered")
	}

	if req.IsMerchant {
		if !utils.IsValidCNPJ(req.CPF) {
			return fmt.Errorf("Invalid CNPJ")
		}
	}

	return nil
}
