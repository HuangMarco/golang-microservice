package service

import (
	"log"
	"net/http"
)

func StartWebServer(port string) error {
	log.Println("Startig the web server at port:", port)

	//注意：使用如下gorilla配置将会失效，如果http.ListenAndServe(":" + port, RequestHandler(port))也存在的话
	//http.ListenAndServe(":8080", nil)生效，同时注释http.ListenAndServe(":" + port, RequestHandler(port))段配置
	// r := NewGorillaRouter()
	//仅匹配/accounts/v2/绝对路径
	// r.HandleFunc("/accounts/v2",AccountHandler)
	// http.Handle("/accounts/", r)
	// http.ListenAndServe(":8080", nil)

	// For http.ServeMux: https://pkg.go.dev/net/http#ServeMux
	err := http.ListenAndServe(":" + port, CommonHttpServeMuxHandler(port))
	if err != nil {
		log.Println("Error happened when the server is listening to port:" + port, err)
		log.Println("error:", err.Error())
		return err
	}
	return nil
}
