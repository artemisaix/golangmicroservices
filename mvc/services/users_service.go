package services

import "github.com/artemisaix/golangmicroservices/mvc/domain"

//The services comunicates with the database to get the data
func GetUser(userId int64)(*domain.User, error){
return domain.GetUser(userId)

}