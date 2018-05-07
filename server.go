package main

import (
  "net/http"
  "strings"
  _ "github.com/go-sql-driver/mysql"
  "fmt"
)



func main() {
  if err := http.ListenAndServe(":3000", nil); err != nil {
    panic(err)
  }
  http.HandleFunc("/create", createtodo)
  http.HandleFunc("/update",updatetodo)
  http.HandleFunc("/delete",deletetodo)
  http.HandleFunc("/read",readtodo)

}