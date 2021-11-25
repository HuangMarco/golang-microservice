package service

import (
	// "fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//Definbe the route includes: the HTTP method, the pattern, the function that will be executed when the route is
//is matched
type Route struct {
	Name string `json:"name"`
	Method string `json:"method"`
	Pattern string `json:"pattern"`
	HandlerFunc http.HandlerFunc
}

//包含一组路由对象
type Routes []Route

//initialize our own route object
var routes = Routes{
	//声明一个route对象
	Route{
		Name: "GetAccount",
		Method: "GET",
		Pattern: "/account/{accountId}",
		HandlerFunc: func(w http.ResponseWriter, r *http.Request){
			w.Header().Set("Content-Type", "application/json;charset=UTF-8")
			// w.Write([]byte("{\"result\": {\"OK\"}}"))
			w.Write([]byte("{\"result\": {\"OK\"}}"))
		},
	},
	Route{
		Name: "SetAccount",
		Method: "Post",
		Pattern: "/account/{accountId}",
		HandlerFunc: func(w http.ResponseWriter, r *http.Request){
			w.Header().Set("Content-Type", "application/json;charset=UTF-8")
			// w.Write([]byte("{\"result\": {\"OK\"}}"))
			w.Write([]byte("{\"result\": {\"OK\"}}"))
		},
	},
}

//使用Gorrila Router: https://github.com/gorilla/mux#overview
//返回一个指针对象，该对象是一个Girilla Router
func NewRouter() *mux.Router {

	//创建Gorilla Router的一个实例
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		//Attach each route, uses a Builder-like pattern to set each route up
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(route.HandlerFunc)
	}

	return router
}
