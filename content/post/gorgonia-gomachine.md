---
title: "Think like a vertex: using Go's concurrency for graph computation"
date: 2019-10-14T22:26:42+02:00
lastmod: 2019-10-14T22:26:42+02:00
draft: true
keywords: []
description: ""
tags: []
categories: []
author: ""

# You can also close(false) or open(true) something for this content.
# P.S. comment can only be closed
comment: false
toc: false
autoCollapseToc: false
# You can also define another contentCopyright. e.g. contentCopyright: "This is another copyright."
contentCopyright: false
reward: false
mathjax: true
---

<!--more-->

I am often asked this question:

_Why do you do machine learning in Go?_

Of course, the main reason is that I like the language. But there are other, more generic, reasons.

In the fifth episode of the third season of [Command Line Heroes](https://www.redhat.com/en/command-line-heroes/season-3/the-infrastructure-effect),
Saron Yitbarek exposes the fact that Go's design is tidily linked to the cloud infrastructure. Indeed, the concurrency mechanism makes it super easy to write a program the
can run at scale on inexpensive machines.

And I genuinely believe that this power is underused in the data-science community.
But those are just thoughts. Only facts lead to a dispassionate debate.

In this article, I describe the prototype of a new computation machine for graph processing.
This machine takes its inspiration from the Pregel paradigm and uses Go's concurrency mechanism as a lever for a simple implementation.

# Computation model for graph processing

Efficient graph processing is actually a key to modern computing and to machine learning success.

In graph processing, [_Spark_](https://spark.apache.org/) is a reference. It is known for its abitilty to process large computation graph.

Spark is base on the [Pregel paradigm](https://kowshik.github.io/JPregel/pregel_paper.pdf), which is a system or large graph processing.
let's take a closer look at this piece of art involved behind the scene.

### About pregel

Pregel is a computation model, not an algorithm. From the original paper, Pregel's goal is to define a program as

> (...) a sequence of iterations, in each of which a vertex can
> receive messages sent in the previous iteration, send messages to other vertices, and modify its own state and that of
> its outgoing edges or mutate graph topology.
> This vertex-centric approach is flexible enough to express a broad set of algorithms

On top of this definition, those [notes from the CME 323](https://stanford.edu/~rezab/classes/cme323/S15/notes/lec8.pdf)
from Stanford University gives a useful résumé of what Pregel is:

> Pregel is essentially a message-passing interface constrained to the edges of a graph. The idea
> is to ”think like a vertex” - algorithms within the Pregel framework are algorithms in which the
> computation of state for a given node depends only on the states of its neighbours.

### Pregel in Go ?

So, Pregel is a computation graph designed to solve the problem of graph processing by leveraging the power of cloud computing
(I mean, using a cluster of inexpensive machines).

Distributed programming can be hard, but the original paper mention that:

> The model (...) implied synchronicity makes reasoning about programs easier.

And Go's concurrency mechanism makes it (super) easy to synchronize concurrent routines.

Let's draw a very basic computation graph.

As a support, let's consider this equation (which is a typical layer of a neural network):

$$f(X) = \sigma(W \cdot X+b)$$

Let's turn it into something more "functional":
$$f(X) = \sigma(add(mul(W,X),B))$$

Now we can express it into a graph:

<center>
![graph](/assets/pregel/graph2.png)
</center>

Now, let's think _like a vertex_:

- I am the `mul` node:
    - I am waiting for `X` and `W` to tell me their values through channel `A` and `B`
    - I am computing the value
    - I am sending the result on channel `C`
- I am the `add` node:
    - I am waiting for `mul` and `b` to tell me their values through channel `C` and `D`
    - I am computing the value
    - I am sending the result on channel `E`
...

### Trivial implementation
Implementing this in Go is fairly easy.

Consider the message as a `float64` value that will flow through channels of communication.
A vertex is then a function that reads from the channels apply its content body and write its result to the output channel:

```go
type message float64
type Vertex func(output chan<- message, input ...<-chan message)
```

The vertex implementation is straightforward:

```go
add := func(output chan<- message, input ...<-chan message) {
        a := <-input[0]
        b := <-input[1]
        output <- message(a + b)
}
```

And the original equation is encoded like this:
```go
A <- message(1.0)
B <- message(1.0)
D <- message(1.0)

mul(C, A, B)
add(E, C, D)
sigma(output, E)

fmt.Println(<-output)
```

Running this code prints `0.8807970779778823` (see the full code [here](https://gist.github.com/owulveryck/b1255077d7e1d940f9cc472bc69ef733))

### Adding concurrency

The main issue in the trivial implementation is that it is not possible to set the values after the operators application. Doing this would lead to a deadlock:

```go
mul(C, A, B)
A <- message(1.0)
```

because assignation to `A` will happen after `mul`'s execution, but `mul` is waiting for a value in channel `A`.

Solution to this problem can be solve thanks to go-routines like this:

```go
go mul(C, A, B)
go add(E, C, D)
go sigma(output, E)

A <- message(1.0)
B <- message(1.0)
D <- message(1.0)
```

Now every vertex runs in a goroutine and the deadlock's gone. Even better, this mechanism has implicit synchronization.
Therefore, computing this graph is more efficient with this mechanism that coding it sequentially because the `mul` operation will be computed in parralel:

<center>
![graph](/assets/pregel/graph3.png)
</center>

Let's do a simple benchmark to validate this hypothesis.
I wrote two simple bench functions:

* concurrent
* sequential

(I do not copy the code for clarity, but you can find it [on gist](https://gist.github.com/owulveryck/b1255077d7e1d940f9cc472bc69ef733#file-pregel_test-go))


```text
benchcmp /tmp/sequential /tmp/concurrent
benchmark           old ns/op     new ns/op     delta
BenchmarkTest-4     693           4253          +513.71%
```

This result shows that the sequential implementation is way faster. This is understandable because the operation is trivial and highly optimized, and there is an overhead
induced by the concurrency mechanism.

Let's put bias and add a 50 microseconds sleep the `mul` and `add` operations to simulate a real-world and less trivial computation.

```go
mul = func(output chan<- message, input ...<-chan message) {
        a := <-input[0]
        b := <-input[1]
        time.Sleep(50 * time.Microsecond)
        output <- message(a * b)
}
```

With a ballast of 50 microseconds, the concurrent implementation is 22% faster.

```text
benchcmp /tmp/sequential50 /tmp/concurrent50
benchmark           old ns/op     new ns/op     delta
BenchmarkTest-4     276892        213892        -22.75%
```

Computing machine Learning equations is dealing with of massive objects (matrices) on significant graphs. Let's try this implementation for real in Gorgonia.

# About Gorgonia

