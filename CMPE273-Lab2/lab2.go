package main

import (
	"net/http"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type AbcRequest struct{
	Name string
}

type AbcResponse struct {
	Greeting string
}

func Webserv(rw http.ResponseWriter, req *http.Request){

	var jsonobject AbcRequest

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	if err:= req.Body.Close();
	err != nil{
		panic(err)
	}

	fmt.Println(string (body))


	if err:= json.Unmarshal(body, &jsonobject);
	err != nil {
		panic(err)
		fmt.Println("Unmarshal failed")
	}
	fmt.Println("Unmarshaled string is", jsonobject.Name)

        jsonobjectR := AbcResponse{"Hello,"+jsonobject.Name}
        js, err := json.Marshal(jsonobjectR);
        if err != nil {
        panic(err)
        }
        fmt.Println("Output string is", string (js))

        rw.Header().Set("Content-Type", "application/json")
        rw.Write(js)
}

func main() {
	http.HandleFunc("/hello", Webserv)
	err:= http.ListenAndServe(":8080",nil)
	if err != nil {
          panic(err)
        }
}