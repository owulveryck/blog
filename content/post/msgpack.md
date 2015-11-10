+++
date = "2015-11-10T08:56:36+01:00"
draft = true
title = "a IaaS-like RESTfull API based on microservices"

+++

# Absract

Recently, I've been looking at the principles of a middleware layer and especially on how [API would glue a system](http://insertpulereference).

I've also seen this excellent video made by [Mat Ryer](http://reference) about how to code an API in GO and why go would be the perfect language to code such a portal.

The problem I'm facing is that in the organization I'm working for, the developments are heterogenous and therefore you can find *ruby* teams as well as *python* team and myself as a *go* team (That will change in the future anyway)
The point is that I would like my middleware to serve as an entry point to the services provided by the departement.

We (as a department) would then be able to present the interface via, for example, a [swagger](http://swagger.io) like interface, take care of the API and do whatever RPC to any submodule.

# An example: a IAAS like interface

Let's consider a node compute lifecycle.
What I want to do is :

* create a node
* update a node (maybe)
* delete a node
* get the status of the node

## The backend

The backend is whatever service able to create a node, suchs as openstack, vmware vcac, juju, or whatever. Thoses services usually provide RESTfull API.

What I've seen in my experience, is that usually, the API are given with a library in whatever modern language. This aim to simplify the developpement of the clients.
Sometimes this library may also be developped by an internal team that will take care of the maintenance.

## The library

In my example, we will consider that the library is a simple _gem_ file developped in ruby. Therefore, our service will be a simple server that will get RPC calls, call the good method in the _gemfile_ 
and that will, _in fine_ transfer it to the backend.

## The RestFull API.

I will use the example described [here](http://blogpost) as a basis for my work.

## The glue: MSGPACK-RPC

There are severeal method for RPC-ing different languages. Ages ago, there was xml-rpc; then there has been json-rpc; I will use msgpack-rpc which is a binary, json base codec.
The communication between the Go Server and the ruby client will be donc over TCP via HTTP for example.

Later on, outside of the scope of this post, I may use ZMQ (as I have already blogged about 0MQ communication between thoses languages).

# The implementation

I will describe here the node creation via a POST method, and consider that the other method could be implemented in a similar way.
