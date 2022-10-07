package main

import(
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"net/http"

	"github.com/destafajri/golang-mongodb/controllers"
)

func main() {
	
	r := httprouter.New()
	usercontrol := controllers.NewUserController(getSession())
	r.GET("/user/:id", usercontrol.GetUser)
	r.POST("/user", usercontrol.CreateUser)
	r.DELETE("/user/:id", usercontrol.DeleteUser)

	http.ListenAndServe("localhost:8000", r)
}

func getSession() *mgo.Session{

	s, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil{
		panic(err)
	}
	return s

}