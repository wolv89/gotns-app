package handler

import (
	"net/http"

	"github.com/wolv89/gotnsapp/model"
	"github.com/wolv89/gotnsapp/util"
)


func ToolboxNuke(w http.ResponseWriter, req *http.Request) {

	model.DatabaseNuke()
	util.HttpSuccess(w, "DB Nuked")

}


func ToolboxLaunch(w http.ResponseWriter, req *http.Request) {

	model.DatabaseLaunch()
	util.HttpSuccess(w, "DB Launched")

}