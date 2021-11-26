package service

import (
	"fmt"
	"net/http"
	"log"
	"os"

	"github.com/gorilla/mux"
)

func AccountHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
  w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
}

//使用Gorrila Router: https://github.com/gorilla/mux#overview
//返回一个指针对象，该对象是一个Girilla Router
func NewGorillaRouter() *mux.Router {

	//创建Gorilla Router的一个实例
	router := mux.NewRouter()
	router.HandleFunc("/api/accounts/", AccountHandler)

	//遍历我们在routes.go中声明的所有routes，都添加到router中
	for _, route := range routes {
		//Attach each route, uses a Builder-like pattern to set each route up
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(route.HandlerFunc)
	}
	return router
}

//不同api对应的hanlder，封装对应的业务逻辑
type ApiV1Hanlder struct {}
//等效于func (apiHanlder) ServeHTTP(w http.ResponseWriter, r *http.Request)
//为结构体apiHanlder添加ServeHTTP方法,通常是业务对应的方法
//在golang中，标准的为某个struct添加对于某个方法的实现
func (handler ApiV1Hanlder) ServeHTTP(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the %v page provided by the CommonApiHanlder!\n", r.URL.Path)
}

//第二种业务逻辑封装的方式
func BusinessForTestAPI(path string){
	log.Printf("The request is for path: %v", path)
	fmt.Fprintf(os.Stdout, "Do something for /test api request.\n")
}

func HandleApiV2(w http.ResponseWriter, r *http.Request){
	// The "/" pattern matches everything, so we need to check
	// that we're at the root here.
	log.Printf("The request url is: %v", r.URL.Path)
	if r.URL.Path == "/api/v2/test/" {
		// log.Printf("Welcome to the %v page",r.URL.Path)
		//fmt.Fprintf将会将内容输出到输出流中:https://pkg.go.dev/fmt#Fprintf
		fmt.Fprintf(w, "Welcome to the %v page!",r.URL.Path)
		BusinessForTestAPI(r.URL.Path)
		//此处退出，而不会将server暂停，只是针对于/test请求的处理结束
		return

	}else if r.URL.Path == "/api/"{
		fmt.Fprintf(w, "Welcome to the %v page!",r.URL.Path)
		return
	}else{
		http.NotFound(w, r)
		log.Printf("The url %v you try to access is not valid", r.URL.Path)
		// errors.New("The url  is not valid")
		return
	}
}

// How to use ServeMux: https://pkg.go.dev/net/http#ServeMux.Handle
func CommonHttpServeMuxHandler(port string) http.Handler {

	log.Println("The requested port is:", port)
	
	//创建一个新的ServeMux来处理restful api
	//也可以使用Gorrila Router: https://github.com/gorilla/mux#overview
	mux := http.NewServeMux()
	//单独为/api/添加一个handler，而不会再交给HandleFunc
	mux.Handle("/api/v1/", ApiV1Hanlder{})
	
	//不管是/api, /account/{accountId}, /apitest等，都是以/开头
	mux.HandleFunc("/api/v2/", HandleApiV2)

	return mux
}