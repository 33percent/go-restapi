package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/sandeepkumardev/go-restapi/app/model"
)

func createtask(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

}