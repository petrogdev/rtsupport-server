package main

import (
  "net/http"
  r "github.com/dancannon/gorethink"
  "log"
)

func main() {
  session, err := r.Connect(r.ConnectOpts{
    Address: "localhost:28015",
    Database: "rtsupport",
  })

  if err != nil {
   log.Panic(err.Error())
  }

  router := NewRouter(session)

  router.Handle("channel add", addChannel)
  router.Handle("channel subscribe", subscribeChannel)
  router.Handle("channel subscribe", unsubscribeChannel)

  router.Handle("user edit", editUser)
  router.Handle("user subscribe", subscribeUser)
  router.Handle("user unsubscribe", unsubscribeUser)

  router.Handle("message add", addChannelMessage)
  router.Handle("message subscribe", subscribeChannelMessage)
  router.Handle("message unsubscribe", unsubscribeChannelMessage)

  http.Handle("/", router)
  http.ListenAndServe(":4000", nil)
}
