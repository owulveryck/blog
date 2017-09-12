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

### A custom UI

# About concurrency

## Implementing a new _push_ command

## Setting the working directory

# going further...

## Implementing a micro-service of backend

# Hip[^1] is _cooler than cool_: Introducing _Nhite_

## The organisation structure

### Demo?

[^1]: [hip definition on wikipedia](https://en.wikipedia.org/wiki/Hip_(slang))

I have packed everything into an organization called nhite.
There is still a lot to do, but I really think that this could make sense to create a community. I may give a product by the end, or go in my attic of dead project.
Anyway, so far I've had a lot of fun!



