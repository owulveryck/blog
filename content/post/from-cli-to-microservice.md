---
categories:
date: 2017-09-02T13:28:36+02:00
description: ""
draft: true
images:
- /assets/images/terraformcli.png
tags:
title: From a command line tools to a microservice - The example of hashicorp tools (terraform) and grpc
---

This post is a little bit different from the last ones. As usual the introduction tries to be open, but it quickly goes deeper into a go implementation.
Some of the explanations may be tricky from time to times and therefore not very clear. As usual, do not hesitate to send me any comment via this blog or via twitter [@owulveryck](https://twitter.com/owulveryck).

# About the cli utilities

I come from the sysadmin world... Precisely the Unix world (I have been a BSD user for years). Therefore I have learned to use and love "_the cli utilities_". Cli utilities are all those tools that makes Unix sexy and "user friendly". 

<center>
Because, yes, Unix **is user-friendly** (it's just picky about its friends[^1]).
</center>

[^1]: This sentence is not from me. I read it once somewhere on the Internet but I cannot find anybody to give the credit to.

From a user perspective, cli tools remains a must nowadays because:

* there are usually developed in the pure Unix philosophy: simple enough to use for what they were made for;
* they can be easily wrapped into scripts. Therefore, it is easy to automate cli actions.

The point with cli application is that they are mainly developed for an end-user that we call "an operator". As Unix is a multi-user operating system, several operators can use the same tool, but they have to be logged onto the same host.

In case of a remote execution, it's possible to execute the cli via `ssh`, but dealing with automation, network interruption and resuming starts to be tricky.
For remote and concurrent execution web-services are more suitable.

Let's see if turning a cli tool into a webservice without recoding the all logic is easy in go?

## Hashicorp's cli

For the purpose of this post, and because I am using Hashicorp tools at work, I will take [@mitchellh](https://twitter.com/mitchellh)'s framework for developping command line utilities.
This package is used in all of the Hashicorp tools and is called................ "[cli](https://github.com/mitchellh/cli)"! 

This library provides a [`Command`](https://godoc.org/github.com/mitchellh/cli#Command) type that represents any action that the cli will execute.
`Command` is a go `interface` composed of three methods:

* `Help()` that returns a string describing how to use the command;
* `Run(args []string)` that takes an array of string as arguments (all cli parameters of the command) and returns and integer (exit code);
* `Synopsis()` that returns a string describing what the command is about.
 
 _Note_: I assume that you know what an interface is (specially in go). If you don't, just google, or even better, buy the book [The Go Programming Language](https://www.amazon.com/Programming-Language-Addison-Wesley-Professional-Computing/dp/0134190440) and read the _chapter 7_ :).

The main object that holds the business logic of the cli package is an implementation of [`Cli`](https://godoc.org/github.com/mitchellh/cli#CLI). 
One of the element of the Cli structure is `Commands` which is a `map` that takes the name of the action as key. The name passed is a string and is the one that will be used on the command line. The value of the `map` is a function that returns a `Command`. This function is named [`CommandFactory`](https://godoc.org/github.com/mitchellh/cli#CommandFactory). According to the documentation, the factory is needed because _we may need to setup some state on the struct that implements the command itself_. Good idea!

## Example

First, let's create a very simple tool using now.
The tool will have two "commands":

* hello: will display _hello args...._  on `stdout` 
* goodbye: will display _goodbye args..._ on `stderr`

{{< highlight go >}}
func main() {
      c := cli.NewCLI("server", "1.0.0")
      c.Args = os.Args[1:]
      c.Commands = map[string]cli.CommandFactory{
            "hello": func() (cli.Command, error) {
                      return &HelloCommand{}, nil
            },
            "goodbye": func() (cli.Command, error) {
                      return &GoodbyeCommand{}, nil
            },
      }
      exitStatus, err := c.Run()
      ... 
}
{{</ highlight >}}
As seen before, the first object created is a `Cli`. Then the `Commands` field is filled with the two commands "hello" and "goodbye" as keys, and an anonymous function that simply returns two structures that will implement the `Command` interface.

Now, let's create the `HelloCommand` structure that will fulfill the [`cli.Command`](https://godoc.org/github.com/mitchellh/cli#Command) interface:

{{< highlight go >}}
type HelloCommand struct{}

func (t *HelloCommand) Help() string {
      return "hello [arg0] [arg1] ... says hello to everyone"
}

func (t *HelloCommand) Run(args []string) int {
      fmt.Println("hello", args)
      return 0
}

func (t *HelloCommand) Synopsis() string {
      return "A sample command that says hello on stdout"
}
{{</ highlight >}}

The `GoodbyeCommand` is similar, and I omit it for brevity.

After a simple `go build`, here is the behaviour of our new cli tool:
{{< highlight shell >}}
~ ./server help
Usage: server [--version] [--help] <command> [<args>]

Available commands are:
    goodbye    synopsis...
    hello      A sample command that says hello on stdout

~ ./server hello -help
hello [arg0] [arg1] ... says hello to everyone

~ ./server/server hello a b c
hello [a b c]
{{</ highlight >}}

So far, so good!
Now, let's see if we can turn this into a webservice.

# Micro-services

<center>_The biggest issue in changing a monolith into microservices lies in changing the communication pattern._[^2]</center>

[^2]: from [Martin Fowler's Microservices definition](https://martinfowler.com/articles/microservices.html#SmartEndpointsAndDumbPipes).

There is, according to me, two options to consider to turn our application into a webservice:

* a RESTish communication and interface;
* a RPC based communication.

SOAP is not an option anymore because it does not provide any advantage over the REST and RPC methods.

## Rest? 

I've always been a big fan of the REST "protocol". It is easy to understand and to write. On top of that it is verbose and allows a good description of "business objects".
But, its verbosity that is a strength quickly become a weakness when applied for machine to machine communication.
The "contract" between the client and the server have to be documented manually (via something like swagger for example). And, as you only transfer objects and states, the server must handle the request, understand it, and apply it to any business logic before returning a result.
Don't get me wrong, REST remains a very good thing. But it is very good when you think about it from the beginning of your conception (and with a user experience in mind).

Indeed, it may not be a good choice for easily turning a cli into a webservice.

## RPC!

RPC on the other hand may be a good fit because there would be a very little modification of the code.
Actually, the principle would be to:

1. trigger a network listener
2. receive a _procedure call with arguments_,
3. execute the function
4. send back the result

The function that holds the business logic does not need any change at all.

The drawbacks of RPCs are:

* the development language need a library that supports RPC,
* the client and the server must use the same communication protocol.

Those drawbacks have been adressed by Google. They gave to the community a polyglot RPC implementation called gRPC. 

Let me quote this from the chapter "[The Production Environment at Google, from the Viewpoint of an SRE](https://landing.google.com/sre/book/chapters/production-environment.html#our-software-infrastructure-XQs4iw)" of the SRE book:

> _All of Google's services communicate using a Remote Procedure Call (RPC) infrastructure named Stubby; an open source version, gRPC, is available. Often, an RPC call is made even when a call to a subroutine in the local program needs to be performed. This makes it easier to refactor the call into a different server if more modularity is needed, or when a server's codebase grows. GSLB can load balance RPCs in the same way it load balances externally visible services._

Sounds cool! Let's dig into gRPC!

### gRPC

We will now implement a gRPC server that will triggers the `cli.Commands`.

It will receive "orders", and depending of the expected call, it will: 

* Implements a `HelloCommand` and trigger its `Run()` functio or,
* Implements a `GoodbyeCommand` and trigger its `Run()` functio or,

We will also implement a gRPC client.

For the server and the client to communicate, they have to share the same protocol and understand each other with a contract.
_Protocol Buffers (a.k.a., protobuf) are Google's language-neutral, platform-neutral, extensible mechanism for serializing structured data_ 
Even if it's not mandatory, gRPC is usually used with the _Protocol Buffer_. 

So, first, let's implement the _contract_ with/in _protobuf_!

### The protobuf contract

The protocol is described with a simple text file and a specific DSL. Then there is a compiler that serialize the description and turns it into a contract that can be understood by the targeted language.

Here is a simple definition that match our need:

{{< highlight protobuf >}}
syntax = "proto3";

package myservice;

service MyService {
    rpc Hello (Arg) returns (Output) {}
    rpc Goodbye (Arg) returns (Output) {}
}

message Arg {
    repeated string args = 1;
}

message Output {
    int32 retcode = 1;
}
{{</highlight >}}

Here is the English description of the contract:

----
Let's take a service called _MyService_. This service provides to actions (commands) remotely:

* _Hello_ 
* _Goodbye_

Both takes as argument an object called _Arg_ that contains an infinite number of _string_ (this array is stored in a field called _args_).

Both actions returns an object called _Output_ that returns an integer.

----

The specification are clear enough to code a server and a client. But the string implementation may differ from a language to another.
You may now understand why we need to "compile" the file.
Let's generate a definition suitable for the go language:

`protoc --go_out=plugins=grpc:. myservice/myservice.proto`

_Note_ the definition file has been placed into a subdirectory `myservice`

This command generates a `myservice/myservice.pb.go` file. This file is part of the `myservice` package, **as specified in the myservice.proto**.

The package myservice holds the "contract" translated in `go`. It is full of interfaces and holds helpers function to easily create a server and/or a client.
Let's see how.

### The implementation of the "contract" into the server

Let's go back to the roots and read the doc of gRPC. In the [gRPC basics -  go](https://grpc.io/docs/tutorials/basic/go.html) tutorial is written:

_To build and start a server, we:_

1. _Specify the port we want to use to listen for client requests..._
2. _Create an instance of the gRPC server using grpc.NewServer()._
3. *__Register our service implementation with the gRPC server.__*
4. _Call Serve() on the server with our port details to do a blocking wait until the process is killed or Stop() is called._

Let's decompose the third step.

#### "service implementation"
In the `myservice/myservice.pb.go` file is defined an interface for our service.

{{< highlight go >}}
type MyServiceServer interface {
      // Sends a greeting
      Hello(context.Context, *Arg) (*Output, error)
      Goodbye(context.Context, *Arg) (*Output, error)
}
{{</highlight >}}

To create a "service implementation" in our "cli" utility, we need to create any structure that implements the Hello(...) and Goodbye(...) methods.
Let's call our structure `grpcCommands`:

{{< highlight go >}}
package main

...
import "myservice"
...

type grpcCommands struct {}

func (g *grpcCommands) Hello(ctx context.Context, in *myservice.Arg) (*myservice.Output, error) {
    return &myservice.Output{int32(0)}, err
}
func (g *grpcCommands) Goodbye(ctx context.Context, in *myservice.Arg) (*myservice.Output, error) {
    return &myservice.Output{int32(0)}, err
}
{{</ highlight >}}

_Note_: *myservice.Arg is a structure that holds an array of string named Args. It corresponds to the `proto` definition exposed before.

#### "service registration"

As written in the doc, we need to register the implementation.
In the generated file `myservice.pb.go`, there is a `RegisterMyServiceServer` function.
This function is simply an autogenerated wrapper around the [`RegisterService`](https://godoc.org/google.golang.org/grpc#Server.RegisterService) method of the gRPC [`Server`](https://godoc.org/google.golang.org/grpc#Server) type.

This method takes two arguments: 

* An instance of the gRPC server
* the implementation of the contract.

The 4 steps of the documentation can be implemented like this:

{{< highlight go >}}
listener, _ := net.Listen("tcp", "127.0.0.1:1234")
grpcServer := grpc.NewServer()
myservice.RegisterMyServiceServer(grpcServer, &grpcCommands{})
grpcServer.Serve(listener)
{{</ highlight >}}

So far so good... The code compiles, but does not perform any action and always return 0.

#### Actually calling the `Run()` method

Now, let's use the `grpcCommands` structure as a bridge between the `cli.Command` and the grpc service.

What we will do is to embed the `c.Commands` object inside the structure and trigger the appropriate objects's `Run()` method from the corresponding gRPC procedures.

So first, let's embed the `c.Commands` object.

{{< highlight go >}}
type grpcCommands struct {
      commands map[string]cli.CommandFactory
}
{{</ highlight >}}

Then change the `Hello` and `Goodbye` methods of `grpcCommands` so they trigger respectively:

* `HelloCommand.Run(args)`
* `GoodbyeCommand.Run(args)`

with `args` being the array of string passed via the `in` argument of the protobuf.

as defined in `myservice.Arg.Args` (the protobuf compiler has transcribed the `repeated string args` argument into a filed `Args []string` of an type `Arg`. 

{{< highlight go >}}
func (g *grpcCommands) Hello(ctx context.Context, in *myservice.Arg) (*myservice.Output, error) {
      runner, err := g.commands["hello"]()
      if err != nil {
            return int32(0), err
      }
      ret = int32(runner.Run(in.Args))
      return &myservice.Output{int32(ret)}, err
}
func (g *grpcCommands) Goodbye(ctx context.Context, in *myservice.Arg) (*myservice.Output, error) {
      runner, err := g.commands["goodbye"]()
      if err != nil {
            return int32(0), err
      }
      ret = int32(runner.Run(in.Args))
      return &myservice.Output{int32(ret)}, err
}
{{</ highlight >}}

Let's factorize a little bit and create a wrapper (that will be useful in the next section):

{{< highlight go >}}
func wrapper(cf cli.CommandFactory, args []string) (int32, error) {
      runner, err := cf()
      if err != nil {
            return int32(0), err
      }
      return int32(runner.Run(in.Args)), nil
}

func (g *grpcCommands) Hello(ctx context.Context, in *myservice.Arg) (*myservice.Output, error) {
      ret, err := wrapper(g.commands["hello"])
      return &myservice.Output{int32(ret)}, err
}
func (g *grpcCommands) Goodbye(ctx context.Context, in *myservice.Arg) (*myservice.Output, error) {
      ret, err := wrapper(g.commands["goodbye"])
      return &myservice.Output{int32(ret)}, err
}
{{</ highlight >}}

Now we have everything needed to turn our cli into a gRPC service. With a little bit of plumbing, the code compiles and the service runs.
The full implementation of the service can be found [here](https://github.com/owulveryck/cli-grpc-example/blob/master/server/main.go).

## A very quick client

The principle is the same for the client. All the needed methods are autogenerated and wrapped by the `protoc` command.

The steps are:

1. create a network connection to the gRPC server (with TLS)
2. create a new instance of myservice'client
3. Actually call a function and get a result

for example:

{{< highlight go >}}
conn, _ := grpc.Dial("127.0.0.1:1234", grpc.WithInsecure())
defer conn.Close()
client := myservice.NewMyServiceClient(conn)
output, err := client.Hello(context.Background(), &myservice.Arg{os.Args[1:]})
{{</ highlight >}}

_Note_: By default, gRPC requires some TLS. I have specified the `WithInsecure` option because I am running on the local loop and it is just an example. Don't do that in production.

# Going further

Normally, Unix tools should respect a [certain philosophy](http://www.faqs.org/docs/artu/ch01s06.html) such as:

<center>**Rule of Silence: When a program has nothing surprising to say, it should say nothing.**</center>

Anyway, we all know that tools are verbose, so let's add a feature that sends the content of stdout and stderr back to the client. (And anyway, we are implementing a service greeting. It would be useless if it was silent :))

## stdout / stderr

What we want to do is to change the output of the commands. 
Therefore, we simply add two more fields to the `Output` object in the protobuf definition:
{{< highlight protobuf >}}
message Output {
    int32 retcode = 1;
    bytes stdout = 2;
    bytes stderr = 3;
}
{{</highlight >}}

The generated file contains the following definition for `Output`:

{{< highlight go >}}
type Output struct {
      Retcode int32  `protobuf:"varint,1,opt,name=retcode" json:"retcode,omitempty"`
      Stdout  []byte `protobuf:"bytes,2,opt,name=stdout,proto3" json:"stdout,omitempty"`
      Stderr  []byte `protobuf:"bytes,3,opt,name=stderr,proto3" json:"stderr,omitempty"`
}
{{</highlight >}}

TODO: Complete here!
{{< highlight go >}}
func wrapper(cf cli.CommandFactory, args []string) (int32, []byte, []byte, error) {
	var ret int32
	oldStdout := os.Stdout // keep backup of the real stdout
	oldStderr := os.Stderr

	// Backup the stdout
	r, w, err := os.Pipe()
        //...
	re, we, err := os.Pipe()
        //...
	os.Stdout = w
	os.Stderr = we

	runner, err := cf()
        //...
	ret = int32(runner.Run(args))

	outC := make(chan []byte)
	errC := make(chan []byte)
	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.Bytes()
	}()
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, re)
		errC <- buf.Bytes()
	}()

	// back to normal state
	w.Close()
	we.Close()
	os.Stdout = oldStdout // restoring the real stdout
	os.Stderr = oldStderr
	stdout := <-outC
	stderr := <-errC
	return ret, stdout, stderr, nil
}
{{</highlight >}}


## Terraform ?

$$\frac{\partial terraform}{\partial cli} + grpc^{protobuf} = \mu service(terraform)$$ [^3]
 
[^3]: What I mean is that we are going to derivate terraform by altering the "cli" module, add it a pinch of grpc powered by protobuf, and it will give us a beautiful terraform microservice. I know, this mathematical formulae comes from nowhere. But I simply like the beautifulness of this language. (I would have been damned by my math teachers because I have used the mathematical language to describe something that is not mathematical... Would you please forgive me, gentlemen :)


