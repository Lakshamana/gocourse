package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Address struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

// type WriteCounter struct {
// 	TotalRead     uint64
// 	ContentLength int64
// 	progress      int
// }
//
// func (wc *WriteCounter) PrintProgress() {
//   fmt.Printf("\r%s", strings.Repeat(" ", 35))
// 	steps := [4]string{"|", "/", "-", "\\"}
// 	fmt.Printf("\rDownloaded %d%%... %s", wc.TotalRead/uint64(wc.ContentLength), steps[wc.progress%4])
// 	wc.progress++
// }
//
// func (wc *WriteCounter) Write(p []byte) (int, error) {
//   println(">> write")
// 	n := len(p)
// 	wc.TotalRead += uint64(n)
// 	wc.PrintProgress()
// 	return n, nil
// }

func main() {
	http.HandleFunc("/", searchZipCodeHandler)
	http.ListenAndServe(":8080", nil)
}

func searchZipCodeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	zipCode := r.URL.Query().Get("zip_code")
	if zipCode == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	address, err := searchZipCode(zipCode)
	if err != nil {
		w.WriteHeader(http.StatusFailedDependency)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(&address)
}

func searchZipCode(zipCode string) (*Address, error) {
	res, err := http.Get(fmt.Sprintf("https://viacep.com.br/ws/%s/json", zipCode))
	// counter := &WriteCounter{ContentLength: res.ContentLength}
	// responseBody := io.TeeReader(res.Body, counter)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var address Address
	err = json.Unmarshal(body, &address)
	if err != nil {
		return nil, err
	}

	return &address, nil
}
