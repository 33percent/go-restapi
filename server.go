package main

import (
  "fmt"
  "github.com/sandeepkumardev/go-restapi/config"
  "github.com/sandeepkumardev/go-restapi/app"
)



func main() {
	keys := keys.makeconnection()
	fmt.Println(keys);
	fmt.Println("ajhjash")
	app := &app.App{}
	app.Initialize(config)
	app.Run(":7080")

}