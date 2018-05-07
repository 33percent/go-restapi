package main

import (
  "fmt"
  "github.com/sandeepkumardev/go-restapi/keys"
)



func main() {
	keys := keys.makeconnection()
	fmt.Println(keys)
	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")

}