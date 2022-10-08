package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"

	"github.com/destafajri/golang-mongodb/models"
	
)

type UserController struct {
	session *mongo.Client
}

func NewUserController(s *mongo.Client) *UserController {
	return &UserController{s}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id:= p.ByName("id")
	ctx := context.Background()

	oid,_ := primitive.ObjectIDFromHex(id)

	var result models.User
	
	if err := uc.session.Database("mongodb-golang").Collection("users").FindOne(ctx, bson.M{"_id": oid}).Decode(&result); err != nil {
		w.WriteHeader(404)
		fmt.Println(err)
		return
	}

	uj, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(404)
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{}
	ctx := context.Background()

	json.NewDecoder(r.Body).Decode(&u)


	uc.session.Database("mongodb-golang").Collection("users").InsertOne(ctx, u)

	uj, err := json.Marshal(u)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", uj)
}

func (uc UserController) RemoveUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	ctx := context.Background()
	oid,_ := primitive.ObjectIDFromHex(id)

	var result models.User

	if err := uc.session.Database("mongodb-golang").Collection("users").FindOneAndDelete(ctx, bson.M{"_id": oid}).Decode(&result); err != nil {
		w.WriteHeader(404)
		fmt.Println(err)
		return
	}

	w.WriteHeader(200)
	fmt.Fprint(w, "Deleted user", oid, "\n")
}
