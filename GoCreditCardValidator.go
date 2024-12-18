package main

import (
  "encoding/json"
  "fmt"
  "log"
  "net/http"
  "strings"
  "unicode"
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
  cardNumber = strings.Map(func(r rune) rune {
    if unicode.IsDigit(r) {
      return r
    }
    return -1
  }, cardNumber)

  // Check if the card number is empty or contains non-digits
  if len(cardNumber) == 0 {
    return false
  }

  // Convert string to slice of digits
  digits := make([]int, len(cardNumber))
  for i, r := range cardNumber {
    digits[i] = int(r - '0')
  }

  // Starting from the rightmost digit and moving left
  for i := len(digits) - 2; i >= 0; i -= 2 {
    doubled := digits[i] * 2
    if doubled > 9 {
      doubled = doubled - 9
    }
    digits[i] = doubled
  }

  // Sum all digits
  sum := 0
  for _, digit := range digits {
    sum += digit
  }

  // Check if sum is divisible by 10
  return sum%10 == 0
}
