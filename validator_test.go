package main

import (
  "bytes"
  "encoding/json"
  "net/http"
  "net/http/httptest"
  "testing"
)

func TestLuhnAlgorithm(t *testing.T) {
  tests := []struct {
    name string
    cardNumber string
    want bool
  }{
    {"Valid Visa", "4532015112830366", true},
    {"Valid Mastercard", "5425233430109903", true},
    {"Valid Amex", "374245455400126", true},
    {"Invalid Number", "4532015112830367", false},
    {"Empty String", "", false},
    {"Non-numeric", "123abc456", false},
    {"With Spaces", "4532 0151 1283 0366", true},
    {"With Dashes", "4532-0151-1283-0366", true},
  }

  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      got := luhnAlgorithm(tt.cardNumber)
      if got != tt.want {
        t.Errorf("luhnAlgorithm(%s) = %v, want %v", tt.cardNumber, got, tt.want)
      }
    })
  }
}

func TestValidateHandler(t *testing.T) {
  tests := []struct {
    name string
    cardNumber string
    expectedStatus int
    expectedValid bool
  }{
    {"Valid Card", "4532015112830366", http.StatusOK, true},
    {"Invalid Card", "4532015112830367", http.StatusOK, false},
    {"Invalid Format", "123abc456", http.StatusOK, false},
  }


}
