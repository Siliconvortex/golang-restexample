package main

import (
  "fmt"
  "encoding/json"
  "net/http"
  "github.com/siliconvortex/golang-restexample/produce"
  "github.com/siliconvortex/golang-restexample/common"
  "strconv"
  "flag"
)

func serveRest(w http.ResponseWriter, r *http.Request) {
  response, err := getJsonResponse()
  if err != nil {
    panic(err)
  }

  fmt.Fprintf(w, string(response))
}

func main() {
  
  portPtr := flag.Int("port", common.DEFAULT_PORT, "the port to listen on")
  flag.Parse()

  fmt.Println("Listening on port", *portPtr)

  http.HandleFunc("/", serveRest)
  http.ListenAndServe("localhost:" + strconv.Itoa(*portPtr), nil)
}

func getJsonResponse() ([]byte, error) {
  // make the fruit basket
  fruits := make(produce.Fruits)
  fruits.Set("Apples", 25)
  fruits.Set("Oranges", 11)

  // make the veggie basket
  vegetables := make(produce.Vegetables)
  vegetables.Set("Carrots", 10)
  vegetables.Set("Peppers", 0)

  // put the baskets together (per the original youtube post)
  d := produce.Data{fruits, vegetables}
  p := produce.Payload{d}

  //serialize to json and return
  return json.Marshal(p)
}
