package service

import (
	"log"
	"net/http"
	"fmt"
)

//new handler
type apiHanlder struct {}
//等效于func (apiHanlder) ServeHTTP(w http.ResponseWriter, r *http.Request)
// 为结构体apiHanlder添加ServeHTTP方法
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

// How to use ServeMux: https://pkg.go.dev/net/http#ServeMux.Handle
func RequestHandler(port string) http.Handler {
	log.Println("The requested port is:", port)
	mux := http.NewServeMux()
	mux.Handle("/api/", apiHanlder{})
	mux.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
		// The "/" pattern matches everything, so we need to check
		// that we're at the root here.
		log.Printf("The request url is: %v", r.URL.Path)
		if r.URL.Path == "/test" {
			// log.Printf("Welcome to the %v page",r.URL.Path)
			fmt.Fprintf(w, "Welcome to the %v page!",r.URL.Path)
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