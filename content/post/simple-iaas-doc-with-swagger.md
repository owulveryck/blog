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
