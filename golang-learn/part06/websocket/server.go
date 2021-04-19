package websocket

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func mian()  {
	router := mux.NewRouter()

	go h.run()

	router.HandleFunc("/ws", myws)

	if err := http.ListenAndServe("127.0.0.1:8000", router); err != nil{
		fmt.Println("err:", err)
	}

}