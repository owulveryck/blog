+++
date = "2015-10-22T20:40:36+02:00"
draft = true
title = "Ruby and go dialog via ZMQ"

+++

# Abtract

I really like go as a programming language. It is a good tool to develop web restful API service.

On the other hand, ruby and its framework rails has also been wildly used to achieve the same goal.

Therefore we may be facing a "legacy" ruby developpement that we would like to connect to our brand new go framework.
0MQ may be a perfect choice for intefacing the two languages.

Anyway, it is, at least, a good experience to do a little bit of code to make them communicate.

# ZeroMQ

I will use the ZeroMQ version 4 as it is the latest available one.
On top of that, I can see in the [release notes](http://zeromq.org/docs:changes-4-0-0) that there is an implementation of a strong encryption, and I may use it later on 
# Go

## Installation of the library

As written in the README file, I try a `go get` installation on my chromebook.
```
~ go get github.com/pebbe/zmq4
# pkg-config --cflags libzmq
Package libzmq was not found in the pkg-config search path.
Perhaps you should add the directory containing `libzmq.pc'
to the PKG_CONFIG_PATH environment variable
No package 'libzmq' found
pkg-config: exit status 1
```

The go binding is not a pure go implementation, and it still needs the C library of zmq.

Let's _brew installing_ it:

```
~  brew install zmq
==> Downloading http://download.zeromq.org/zeromq-4.1.3.tar.gz
######################################################################## 100.0%
==> ./configure --prefix=/usr/local/linuxbrew/Cellar/zeromq/4.1.3 --without-libsodium
==> make
==> make install
/usr/local/linuxbrew/Cellar/zeromq/4.1.3: 63 files, 3.5M, built in 73 seconds
```

Let's do the go-get again:

```
~ go get github.com/pebbe/zmq4
```

so far so good. Now let's test the installation with a "hello world" example.

_Note_: the [examples directory](https://github.com/pebbe/zmq4/blob/master/examples) contains a go implementation of all the example of the ZMQ book
I will use the [hello world client](https://github.com/pebbe/zmq4/blob/master/examples/hwclient.go) and the [hello world server](https://github.com/pebbe/zmq4/blob/master/examples/hwserver.go) for my tests

The hello world client/server is implementing a Request-Reply patternt and are communicating via a TCP socket.
* The server is the *replier* and is listening on the TCP port 5555
```go
...
func main() {
    //  Socket to talk to clients
    responder, _ := zmq.NewSocket(zmq.REP)
    defer responder.Close()
    responder.Bind("tcp://*:5555")
    ...
}
```
* The client is the *requester* and is dialing the same TCP port
```go
...
func main() {
    //  Socket to talk to server
    fmt.Println("Connecting to hello world server...")
    requester, _ := zmq.NewSocket(zmq.REQ)
    defer requester.Close()
    requester.Connect("tcp://localhost:5555")
    ...
}
```

Then, the client is sending (requesting) a _hello_ message, and the server is replying a _world_ message.

## Running the example
First, start the server:

```
~ cd $GOPATH/src/github.com/pebbe/zmq4/examples
~ go run hwserver.go
```

Then the client

```
~ cd $GOPATH/src/github.com/pebbe/zmq4/examples
~ go run hwclient.go
Connecting to hello world server...
Sending  Hello 0
Received  World
Sending  Hello 1
Received  World
Sending  Hello 2
...
```
