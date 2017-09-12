---
categories:
date: 2017-09-12T13:28:36+02:00
description: "This is a second part of the last article. I now really dig into terraform. This article will explain how to use the Terraform sub-packages in order to create a brand new binary that acts as a gRPC server instead of a cli."
draft: false
images:
- https://nhite.github.io/images/logo.png
tags:
title: Terraform is hip... Introducing Nhite
---

In a previous post, I did some experiments with gRPC, protocol buffer and Terraform.
The idea was to transform the "terraform" cli tool into a micro-service thanks to gRPC.

This post is the second part of the experiment. I will go deeper in the code and see if it is possible
to create a brand new utility, without hacking terraform. The idea is to import some of the packages that compose the binary
and create my own service based on gRPC.

# The terraform structure

Terraform is a binary utility written in `go`.
The `main` package resides in the root directory of the `terraform` directory.
As usual with go projects, all other subdirectories are different modules.

The whole business logic of terraform is coded into the subpackages. The "`main`" package is simply an envelop for kick-starting the utility (env variables, etc.) and to initiate the command line.

### The cli implementation

The command line flags are instantiated by Mitchell Hashimoto's cli package.
As explained in the previous post, this cli package is calling a specific function for every action.

### The _command_ package

Every single action is fulfilling the `cli.Command` interface and is implemented in the [`command`](https://godoc.org/github.com/hashicorp/terraform/command) subpackage.
Therefore, every "action" of terraform has a definition in the command package and the logic is coded into a `Run(args []string) int` method (see the [doc of the Command interface for a complete definition](https://godoc.org/github.com/mitchellh/cli#Command).

# Creating a new binary

The idea is not to hack any of the packages of terraform to allow an easier maintenance of my code. 
In order to create a custom service, I will instead implement a new utility; therefore a new `main` package.
This package will implement a gRPC server. This server will implement wrappers around the functions declared in the `terraform.Command` package.

For the purpose of my poc, I will only implement three actions of terraform:

* `terraform init`
* `terraform plan`
* `terraform apply`

## The gRPC contract

In order to create a gRPC server, we need a service definition.
To keep it simple, let's consider the contract defined in the previous post ([cf the section: Creating the protobuf contract](https://blog.owulveryck.info/2017/09/02/from-command-line-tools-to-microservices---the-example-of-hashicorp-tools-terraform-and-grpc.html#creating-the-protobuf-contract)).
I simply add the missing procedure calls:

{{< highlight protobuf >}}
syntax = "proto3";

package pbnhite;

service Terraform {
    rpc Init (Arg) returns (Output) {}
    rpc Plan (Arg) returns (Output) {}
    rpc Apply (Arg) returns (Output) {}
}

message Arg {
    repeated string args = 2;
}

message Output {
    int32 retcode = 1;
    bytes  stdout = 2;
    bytes stderr = 3;
}
{{</ highlight >}}

## Fulfilling the contract

As described previoulsy, I am creating a `grpcCommand` structure that will have the required methods to fulfill the contract:

{{< highlight go >}}
type grpcCommands struct {}

func (g *grpcCommands) Init(ctx context.Context, in *pb.Arg) (*pb.Output, error) {
    ....
}
func (g *grpcCommands) Plan(ctx context.Context, in *pb.Arg) (*pb.Output, error) {
    ....
}
func (g *grpcCommands) Apply(ctx context.Context, in *pb.Arg) (*pb.Output, error) {
    ....
}
{{</ highlight >}}

In the previous post, I have filled the `grpcCommand` structure with a `map` filled with the command definition.
The idea was to keep the same cli interface.
As we are now building a completely new binary with only a gRPC interface, we don't need that anymore.
Indeed, there is still a need to execute the `Run` method of every terraform command.

Let's take the example of the Init command. 

Let's see the definition of the command by looking at the [godoc](https://godoc.org/github.com/hashicorp/terraform/command#InitCommand):

{{< highlight go >}}
type InitCommand struct {
    Meta
    // contains filtered or unexported fields
}
{{</ highlight >}}

This command holds a substructure called `Meta`. `Meta` is defined [here](https://godoc.org/github.com/hashicorp/terraform/command#Meta) and holds _the meta-options that are available on all or most commands_. Obviously we need a Meta definition in the command to make it work properly.

For now, let's add it to the `grpcCommand` globally, and we will see later how to implement it.

Here is the gRPC implementation of the contract:

{{< highlight go >}}
type grpcCommands struct {
    meta command.Meta
}

func (g *grpcCommands) Init(ctx context.Context, in *pb.Arg) (*pb.Output, error) {
    // ...
    tfCommand := &command.InitCommand{
        Meta: g.meta,
    }
    ret := int32(tfCommand.Run(in.Args))
    return &pb.Output{ret, stdout, stderr}, err
}
{{</ highlight >}}

## How to initialize the _grpcCommand_  object

Now that we have a proper `grpcCommand` than can be registered to the grpc server, let's see how to create an instance.
As the grpcCommand only contains one field, we simply need to create a `meta` object.

Let's simply copy/paste the code done in terraform's main package and we now have:

{{< highlight go >}}
var PluginOverrides command.PluginOverrides
meta := command.Meta{
    Color:            false,
    GlobalPluginDirs: globalPluginDirs(),
    PluginOverrides:  &PluginOverrides,
    Ui:               &grpcUI{},
}
pb.RegisterTerraformServer(grpcServer, &grpcCommands{meta: meta})
{{</ highlight >}}

According to the comments in the code, the `globalPluginDirs()` _returns directories that should be searched for
globally-installed plugins (not specific to the current configuration)_.
I will simply copy the function into my main package

## About the UI

In the example cli that I developped in the previous post, what I did was to redirect stdout and stderr to an array of bytes, in order to capture it and send it back to a gRPC client.
I noticed that this was not working with Terraform.
This is because of the UI!
UI is an interface whose purpose is to get the output stream and write it down to a specific io.Writer.

Our tool will need a custom UI.

### A custom UI

As UI is an interface ([see the doc here](https://godoc.org/github.com/mitchellh/cli#Ui)), it is easy to implement our own. Let's define a structure that holds two array of bytes called `stdout` and `stderr`. Then let's implement the methods that will write into this elements:

{{< highlight go >}}
type grpcUI struct {
    stdout []byte
    stderr []byte
}

func (g *grpcUI) Output(msg string) {
    g.stdout = append(g.stdout, []byte(msg)...)
}
{{</ highlight>}}

_Note 1_: I omit the methods `Info`, `Warn`, and `Error` for brevity.

_Note 2_: For now, I do not implement any logic into the `Ask` and `AskSecret` methods. Therefore, my client will not be able to ask something. But as gRPC is bidirectional, it would be possible to implement such an interaction.

Now, we can instantiate this UI for every call, and assing it to the meta-options of the command:

{{< highlight go >}}
var stdout []byte
var stderr []byte
myUI := &grpcUI{
    stdout: stdout,
    stderr: stderr,
}
tfCommand.Meta.Ui = myUI
{{</ highlight >}}

So far, so good: we now have a new terraform binary, that is working via gRPC with a very little bit of code.

# What did we miss?

## Multi-stack
It is fun but not usable for real purpose because the server needs to be launched from the directory where the tf files are... 
Therefore the service can only be used for one single terraform stack... Come on!

Let's change that and pass as a parameter of the RPC call the directory where the server needs to work. It is as simple as adding an extra argument to the `message Arg`:

{{< highlight protobuf >}}
message Arg {
    string workingDir = 1;
    repeated string args = 2;
}
{{</ highlight >}}

and then, simply do a `change directory` in the implementation of the command:

{{< highlight go >}}
func (g *grpcCommands) Init(ctx context.Context, in *pb.Arg) (*pb.Output, error) {
    err := os.Chdir(in.WorkingDir)
    if err != nil {
        return &pb.Output{int32(0), nil, nil}, err
    }
    tfCommand := &command.InitCommand{
        Meta: g.meta,
    }
    var stdout []byte
    var stderr []byte
    myUI := &grpcUI{
        stdout: stdout,
        stderr: stderr,
    }
    ret := int32(tfCommand.Run(in.Args))
    return &pb.Output{ret, stdout, stderr}, err
}
{{</ highlight >}}

## Implementing a new _push_ command

Now that we have a terraform service how can it be used by a client.
Assume that this service is deployed on a remote machine. We do not want everybody to log into the host to write their stack in the local filesystem.


# going further...

## Implementing a micro-service of backend

# Hip[^1] is _cooler than cool_: Introducing _Nhite_

## The organisation structure

### Demo?

[^1]: [hip definition on wikipedia](https://en.wikipedia.org/wiki/Hip_(slang))

I have packed everything into an organization called nhite.
There is still a lot to do, but I really think that this could make sense to create a community. I may give a product by the end, or go in my attic of dead project.
Anyway, so far I've had a lot of fun!



