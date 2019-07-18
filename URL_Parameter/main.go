package main

import (
	//"fmt"
	"net/http"
	"io"
)
func main() {
	
	handler := func(w http.ResponseWriter, req *http.Request) {
	
	/* handler := func(w http.ResponseWriter, req *http.Request) {
	
		io.WriteString(w, "HelloWorld!")
	}
	name := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w,"Hello niha!")
	}*/
	/*if req.Method == "GET" {
			io.WriteString(w, "HelloWorld!")	
		} else if req.Method == "POST" {
			io.WriteString(w,"Hello niha!")
		}
		
	}*/
	//params := req.URL.Query() 
	//for k,v:= range params {
		//fmt.Printf("%s:%s\n", k,v)
		//w.Write([]byte("name: "+req.FormValue("name")))
		//v := params["name"]
		io.WriteString(w,"fname: "+req.FormValue("fname"))
		io.WriteString(w,"\nlname: "+req.FormValue("lname"))
		//fmt.Println(params["name"])
	//}

	//}
		
	}


	http.HandleFunc("/hello",handler)
	
	//http.HandleFunc("/niha", name)
	http.ListenAndServe(":8080",nil)
}
