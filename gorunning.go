package main

import (
  "flag"
  "fmt"
  "log"
  "os"
  "os/signal"
  "syscall"
  "reflect"

  "github.com/coreos/pkg/flagutil"
  "github.com/dghubble/go-twitter/twitter"
  "github.com/dghubble/oauth1"
)

//func hello(w http.ResponseWriter, r *http.Request) {
//  io.WriteString(w, "Hello world!")
//}

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
  
  fmt.Println("All requirements met...")
  
  config := oauth1.NewConfig(*consumerKey, *consumerSecret)
  token := oauth1.NewToken(*accessToken, *accessSecret)
  // OAuth1 http.Client will automatically authorize Requests
  httpClient := config.Client(oauth1.NoContext, token)

  // Twitter Client
  client := twitter.NewClient(httpClient)

  // Convenience Demux demultiplexed stream messages
  demux := twitter.NewSwitchDemux()
  demux.Tweet = func(tweet *twitter.Tweet) {
    //This is the tweet
    //fmt.Println(tweet.Text)
    fmt.Println("tweet is a type of: ", reflect.TypeOf(tweet))
    fmt.Println("place: ", tweet.Place)
    //fmt.Println(tweet.Place.Attributes)
    //if tweet.Place != nil {
    //  fmt.Println("has a place: ", tweet.Place)
    //  if tweet.Place.BoundingBox != nil {
    //    fmt.Println("has a BOUNDING BOX")
    //      fmt.Println(fmt.Printf("%+v\n", tweet.Place.BoundingBox))
    //  }
    //} else {
    //  fmt.Println("does not have a place")
    //}
    if tweet.Coordinates != nil {
      fmt.Println(fmt.Printf("%+v\n", tweet.Coordinates))
    } else {
    }
    fmt.Println(fmt.Printf("%+v\n", tweet.Place))
  }
  demux.DM = func(dm *twitter.DirectMessage) {
    fmt.Println(dm.SenderID)
  }
  demux.Event = func(event *twitter.Event) {
    fmt.Printf("%#v\n", event)
  }

  fmt.Println("Starting Stream...")

  // FILTER
  filterParams := &twitter.StreamFilterParams{
    Track:         []string{"and"},
    StallWarnings: twitter.Bool(true),
  }
  stream, err := client.Streams.Filter(filterParams)
  if err != nil {
    log.Fatal(err)
  }

  // USER (quick test: auth'd user likes a tweet -> event)
  // userParams := &twitter.StreamUserParams{
  //  StallWarnings: twitter.Bool(true),
  //  With:          "followings",
  //  Language:      []string{"en"},
  // }
  // stream, err := client.Streams.User(userParams)
  // if err != nil {
  //  log.Fatal(err)
  // }

  // SAMPLE
  // sampleParams := &twitter.StreamSampleParams{
  //  StallWarnings: twitter.Bool(true),
  // }
  // stream, err := client.Streams.Sample(sampleParams)
  // if err != nil {
  //  log.Fatal(err)
  // }

  // Receive messages until stopped or stream quits
  go demux.HandleChan(stream.Messages)

  // Wait for SIGINT and SIGTERM (HIT CTRL-C)
  ch := make(chan os.Signal)
  signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
  log.Println(<-ch)

  fmt.Println("Stopping Stream...")
  stream.Stop()
   
}
