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

([youtube version](https://www.youtube.com/watch?v=iYYRH4apXDo) for those who are spotify-less)

----
This is my third article about the distributed coordination language Linda.

The final target of the work is to use this coordination mechanism to deploy and maintain applications based on the description of their topology (using, for example, TOSCA as a DSL).

Last time, I introduced a lisp based language (zygomys) as an embedded programing mechanism to describe the business logic.

Today I will explain how I have implemented a new _action_ in the linda language to achieve a new step: to distribute the work among different nodes.

My test scenario remains the "dining of the philosophers".

# Introducing _evalc_

Linda is a coordination language, but the language which is more than 30 years old, has not been designed with the idea of running on multiple hosts.
The basic primitives of the language do not allow remote execution.

What I need is a sort of _eval_ function that would trigger the execution of the evaluation on another host instead of another goroutine.

I do not care about catching the result of the execution as it will be posted to the tuple space.
Indeed, if more coordination between the actors of this RPC is needed, it can be encoded using the in/out mechanism of linda.

Therefore, I have decided to introduce a new primitive called _evalc_ (for eval compute... Yeah I know, I have imagination)

# Implementing _evalc_

The evalc will not trigger a function on a new host.
Instead, each participating host will run a sort of agent (actually a clone of thw zygo interpreter) that will watch a certain type of event (tainted with the evalc) and will then execute a function.

The tuple space acts like a communication channel and this implementation may be apparented to a kind of [CSP](https://en.wikipedia.org/wiki/Communicating_sequential_processes)

The _evalc_ will work exactly as its equivalent _eval_. Therefore the function declaration in go will look like this:

{{< highlight go >}}
func (l *Linda) EvalC(env *zygo.Glisp, name string, args []zygo.Sexp) (zygo.Sexp, error) {
    return zygo.SexpNull, nil
}
{{< /highlight >}}

## First attempt

# Runtime

![Runtime screenshot](https://raw.githubusercontent.com/ditrit/go-linda/master/doc/v0.3.png)

# Conclusion



----
Credit:

The illustration has been found [here](https://www.flickr.com/photos/joebehr/23704122254)


Conversation with [Jason E. Aten](https://www.linkedin.com/in/jason-e-aten-ph-d-45a31318) aka [glycerine](https://github.com/glycerine):

> Evaluating an arbitrary expression remotely will be challenging because an expression can refer to any variable in the environment, and so would theoretically require a copying of the whole environment--the heap as well as the datastack. The set of referenced variables can be hard to predict in advance. 
> 
> So more than likely, when you talk to a remote box, you will need to restrict the variables which you reference. This is typically done by establishing an RPC or remote procedure call convention where is only the parameters to the method call are conveyed to the remote host.  
> 
> However, another way you could approach this would be by adding a convention to the variables to establish that they are either remote or from the tuple space. 
> 
> For example, you could establish a convention for using a sigil on variables in those sexp that meant a variable was stored in the tuple space/etcd. The sigil system in zygomys was designed to support this kind of annotation (if you've seen Perl, then the `$` in from of varialbes is an example of a sigil). See the examples in tests/sigils.zy where the prefix `$`, `#`, and `?` create special symbols.
