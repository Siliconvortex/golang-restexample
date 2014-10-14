package main

import (
  "github.com/siliconvortex/golang-restexample/produce"
  "github.com/siliconvortex/golang-restexample/common"
  "encoding/json"
  "fmt"
  "net/http"
  "io/ioutil"
  "flag"
  "strconv"
)

func main() {
  portPtr := flag.Int("port", common.DEFAULT_PORT, "the port to listen on")
  flag.Parse()

  url := "http://localhost:" + strconv.Itoa(*portPtr)

  fmt.Println("Getting data from '" + url + "'")

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
