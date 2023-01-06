package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/artemisaix/golangmicroservices/mvc/services"
)

func GetUser(resp http.ResponseWriter, req *http.Request){
userId,err:=strconv.ParseInt(req.URL.Query().Get("user_id"),10,64)
if err!=nil{
	//Return the Bad Request to the Client
		resp.WriteHeader((http.StatusNotFound))
	resp.Write([]byte("user_id must be a number"))
	return
}
user, err:= services.GetUser(userId)
if err != nil {
	resp.WriteHeader((http.StatusNotFound))
	resp.Write([]byte(err.Error()))
	//Handle the err and return to the client
	return
}
//return user to client
jsonValue,_:=json.Marshal(user)
resp.Write(jsonValue)
}