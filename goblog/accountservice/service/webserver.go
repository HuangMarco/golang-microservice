package service

import (
	"log"
	"net/http"
)

func StartWebServer(port string) error {

	log.Println("Startig the web server at port:", port)
	
	err := http.ListenAndServe(":" + port, nil)
	if err != nil {
		log.Println("Error happened when the server is listening to port:" + port, err)
		log.Println("error:", err.Error())
		return err
	}
	
	return nil
}