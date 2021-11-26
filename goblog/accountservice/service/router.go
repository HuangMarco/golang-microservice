package service

import (
	"github.com/gorilla/mux"
)

//使用Gorrila Router: https://github.com/gorilla/mux#overview
//返回一个指针对象，该对象是一个Girilla Router
func NewRouter() *mux.Router {

	//创建Gorilla Router的一个实例
	router := mux.NewRouter()
	router.HandleFunc("/api/accounts/",AccountHandler)

	//遍历我们在routes.go中声明的所有routes，都添加到router中
	for _, route := range routes {
		//Attach each route, uses a Builder-like pattern to set each route up
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(route.HandlerFunc)
	}

	return router
}