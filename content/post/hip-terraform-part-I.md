---
categories:
date: 2017-09-12T13:28:36+02:00
description: ""
draft: false
images:
- https://nhite.github.io/images/logo.png
tags:
title: Introducing hip terraform
---

In a previous post, I did some experiments with gRPC, protocol buffer and Terraform.
The idea was to transform the "terraform" cli tool into a micro-service thanks to gRPC.

This post is the second part of the experiment. I will go deeper in the code an see if it is possible
to create a brand new utility, without hacking terraform. The idea is to import some of the packages that compose the binary
and create my own service based on gRPC.

# The terraform structure

Terraform is a binary utility written in `go`.
The `main` package resides in the root directory of the `terraform` directory.
As usual with go projects, all other subdirectories are different modules.

The whole business logic of terraform is coded into the subpackages. The main is simply an envelop for kickstarting the utility (env variables etc.) and to init the command line.

### The cli implementation

The command line flags are instanciated by Mitchell Hashimoto's cli package.
As explained in the previous post, this cli package is calling a specific function for every action.

### The _command_ package

Every single action is fulfilling the `cli.Command` interface and is implemented in the [`command`](https://godoc.org/github.com/hashicorp/terraform/command) subpackage.

# Creating a new binary

The idea is not to hack any if the packages of terraform to allow an easier maintenance of my code. 
In order to create a custom service, I will instead implement a new utility; therefore a new `main` package.
This package will implement a gRPC server. This server will implement wrappers around the functions declared in the `terraform.Command` package.

## The gRPC contract

In order to create a gRPC server, we need a service definition.
To keep it simple for now, I will use the one that has been defined in the previous post ([cf the section: Creating the protobuf contract](https://blog.owulveryck.info/2017/09/02/from-command-line-tools-to-microservices---the-example-of-hashicorp-tools-terraform-and-grpc.html#creating-the-protobuf-contract)).

## Fulfilling the contract

## About the UI

### A custom UI

# About concurrency: a new _push_ command

# going further...

## Implementing a micro-service of backend

# Hip[^1] is _cooler than cool_: Introducing _nhite_

## The organisation structure

### Demo?

[^1]: [hip definition on wikipedia](https://en.wikipedia.org/wiki/Hip_(slang))

I have packed everything into an organization called nhite.
There is still a lot to do, but I really think that this could make sense to create a community. I may give a product by the end, or go in my attic of dead project.
Anyway, so far I've had a lot of fun!



