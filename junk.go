package main

import (
  "encoding/json"
  "fmt"
)

type Message struct {
  Name string
  Data interface{}
}

func main() {
  recRawMsg := []byte(`{name: "channel add", ` +
    `"data: {"name: "Hardware Support"}}`)

  var recMessage Message
  err := json.Unmarshal(recRawMsg, &recMessage)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Fprintf("%#v\n", recMessage)
}


// type Speak interface {
//   Speak()
// }
//
// func (m Message) Speak() {
//   fmt.Println("I'm a " + m.Name + " event!")
// }
//
// func SomeFunc(speaker Speaker) {
//   speaker.Speak()
// }
