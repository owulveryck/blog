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

> You don't need configuration management, what you need is a description of the topology of your application - *[Mark Shuttleworth](http://www.markshuttleworth.com/biography)* in its keynote _The magic of modeling_

> You don't need orchestration, what you need is choreography - Exposed by _[Julian Dunn](https://www.linkedin.com/in/julian)_
(you can find a retranscription [here on youtube](https://www.youtube.com/watch?v=kfF9IATUask))

> What we need is a new way to do configuration management - _[James Shubin](https://www.linkedin.com/in/james-shubin-74a89a44)_, see [his blog post](https://ttboj.wordpress.com/2016/01/18/next-generation-configuration-mgmt/) which ispired my project [khoreia](http://github.com/owulveryck/khoreia)

I came back home very excited about this.

### From system configuration management...

I did system administration and engineering for years. Configuration management was the answer to the growing of the infrastructure.
Config management allowed us to

- Get the systems reliable
- Get the best efficiency possible from the infrastructure
- Maintain a low TCO
...

It was all "system centric", so the application could be deposed and run in best conditions.

### ... to application's full description

A couple of years ago, maybe because of the DevOps movement, my missions were getting more and more application centric (which is good). 
Actually infrastructure has not been considered as a needed cost anymore.

Thanks to _Agility_, _DevOps_, and the emergent notion of product (as opposed to project), **Application and infrastructure are now seen as a whole**.  

Therefore, the setup of the application must not rely only on programmed configuration management tools anymore, but on its complete **representation**

# The self-sufficient application

Some times ago, I wrote a pulse article because I wanted to lay down on paper what I thought about the future of application deployment.
I've described some layers of the application.
I kept on studying, and with a some help from my colleagues and friends, I've  finally been able to put a word on those ideas I had in mind.
This word is **Topology**

## and then came TOSCA

To describe a whole application, I needed a _domain specific language_ (DSL).
All of the languages I was trying to document were by far too system centric.
Then I discovered [TOSCA](http://docs.oasis-open.org/tosca/TOSCA-Simple-Profile-YAML/v1.0/csprd01/TOSCA-Simple-Profile-YAML-v1.0-csprd01.html).
TOSCA is __THE DSL__ for representing the topology of an application.

### Pros...
What's good about Tosca is its goal:

It describes a standard for representing a cloud application. It is written by the Oasis and therefore most of the big brand in IT may be aware of its existence
The promise is that if you describe any application with Tosca, it could be deployed on any plateform, with a decent __orchestrator__.

### ...and cons
But... Tosca is complex.
IT's not that simple to write a Tosca file. The standard want to cover all the possible case, and according [Mr Pareto](https://en.wikipedia.org/wiki/Vilfredo_Pareto)'s law,
I can say that 80% of the customers will only need 20% of the standard.

On top of that, Tosca is young, and I could not find any decent tool to orchestrate and deploy an application. 
Big companies claim their compliance with the standard, but actually very few of them (if any) does really implement it.

# The need of orchestration
As seen before, a Tosca file would need a tool to transform it to a rela application.
This tool is **an orchestrator**.
