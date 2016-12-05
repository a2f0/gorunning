package main

import (
  "io"
  "flag"
  "net/http"
  "os"
  "log"

  "github.com/coreos/pkg/flagutil"
  //"github.com/dghubble/go-twitter/twitter"
  //"github.com/dghubble/oauth1"

)

func hello(w http.ResponseWriter, r *http.Request) {
  io.WriteString(w, "Hello world!")
}

func main() {
  //http.HandleFunc("/", hello)
  //http.ListenAndServe(":8000", nil)

  flags := flag.NewFlagSet("user-auth", flag.ExitOnError)
  consumerKey := flags.String("consumer-key", "", "Twitter Consumer Key")
  consumerSecret := flags.String("consumer-secret", "", "Twitter Consumer Secret")
  accessToken := flags.String("access-token", "", "Twitter Access Token")
  accessSecret := flags.String("access-secret", "", "Twitter Access Secret")
  flags.Parse(os.Args[1:])
  flagutil.SetFlagsFromEnv(flags, "TWITTER")

  if *consumerKey == "" || *consumerSecret == "" || *accessToken == "" || *accessSecret == "" {
    log.Fatal("Consumer key/secret and Access token/secret required")
  } 

}
