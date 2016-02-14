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

I've had the oportunity to attend the [configuration management camp](http://cfgmgmtcamp.eu/) in Gent (_be_) for its 2016 edition.

I really enjoyed those two days of talks, watching people present different ideas of a possible future for
the infrastructure and deployment engineering. 
Beyond the technical demonstrations and the experience sharing, I've spotted a bunch of ideas

Among those, the onesthose that comes to me spontaneously are:

> You don't need configuration management, what you need is a description of the topology of your application - *[Mark Shuttleworth](http://www.markshuttleworth.com/biography)* in its keynote _The magic of modeling_

> You don't need orchestration, what you need is choreography - Exposed by _[Julian Dunn](https://www.linkedin.com/in/julian)_
(you can find a retranscription [here on youtube](https://www.youtube.com/watch?v=kfF9IATUask))

> What we need is a new way to do configuration management - _[James Shubin](https://www.linkedin.com/in/james-shubin-74a89a44)_, see [his blog post](https://ttboj.wordpress.com/2016/01/18/next-generation-configuration-mgmt/) which ispired my project [khoreia](http://github.com/owulveryck/khoreia)

I came back home very excited about this.
this post tries to expose my reflection and how I've implemented an idea (see it as a POC)t want to git clone...)t want to git clone...).
I've passed some time to learn about TOSCA, and the to code an orchestrator. 

In a first part I will expose why, according to me, the topological description of the application may be what
company needs.

Therefore, I will notice the need for orchestration tools.

Even if the concepts remains actuals, the future may be an evolution of this mechanism of central command and control. 
In the last part of this post, I will expose what I've understood of the concept of choreography so far.

Finally I will demonstrate the idea with a POC based on a developement on [the etcd product](https://github.com/coreos/etcd) from CoreOS.
(and a youtube demo for those who don't want to `git clone...`)

## Configuration management and orchestration

Configuration management has been for a long time, a goal for IT automation. 
Years ago, it allowed system engineers to control a huge park of machines while maintaining a TCO at a relatively decent level.

Over the last decade, 4 major tools have emerged and are now part of most CTO common vocabulary.

Let's take a look at the trends from 4 major tools categorized as "configuration management tools":

| Tool        | Founded in |
| ----------- |:----------:|
| Ansible     | 2012       |
| Puppet      | 2005       |
| Chef        | 2009       |
| Salt        | 2011       |

_Note_: I do not represent CFEngine because it is doesn't seem not so widely used in dotcom companies (even if it seems to be a great tool and on a certain extent the father of the others)

The "interest" for those tools as seen by google is be represented like this:

<center>
<script type="text/javascript" src="//www.google.com/trends/embed.js?hl=en&q=/m/0k0vzjb,+/m/03d3cjz,+/m/05zxlz3,+/m/0hn8c6s&date=1/2014+25m&cmpt=q&tz=Etc/GMT-1&tz=Etc/GMT-1&content=1&cid=TIMESERIES_GRAPH_0&export=5&w=700&h=350"></script>
</center>

As we can see, ansible seems to be the emerging technology. Indeed its acquisition by redhat in late 2015 may have boosted a bit the trends, but anyway, the companies that do not implement infrastructure as code may seem to prefer this tool.
Cause or consequence, Gartner has nominated Ansible as a _cool vendor_ for 2015 (according to Gartner, a Cool Vendor is an emerging and innovative vendor that has original, interesting, and unique technology with real market impact)

Why did a newcomer such as ansible did present such interest?

Beside its simplicity, Ansible is not exactly a configuration management tool, it is **an orchestrator** (see [the ansible webpage](https://www.ansible.com/orchestration))

According to [Rogger's theory](https://en.wikipedia.org/wiki/Diffusion_of_innovations) about the diffusion of innovation, and regarding the trends, I think that it is accurate to say
that the position of ansible is near the "late majority"
<center>
![Diffusion of ideas](https://upload.wikimedia.org/wikipedia/commons/thumb/0/0f/Diffusionofideas.PNG/330px-Diffusionofideas.PNG)
</center>

What does this mean ?

To me,it means that people do feel the need for orchestration, or if they don't feel it, they will thanks to Ansible. 
Via orchestration, they may feel the need for representing their product.

We are now talking about **infrastructure as data**; soon we will talk about **architecture as data**

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
(I'm talking of the application "born in the datacenter", it is different for those "born in the cloud")

Therefore, the setup of the application must not rely only on programmed configuration management tools anymore, but on its complete **representation**

# The self-sufficient application

Some times ago, I wrote article published on [pulse](https://www.linkedin.com/pulse/from-integration-self-sufficient-application-olivier-wulveryck?trk=prof-post) because I wanted to lay down on paper what I thought about the future of application deployment.
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

It describes a standard for representing a cloud application. It is written by the Oasis consortium and 
therefore most of the big brand in IT may be aware of its existence.
The promise is that if you describe any application with Tosca, it could be deployed on any plateform, with a decent __orchestrator__.

### ...and cons
But... Tosca is complex.
IT's not that simple to write a Tosca representation. The standard wants to cover all the possible cases, and according [Pareto](https://en.wikipedia.org/wiki/Vilfredo_Pareto)'s law,
I can say that 80% of the customers will only need 20% of the standard.

On top of that, Tosca is young (by now, the YAML version is still in pre-release), and I could not find any decent tool to orchestrate and deploy an application. 
Big companies claim their compliance with the standard, but actually very few of them (if any) does really implement it.

## Let's come back to orchestration (and real world)
As seen before, a Tosca file would need a tool to transform it to a real application.
This tool is **an orchestrator**.

The tool should be called __conductor__, because what is does actually is to conduct the symphony, and yet in our context the symphony is not 
represented by the topology, but by its 'score': its execution plan, and the purpose of the 'orchestrator' is to make every node to play its part
so the application symphony could be rendered in best condition of reliability and efficiency...

Wait, that was the promise of the configuration management tools, isn't it?

### The execution plan
So what is the execution plan.
An execution plan is a program. It describes exactly what needs to be done by systems.
The execution plan is deterministic.

With the description of the application, the execution plan, and the orchestration, the ultimate goal of automation seems fulfilled, indeed!
We have a complete suite of tools that allows to describe the application and architecture base on its functions and it is possible to 
generate and executes all the commands a computer **must** do to get things done.

Why do we neeed more?
Because now systems are so complex that we could not rely anymore on IT infrastructure to do exactly what we told it to.
Mark Burgess, considered by a lot of people as a visionnay, wrote a book entitled: 
[In Search of Certainty: The science of our information infrastructure](http://www.amazon.com/In-Search-Certainty-information-infrastructure/dp/1492389161)

Julian Dunn told about it in its speech, and I've started reading IT.

The conclusion is roughly: 

_as we may not rely on command and control anymore, we should the system to work on its own to reach a level of stability_

# Dancing, Choreography, Jazz ?






# khoreia

## Screencast: a little demo on distributed systems based on event on filesystems

Here is a little screencast I made as a POC.
Two machines are used relied by a VPN:

- my chromebook, linux-based at home in france;
- a FreeBSD server located in canada.

both machines are part of an etcd cluster.
The topology is composed of 8 nodes with dependencies which can be represented like this (same example as the one I used in a previous post):
<img class="img-responsive" src="/blog/assets/images/digraph1.png" alt="digraph example"/> 

The topology is described as a simple yaml file [here](https://github.com/owulveryck/khoreia/blob/master/samples/topology.yaml)

Each node is fulfilling two methods:

* Create
* Configure

And each method is implementing an interface composed of:

* `Check()` which check whether the action has been release and the "role" is ok
* `Do()` which actually implements the action

### Example
Each node will:

1. **Wait for an event**
2. call Create.Check() and Configure.Check().
3. watch for events from their dependencies
3. if an event is detected, call the appropriate Do() method

### Engine
The interfaces `Check()` and `Do()` may be implemented on different engines.

For my demo, as suggested by James I'm using a "file engine" base on iNotify (linux) and kQueue (freebsd).

The `Check()` method is watching the presence of a file. It sends the event "true" if the file is created of "false" if its deleted.
The `Do()` method actually create an empty file.

<center>
<iframe width="560" height="315" src="https://www.youtube.com/embed/l96uFQUrcp8" frameborder="0" allowfullscreen></iframe>
</center>

### Khoreia on github:

[github.com/owulveryck/khoreia](http://github.com/owulveryck/khoreia)
