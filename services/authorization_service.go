package services

import (
	"fmt"
	"log"
	"net/http"
)

func Authorize() (bool, error) {
	url := "https://util.devi.tools/api/v2/authorize"

	log.Printf("Realizando solicitação de autorização para: %s", url)

	resp, err := http.Get(url)
	if err != nil {
		return false, fmt.Errorf("Erro ao realizar solicitação de autorização: %v", err)
	}
	defer resp.Body.Close()

	log.Printf("Status da solicitação de autorização: %s", resp.Status)

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("Solicitação de autorização falhou. Status: %v", resp.Status)
	}

	log.Println("Autorização concedida com sucesso")

	return true, nil
}
