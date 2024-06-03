package tests

import (
	"GoPay/models"
	"GoPay/routes"
	"bytes"
	"encoding/json"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gorilla/mux"
)

func setupRouter() *mux.Router {
	r := routes.SetupRouter()
	return r
}

func randomName() string {
	names := []string{"João", "Maria", "Pedro", "Ana", "Carlos", "Mariana", "José", "Luiza"}
	return names[rand.Intn(len(names))]
}

func randomCPF() string {
	var cpf string
	for i := 0; i < 11; i++ {
		cpf += strconv.Itoa(rand.Intn(10))
	}
	return cpf
}

func randomEmail() string {
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	email := make([]byte, 10)
	for i := range email {
		email[i] = letters[rand.Intn(len(letters))]
	}
	return string(email) + "@example.com"
}

func TestRegisterUser(t *testing.T) {
	models.InitDatabase()
	r := setupRouter()

	user := map[string]interface{}{
		"name":        randomName() + " " + randomName(),
		"cpf":         randomCPF(),
		"email":       randomEmail(),
		"password":    "senha123",
		"is_merchant": false,
	}

	jsonUser, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonUser))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Handler retornou status errado: got %v want %v",
			status, http.StatusCreated)
	}

	var responseUser map[string]interface{}
	json.NewDecoder(rr.Body).Decode(&responseUser)

	if responseUser["name"] != user["name"] {
		t.Errorf("Expected name to be %v. Got %v", user["name"], responseUser["name"])
	}
	if responseUser["cpf"] != user["cpf"] {
		t.Errorf("Expected CPF to be %v. Got %v", user["cpf"], responseUser["cpf"])
	}
	if responseUser["email"] != user["email"] {
		t.Errorf("Expected email to be %v. Got %v", user["email"], responseUser["email"])
	}
	if responseUser["is_merchant"] != user["is_merchant"] {
		t.Errorf("Expected is_merchant to be %v. Got %v", user["is_merchant"], responseUser["is_merchant"])
	}
}

func TestTransfer(t *testing.T) {
	models.InitDatabase()
	r := setupRouter()

	payer := map[string]interface{}{
		"name":        randomName() + " " + randomName(),
		"cpf":         randomCPF(),
		"email":       randomEmail(),
		"password":    "senha123",
		"is_merchant": false,
	}
	jsonPayer, _ := json.Marshal(payer)
	reqPayer, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonPayer))
	reqPayer.Header.Set("Content-Type", "application/json")
	rrPayer := httptest.NewRecorder()
	r.ServeHTTP(rrPayer, reqPayer)

	payee := map[string]interface{}{
		"name":        "Maria Souza",
		"cpf":         "98765432100",
		"email":       "maria@example.com",
		"password":    "senha123",
		"is_merchant": false,
	}
	jsonPayee, _ := json.Marshal(payee)
	reqPayee, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonPayee))
	reqPayee.Header.Set("Content-Type", "application/json")
	rrPayee := httptest.NewRecorder()
	r.ServeHTTP(rrPayee, reqPayee)

	models.Users[0].Balance = 200.0

	transfer := map[string]interface{}{
		"value": 100.0,
		"payer": 1,
		"payee": 2,
	}
	jsonTransfer, _ := json.Marshal(transfer)
	reqTransfer, _ := http.NewRequest("POST", "/transfer", bytes.NewBuffer(jsonTransfer))
	reqTransfer.Header.Set("Content-Type", "application/json")
	rrTransfer := httptest.NewRecorder()
	r.ServeHTTP(rrTransfer, reqTransfer)

	if status := rrTransfer.Code; status != http.StatusBadRequest {
		t.Errorf("Esperava status de erro: %v, mas recebeu: %v", http.StatusBadRequest, status)
	}

	if models.Users[0].Balance != 200.0 {
		t.Errorf("Esperava que o saldo do pagador fosse 200.0, mas foi: %v", models.Users[0].Balance)
	}
	if models.Users[1].Balance != 0.0 {
		t.Errorf("Esperava que o saldo do recebedor fosse 0.0, mas foi: %v", models.Users[1].Balance)
	}
	t.Logf("Saldo Pagador: %v", models.Users[0].Balance)
	t.Logf("Saldo Receptor: %v", models.Users[1].Balance)

}
