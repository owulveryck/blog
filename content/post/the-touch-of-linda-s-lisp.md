---
categories:
- cloud
- distributed systems
date: 2017-02-28T20:57:38+01:00
description: "This is the second part of a series on my attempt to build a deployement language on a cloud scale"
draft: true
images:
- https://upload.wikimedia.org/wikipedia/commons/9/9d/SICP_cover.jpg
tags:
- go
- zygomys
- Lisp
- linda
title: To go and touch Linda's Lisp
---

The title is not a typo nor dyslexia. I will really talk about Lisp.

In a [previous post](/2017/02/03/linda-31yo-with-5-starving-philosophers.../index.html) I explained my will to implement the dining of the philosophers with Linda in GO.

The ultimate goal is to use a distributed and abstract language to go straight from the design to the runtime of an application.

# The problem I've faced

I want to use a GO implementation for the Linda language because a go binary is a container by itself. Therefore if I build my linda language within go, I will be able to run it easily accross the computes nodes without any more dependencies.
The point is that the Linda implementation may be seen as a go package. Therefore every implenentation of every algorithm must be coded in go. Therefore I will lack a lot of flexibility as I will need on agent per host and per algorithm. For example the binary that will solve the problem of the dining of the philosophers will only be usefull for this specific problem.

What would be nice it to use an embedded scripting language. This language would implement the Linda primitives (_in, rd, eval, out_). And the go binary would be in charge to communicate with the tuple space.

## Tuple space: _I want your Sexp_

I have though a lot about a way to encode my tuples for the tuple space.
Of course go as a lot of enconding available:

- json
- xml
- protobuf
- gob

None of them gave me entire satisfaction. The reason is that go is strongly typed. A tuple must be internally represented by an empty *interface{}* to remain flexible.
Obviously I would need to use a lot of reflexion in my code. 

So to keep it simple I though I would need a little refresh about the principles of the reflection. So I took my [book](https://books.google.fr/books/about/The_Go_Programming_Language.html?id=SJHvCgAAQBAJ) about go (I bought it when I started learning go).

In this book there is a full example about encoding and decoding [s-expression](https://en.wikipedia.org/wiki/S-expression). And what is an s-expression? A tuple! __eureka__

## Lisp/zygomys

So I though about s-expression again.... I could use the parser described in my book and that would be enough for the purpose of my test.
I could just create a package _encoding/sexpr_ and that would do the job.

But the more I was reading about s-expression, the more I was digging in list processing. 

List processing, s-expression, embedded language, functionnal programming.... That was it: I really needed a lisp based embedded language to implement my linda solution.

![xkcd 297](https://imgs.xkcd.com/comics/lisp_cycles.png)

# The POC

## Implementing the linda primitives in the REPL

## _etcd_ as a tuple space

## Implememting the algorithm in _zygomys_

## RUN!

# Conclusion and future work

