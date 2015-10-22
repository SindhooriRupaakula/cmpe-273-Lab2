package main

import ("fmt"
        "net/http"
        "encoding/json"
        "github.com/julienschmidt/httprouter")

type JSONObj struct {
  Name string `json:"name"`
}

type JSONResp struct {
  Resp string `json:"response"`
}

func getName(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
  fmt.Fprintf(rw, "Hi %s, This is the response for a GET request.", p.ByName("name"))
}

func postName(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
  var jsonObj JSONObj
  var jsonResp JSONResp

  dec := json.NewDecoder(r.Body)
  err := dec.Decode(&jsonObj)
  if err != nil {
    panic(err)
  }

  fmt.Println(r.Body)

  jsonResp.Resp = "Hi " +jsonObj.Name + ", This is the response for a POST request."

  fmt.Println(jsonResp.Resp)
  output,_ := json.Marshal(jsonResp)
  fmt.Fprintf(rw,string(output))
}

func main() {
  r := httprouter.New()
  r.GET("/hello/:Name",getName)
  r.POST("/hello/", postName)
  http.ListenAndServe("localhost:8080", r)
}

