---
author: Olivier Wulveryck
date: 2016-02-10T17:19:47+01:00
description: What we need is Choreography
draft: true
keywords:
- choreography
- orchestration
- topology
tags:
- choreography
- orchestration
- topology
- TOSCA
- go
- khoreia
title: What we need is choreography
topics:
- application deployement
type: post
---

I've had the oportunity to attend the configuration management camp in Gent (be) for its 2016 edition.

I really enjoyed those two days of talks, watching the folks to expose different ideas of what could be the future
of the infrastructure and deployment engineering. 
Beyond the technical demonstrations and the experience sharing, I've spotted a bunch of ideas

Among them, those that comes to me spontaneously are:

> You don't need configuration management, what you need is a description of the topology of your application - *Mark Shuttleworth* in its keynote _The magic of modeling_

> You don't need orchestration, what you need is choreography - Exposed by _Julian Dunn_ (you can find a retranscription [here on youtube](https://www.youtube.com/watch?v=kfF9IATUask))

> What we need is a new way to do configuration management - _James Shubin_, see [his blog post](https://ttboj.wordpress.com/2016/01/18/next-generation-configuration-mgmt/) which ispired me

I came back home very excited about this.

### From system configuration management...

I did system administration and engineering for years. Configuration management was the answer to the growing of the infrastructure.
Config management allowed us to

- Get the systems reliable
- Get the best efficiency possible from the infrastructure
- Maintain a low TCO
...

It was all "system centric", so the application could be deposed and run in best conditions.

### ... to application full description

A couple of years ago, maybe because of the DevOps movement, my missions were getting more and more application centric (which is good). 
Actually infrastructure has not been considered as a needed cost anymore.

Thanks to _Agility_, _DevOps_, and the emergent notion of product (as opposed to project), **Application and infrastructure are now seen as a whole**.  

Therefore, the setup of the application must not rely only on programmed configuration management tools anymore, but on its complete **topology**







Here is what I may tell you about thoses poins:

# You don't need configuration management...

## What we need is a description of the application

Mark Shuttleworth did a very cool introduction keynote.
Its purpose was to introduce `juju`; but the fact is that juju is an answer to modern needs.

## And an orchestrator to materialize the model

# Actually, what we need is choreography
