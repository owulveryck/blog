---
author: Olivier Wulveryck
date: 2016-10-19T07:24:51+02:00
description: How to write a Single Page Application MVC without blowing your mind with Javascript and a Framework.
draft: true
keywords:
- key
- words
tags:
- one
- two
title: SPA with Gopherjs
topics:
- topic 1
type: post
---

# Introduction

Single page application (SPA) are a standard when dealing with mobile apps.
Unfortunatly, by now, javascript is the only programming language supported by a browser.

Therefore, to code web UI it remains a must.

## Life of an ex sysadmin who wants to code a web app: getting depressed

To make the development easier, your friend, who is "web developper" will recommand you to use a marvelous framework.
Depending on the orientation of the wind, the moon or its reading, he will encourage you to use `reactjs`, `angular`, `ember` or whatever exotic
tool.

With some recommandation from my real friends and from google, I've started an app based on [ionic](http://ionicframework.com/) which is based on [angular](https://angularjs.org/).
As I did not know anything about angular, I've watched a (very good) [introduction](https://www.youtube.com/watch?v=i9MHigUZKEM) and followed the ionic tututorial.

So far so good...

Then I implemented a SSO with facebook. I wrote a backend in go to handle the callback. Everything was working on my browser.

But... There was something wrong on the mobile phone version. A bug!

I tried to debug it, with Xcode, with Safari... The more I was searching, the more I had to dive into the framework.

I asked a friend and his first reply was: "which version of angular? Because in version 2 they have changed a lot of concepts"

That was too much.
I though that definitly this world made of javascript, frameworks, grunt, bower, gulp, npm or whatever fancy tool was not for me.
Too many work to learn something already outdated.

## Out of the depression!

Ok, I abandoned those tools. But I still wanted to code my app, and I'm not the kind of guy that easily give up.

Let's resume:

* I need a MCV
* MVC is not framework dependend
* A SPA is the good choice for mobile app
* Javacsript is mandatory

I dig a little bit and I've found this blog post: [Do you really want an SPA framework?](https://mmikowski.github.io/no-frameworks/) that leads me to "the solution": 

I will code my controller with JQuery and gopherjs!

# An example



