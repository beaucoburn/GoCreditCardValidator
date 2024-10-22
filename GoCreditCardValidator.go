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
  // Remove any spaces or dashes
  cardNumber = strings.ReplaceAll(cardNumber, " ", "")
  cardNumber = strings.ReplaceAll(cardNumber, "-", "")

  // Check if the card number contains only digits
  if _, err := strconv.Atoi(cardNumber); err != nil {
    return false
  }

  sum := 0
  isEven := true

  for i := len(cardNumber) - 1; i >= 0; i-- {
    digit, _ := strconv.Atoi(string(cardNumber[i]))

    if isEven {
      digit *= 2
      if digit > 9 {
        digit -= 9
      }
    }

    sum += digit
    isEven = !isEven
  }

  return sum%10 == 0
}
