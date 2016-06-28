---
author: Olivier Wulveryck
date: 2016-06-23T15:32:54+02:00
description: Playing with websocket for a dynamic presentation.
draft: true
tags:
- websocket
- bootstrap
- golang
- reveal.js 
- Javascript
- D3.js
title: Websockets, Reveal.js, D3 and GO for a dynamic keynote
topics:
- topic 1
type: post
---

# the goal
As all my peers, I have the opportunity to talk about different technological aspects.
As all my peers, I'm asked to present a bunch of slides (powerpoint or keynote, or whatever).

In this post I won't dig into what's good or not to put in a presentation, and if that's what interest you, I 
recommend you to take a look at [Garr Reynold's tips and tricks](http://reference).

_Steve Jobs_ said:

> People who knows what they're talking about don't need PowerPoint

(actually it has been reported in its book) (TODO: references)

As an attendee I tend to agree; usually PowerPoints are boring and they hardly give any interest besides for the writer to say "hey look, I've worked for this presentation".

Indeed, they are a must. So for my next presentation I thought: 

wouldn't it be nice to use this wide display area to make the presentation more interactive.
One of the key point in communication is to federate people. So what if people could get represented for real in the presentation.

## how to: the architecture 

Obviously I cannot use conventional tools, such as PowerPoint, Keynote, Impress, google slides and so.
I need something that I can program; something that can interact with a server, and something that is not a console so I can get
fancy and eye-candy animations.

### The basic

[reveal.js](http://reference) is an almost perfect candidate:

* it is a framework written in JavaScript therefore, I can easily ass code
* it's well designed
* it can be used alongside with any other JavaScript framework

### Graphs, animations, etc...

A good presentation has animations, graphs, diagrams, and stuffs that cannot be expressed simply with words.
I will interact with the audience. I will explain how later, but anyway they will send me some data.
I could process them in whatever server-side application (php, go-template base, python) but I have the feeling that's not 
the idiomatic way of doing modern web content. Actually, I would need anyway to deal with device (mobile, desktop), screen size,
browser... So what's best, I think, is to get the data on the client side and process it via Javascript.

[Data Driver Document (D3)](http://reference) is the framework I will use to process and display the data I will get from the client.

### The attendees 

If I want the attendees to participate they need a device, to act as a client.
About all people I will talk to have a smartphone; that is what I will use. 

It has two advantages:

* it is their own device, I looks more realistic and unexpected: therefore I would get a better reception of the message I'm trying to pass.
* it usually has a Webkit based web browser with a decent Javascript engine.

I won't develop a native app, instead I will a webpage mobile compliant based on the [bootstrap](http://reference) framework.

### The HUB

#### A server

#### with websocket
