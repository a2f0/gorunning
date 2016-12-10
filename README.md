# Overview

Go experimentation.

# Install

Install Go

`brew install go`

Setup a Go environment
```
mkdir -p ~/working/github/go/src/
echo 'export GOPATH=~/working/github/go' >> ~/.bashrc
echo 'export PATH=$PATH:~/working/github/go/bin' >> ~/.bashrc
export PATH=$PATH:~/working/github/go/bin
```

Clone it
```
cd ~/working/github/go/src/
git clone git@github.com:dsulli99/gorunning.git
```

Start it
```
export TWITTER_CONSUMER_KEY=xxx
export TWITTER_CONSUMER_SECRET=xxx
export TWITTER_ACCESS_TOKEN=xxx
export TWITTER_ACCESS_SECRET=xxx
go run gorunning.go
```

Getting help
`go run gorunning.go --help`

# Godep notes 

Install godep
```
go get github.com/tools/godep
```
