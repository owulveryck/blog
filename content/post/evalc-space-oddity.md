---
categories:
- coordination language
date: 2017-03-13T20:54:27+01:00
description: "Third article about writing a distributed linda interpreter"
draft: true
images:
- /assets/images/the_stars_look_different.jpg
tags:
- zygomys
- Linda
- golang
- etcd
title: Linda's evalc, a (tuple)space oddity
---

For a change, I will start with a good soundtrack

<iframe src="https://embed.spotify.com/?uri=spotify:track:72Z17vmmeQKAg8bptWvpVG&theme=white" width="280" height="80" frameborder="0" allowtransparency="true"></iframe>

----
This is my third article about the distributed coordination language Linda.

The final target of the work is to use this coordination mechanism to deploy and maintain applications based on the description of their topology (using, for example, TOSCA as a DSL).

Last time, I introduced a lisp based language (zygomys) as an embedded programing mechanism to describe the business logic.

Today I will explain how I have implemented a new _action_ in the linda language to achieve a new step: to distribute the work among different nodes.

My test scenario remains the "dining of the philosophers".

# Introducing _evalc_

Linda is a coordination language, but the language which is more than 30 years old, has not been design with the idea of running on multiple hosts.
The basic primitives of the language do not allow remote execution.

What I need is a sort of _eval_ function that would trigger the execution of the evaluation on another host instead of another goroutine.

I do not care about catching the result of the execution as it will be posted to the tuple space.
If I need more coordination between the actors of this RPC, I can encode them using the in/out mechanism of  linda.

# Implementing _evalc_


# Runtime

![Runtime screenshot](https://raw.githubusercontent.com/ditrit/go-linda/master/doc/v0.3.png)

# Conclusion



----
Credit:

The illustration has been found [here](https://www.flickr.com/photos/joebehr/23704122254)
