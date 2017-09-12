---
categories:
date: 2017-09-12T13:28:36+02:00
description: ""
draft: false
images:
- /assets/images/terraformcli.png
tags:
title: Introducing hip terraform
---

In a previous post, I did some experiments with gRPC, protocol buffer and Terraform.
The idea was to transform the "terraform" cli tool into a micro-service thanks to gRPC.

This post is the second part of the experiment. I will go deeper in the code an see if it is possible
to create a brand new utility, without hacking terraform. The idea is to import some of the packages that compose the binary
and create my own service based on gRPC.

# The terraform structure

## The cli implementation

## The `commands` package

# Creating a new binary

## The gRPC contract

## Fulfilling the contract

## About the UI

### A custom UI

# About concurrency: a new _push_ command

# going further...

## Implementing a micro-service of backend

# Hip is _cooler than cool_: Introducing _nhite_[^1]

I have packed everything into an organization called nhite.
There is still a lot to do, but I really think that this could make sense to create a community. I may give a product by the end, or go in my attic of dead project.
Anyway, so far I've had a lot of fun!

[^1]: [hip definition on wikipedia](https://en.wikipedia.org/wiki/Hip_(slang))


