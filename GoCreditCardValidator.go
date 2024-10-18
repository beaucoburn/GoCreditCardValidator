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
  fmt.Println("Server is running on http://locahost:8080")
}
