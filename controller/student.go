package controller

import (
	echo "github.com/labstack/echo/v4"
	"net/http"
	"test/model"
	"test/storage"
)

// Handler
func GetStudents(c echo.Context) error {
	students,_:=GetRepoStudents()
	return c.JSON(http.StatusOK, students)
}
func GetRepoStudents()([]model.Students,error){
	db:=storage.GetDbInstance()
	students :=[]model.Students{}
	if err:=db.Find(&students).Error;err!=nil{
		return nil,err
	}
	return students,nil

}