---
author: Olivier Wulveryck
date: 2015-11-17T10:05:42Z
description: Playing with openstack keystone
draft: true
keywords:
- openstack
- keystone
- authentication
tags:
- openstack
- keystone
- authentication
- REST
title: Playing with openstack Keystone
topics:
- topic 1
type: post
---

In the cloud computing, alongside of the hosting monsters such as amazon or google, there is the [Openstack Platform](https://www.openstack.org).

Openstack is not a single software, it is more a galaxy of components aim to control the infrastructure, such as hardware pools, storage, network.
The management can then be done via a Web based interface or via a bunch of RESTful API.

I would like to evalute its identity service named [keystone](http://docs.openstack.org/developer/keystone/) and use it as a AuthN and AuthZ backend for my simple_iaas example.

_Note_ : I will consider that the openstack keystone is installed. As I don't want to rewrite an installation procedure as many exists already on the web.

# My goal

My goal is to have a webservice that will protect the scopes of my IAAS. 
I may declare two users:

- One may list the nodes via a GET request
- The other one may also create and destroy nodes via POST and DELETE request

