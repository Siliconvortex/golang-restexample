package main

import (
  "github.com/siliconvortex/golang-restexample/produce"
  "encoding/json"
  "fmt"
  "net/http"
  "io/ioutil"
)

func main() {
  url := "http://localhost:12337"
  res, err := http.Get(url)
  if err != nil {
    panic(err)
  }
  defer res.Body.Close()

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    panic(err)
  }

  var p produce.Payload

  err = json.Unmarshal(body, &p)
  if err != nil {
    panic(err)
  }

  fmt.Println(p.Stuff.Fruit)
  fmt.Println(p.Stuff.Veggies)
}
