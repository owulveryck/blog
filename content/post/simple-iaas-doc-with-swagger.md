---
author: Olivier Wulveryck
date: 2015-11-11T14:24:43+01:00
description: Experience with swagger-ui as a documentation tool for the simple iaas api
draft: true
tags:
- swagger
- api
- documentation
title: Simple IaaS API documentation with swagger
type: post
---

In a [previous post](http://blog.owulveryck.info/2015/11/10/iaas-like-restfull-api-based-on-microservices/) I have explained how to develop a very simple API server.

Without the associated documentation, the API will be useless. Let's see how we can use [swagger-ui](https://github.com/swagger-api/swagger-ui) 
in this project to generate a beautiful documentation.

*Note* I'm blogging and experimenting, of course, in the "real" life, it's a lot better to code the API interface before implementing the middleware.

# About Swagger

Swagger is a framework. On top of the the swagger project is composed of several tools.

The entry point is to write the API interface using the [Swagger Formal Specification](http://swagger.io/specification/). I will the use the [swagger-ui](https://github.com/swagger-api/swagger-ui) to display the documentation.
The swagger-ui can be modified and recompiled, but I won't do it (as I don't want to play with nodejs). Instead I will rely on the "dist" part which can be used "as-is"


# Defining the API interface with Swagger

## Header and specification version:

Swagger comes with an editor which can be used [online](http://editor.swagger.io/#/).

I will use swagger spec 2.0, as I don't see any good reason not to do so. Morover, I will describe the API using the `YAML` format instead of the JSON format to be human-friendly.

Indeed, in my `YAML` squeleton the header of my specs will then look like thi:

```yaml
swagger: '2.0'
info:
  version: 1.0.0
    title: 'Very Simple IAAS'
```

## The node creation: a POST method
Let's document the Node creation (as it is the method that we have implemented before).

The node creation is a `POST` method, that produces a JSON in output with the request ID of the node created.

The responses code may be:

* 202 : if the request has been taken in account
* 400 : when the request is not formatted correctly
* 500 : if any unhaldled exception occured
* 502 : if the backend is not accessible (etiher the RPC server or the backend)

So far, the YAML spec will look like:
```yaml
paths:
  /v1/nodes:
    post:
      description: Create a node
      produces:
        - application/json
      responses:
        202:
          description: A request ID.
        400:
          description: |
            When the request is malformated or when mandatory arguments are missing
        500:
          desctiption: Unhandled error
        502:
          description: Execution backend not available
```

So far so good, let's continue with the input payload.


