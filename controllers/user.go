package controllers

import(
	"fmt"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"

)

type UserController struct{
	sessiopn *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController{
	return &UserController{s}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params){

}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){

}