package service

import (
	"log"
	"net/http"
	"fmt"
	"os"
)

//new handler
type apiHanlder struct {}
//等效于func (apiHanlder) ServeHTTP(w http.ResponseWriter, r *http.Request)
//为结构体apiHanlder添加ServeHTTP方法
func (handler apiHanlder) ServeHTTP(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the %v page!",r.URL.Path)

}

func StartWebServer(port string) error {

	log.Println("Startig the web server at port:", port)
	// For http.ServeMux: https://pkg.go.dev/net/http#ServeMux
	err := http.ListenAndServe(":" + port, RequestHandler(port))
	if err != nil {
		log.Println("Error happened when the server is listening to port:" + port, err)
		log.Println("error:", err.Error())
		return err
	}
	
	return nil
}

func forTestAPI(path string){
	log.Printf("The request is for path: %v", path)
	fmt.Fprintf(os.Stdout, "Do something for /test api request.\n")
}

// How to use ServeMux: https://pkg.go.dev/net/http#ServeMux.Handle
func RequestHandler(port string) http.Handler {
	log.Println("The requested port is:", port)
	//创建一个新的ServeMux来处理restful api 
	mux := http.NewServeMux()
	//单独为/api/添加一个handler，而不会再交给HandleFunc
	mux.Handle("/api/", apiHanlder{})
	mux.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
		// The "/" pattern matches everything, so we need to check
		// that we're at the root here.
		log.Printf("The request url is: %v", r.URL.Path)
		if r.URL.Path == "/test" {
			// log.Printf("Welcome to the %v page",r.URL.Path)
			//fmt.Fprintf将会将内容输出到输出流中:https://pkg.go.dev/fmt#Fprintf
			fmt.Fprintf(w, "Welcome to the %v page!",r.URL.Path)
			forTestAPI(r.URL.Path)
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
		
	})

	return mux
}