---
title: "Think like a vertex: using Go's concurrency for graph computation"
date: 2019-10-14T22:26:42+02:00
lastmod: 2019-10-14T22:26:42+02:00
draft: true
keywords: []
description: "In this article, I describe the prototype of a new computation machine for graph processing.  This machine takes its inspiration from the Pregel paradigm and uses Go's concurrency mechanism as a lever for a simple implementation."
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

So, Pregel's goal is to solve the problem of graph processing by leveraging the power used for cloud computing.
(a cluster of inexpensive machines).

Distributed programming, most of the time efficient, is nevertheless hard. But the original paper mention that:

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

eescription: "In this article, I describe the prototype of a new computation machine for graph processing.  This machine takes its inspiration from the Pregel paradigm and uses Go's concurrency mechanism as a lever for a simple implementation."
Computing machine Learning equations is dealing with of massive objects (matrices) on significant graphs. Let's try this implementation for real in Gorgonia.

# About Gorgonia

Gorgonia is a computation library written in Go.
Its goal it to facilitate machine learning in this language.

The principle of Gorgonia is:

* it gives the primitives to build an expression graph;
* it implements "machines" to compute the expression graph and provide the result;
* it can also do automatic differentiation, but let's put this aside for now;

## The Expression Graph

The vertices of the ExprGraph are Go structures called [`Node`](https://godoc.org/gorgonia.org/gorgonia#Node).

A node carries a [`Value`](https://godoc.org/gorgonia.org/gorgonia#Value) and an [`Op`eration](https://godoc.org/gorgonia.org/gorgonia#Op).

The Operation is an object with a special method Do:

```go
Do(...Value) (Value, error)
```

It looks possible to write a computation engine on the principle we've evaluated before.

* create a channel of `Value` for every edge of the graph.
* create a goroutine for every node of the graph
  * The goroutine takes the input from the channels that reach the node
  * The goroutine executes the `Do` statement of the operation
  * The goroutine write the output value to every channel issued from the current node

## How? Gorgonia's VM

Gorgonia describes a [VM](https://godoc.org/gorgonia.org/gorgonia#VM) via a Go interface
From the documentation:

> VM represents a structure that can execute a graph or program.

So, the "pregel" implementation we are seeking is eventually an implementation of the VM interface.

I made such experimental implementation called `GoMachine`. It can be found in the master branch of Gorgonia.
The code and godoc are accessible [here](https://godoc.org/gorgonia.org/gorgonia/x/vm#GoMachine).

Of course there are caveats in the trivial implementation described in this post. For example, the vertex cannot read or write the IO channels
sequentially, otherwise it may end with a deadlock.
The GoMachine takes care of that. But the principle is not different from what has been described here; nor the code is mode complicated.

I've used this machine for onnx-go, and successfully run some models with very good performances with it.

Do not hesitate to give it a try.

# Conclusion

Some of the coolest features of the Go language are its simplicity from the development process to the distribution of the final binary.

This is why I started using Go for machine learning in first place. It was the easiest way to run a neural net into production without worrying about
dependencies.

But the real power of the Go language goes far beyond those principle.
Actually concurrency is a key point in this. Using this lever can really improve efficiency while keeping things simple.
The concurrent machine is a perfect example of this.

The full implementation, able to run some neural net models such as Resnet or (tiny)Yolo is less than 200 lines of code.

I do not want to stop the experiment here. Some stuffs I'd like to see are:

* The gradient computation to the machine
* The usage of CUDA for the operator supporting it
* ...
* A true distributed graph computation over several machines in the cloud maybe?
