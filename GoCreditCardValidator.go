package main

import (
  "encoding/json"
  "fmt"
  "log"
  "net/http"
  "strconv"
  "strings"
)

func main() {
  http.HandleFunc("/validate", validateHandler)
  fmt.Println("Server is running on http://locahost:8081")
  log.Fatal(http.ListenAndServe(":8081", nil))
}

func validateHandler(w http.ResponseWriter, r *http.Request){
  if r.Method != http.MethodPost {
    http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    return
  }

  var request struct {
    CardNumber string `json:"card_number"`
  }

  err := json.NewDecoder(r.Body).Decode(&request)
  if err != nil {
    http.Error(w, "Invalid request body", http.StatusBadRequest)
    return
  }

  isValid := luhnAlgorithm(request.CardNumber)

  response := struct {
    IsValid bool `json:"is_valid"`
  }{
    IsValid: isValid,
  }

  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(response)
}

func luhnAlgorithm(cardNumber string) bool {

}
