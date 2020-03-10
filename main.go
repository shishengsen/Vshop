package main

import "net/http"

//
func sayHello(w http.ResponseWriter, r *http.Request){
	w.Write([]byte ("Hello world"))
}

func main(){

	fileHandler := http.FileServer(http.Dir("./video"))
	http.Handle("/video",http.StripPrefix("/video",fileHandler))
	//注册handle
	http.HandleFunc("/sayHello", sayHello)

	//启动web服务
	http.ListenAndServe(":8090",nil)

}
