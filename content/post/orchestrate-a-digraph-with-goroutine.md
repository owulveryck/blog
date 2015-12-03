---
author: Olivier Wulveryck
date: 2015-12-02T14:24:21Z
description: |
    A directed graph may be represented by its adjacency matrix.
    Considering each vertice as a runable element and any edge as a dependency,
    I will describe a method to "run" the graph in a concurrent way using goalang's goroutine
draft: true
tags:
- golang
- digraph
title: Orchestrate a digraph with goroutine
type: post
---

I've read a lot about graph theory recently.
They have changed the world a lot. From the simple representation to Bayesian network via Markov chains, the applications are numerous.

Today I would like to imagine a graph as a workflow of execution. Every node would be considered as runnable. And every  edge would be a dependency.

# The use case 
If we consider this very simple graph (example taken from the french wikipedia page)

<img class="img-responsive" src="/assets/images/digraph1.png" alt="digraph example"/>

its corresponding adjacency matrix is:

 <img class="img-responsive" src="/assets/images/matrix1.png" alt="Adjacency matrix"/>

its dimension is 8x8

For the lab, I will consider that each node has to do a simple task which is to wait for a random number of millisecond (such as Rob Pike's _boring_ function, see references)

# Let's GO

## How will it work

Every node will be run in a `goroutine`. That is a point. But how do I deal with concurrency ?

Every single goroutine will be initially launched and then wait for an information.

It will have an input communication channel, and a _conductor_ will feed this channel with enough information for the goroutine to decides whether it should run or not. 
This information is simply the adjacency matrix up-to-date. That means that is a node is done, its value is set to zero.

Every goroutine will then check in the adjacency matrix, whether it has predecessor and therefor will execute the step or not.

Once the execution of task is over, the goroutine will then feed another channel to tell the conductor that its job is done. and then the conductor will broadcast the information.

* __(1)__ The conductors feed the nodes with the matrix
<a href="/assets/orchestrate-a-digraph-with-goroutine/digraph_step1.dot"><img class="img-responsive" src="/assets/images/digraph_step1.png" alt="digraph example"/></a>
* __(2)__ Every node get the data and analyse the matrix
<a href="/assets/orchestrate-a-digraph-with-goroutine/digraph_step2.dot"><img class="img-responsive" src="/assets/images/digraph_step2.png" alt="digraph example"/> </a>
* __(3)__ Nodes 3, 5 and 7 have no predecessor, they can run
<a href="/assets/orchestrate-a-digraph-with-goroutine/digraph_step3.dot"><img class="img-responsive" src="/assets/images/digraph_step3.png" alt="digraph example"/> </a>
* __(4)__ Nodes 3 and 5 are done, they informs the conductor
<a href="/assets/orchestrate-a-digraph-with-goroutine/digraph_step4.dot"><img class="img-responsive" src="/assets/images/digraph_step4.png" alt="digraph example"/> </a>
* __(5)__ conductor update the matrix
<a href="/assets/orchestrate-a-digraph-with-goroutine/digraph_step5.dot"><img class="img-responsive" src="/assets/images/digraph_step5.png" alt="digraph example"/> </a>
* __(6)__ The conductor feeds the nodes with the matrix
<a href="/assets/orchestrate-a-digraph-with-goroutine/digraph_step6.dot"><img class="img-responsive" src="/assets/images/digraph_step6.png" alt="digraph example"/> </a>
* __(7)__ The nodes analyse the matrix
<a href="/assets/orchestrate-a-digraph-with-goroutine/digraph_step7.dot"><img class="img-responsive" src="/assets/images/digraph_step7.png" alt="digraph example"/> </a>
* __(8)__ Node 2 can run...
<a href="/assets/orchestrate-a-digraph-with-goroutine/digraph_step8.dot"><img class="img-responsive" src="/assets/images/digraph_step8.png" alt="digraph example"/> </a>

## The representation of the use case in go

### Data representation
to keep it simple, I won't use a `list` or a `slice` to represent the matrix, but instead I will rely on the [package mat64](https://godoc.org/github.com/gonum/matrix/mat64).

A slice may be more efficient, but by now it is not an issue. 

On top of that, I may need later to transpose or look for eigenvalues, and this package does implement the correct method to do so.

```golang
// Allocate a zeroed array of size 8×8
m := mat64.NewDense(8, 8, nil)
m.Set(0, 1, 1); m.Set(0, 4, 1) // First row
m.Set(1, 6, 1); m.Set(1, 6, 1) // second row
m.Set(3, 2, 1); m.Set(3, 6, 1) // fourth row
m.Set(5, 0, 1); m.Set(5, 1, 1); m.Set(5, 2, 1) // fifth row
m.Set(7, 6, 1) // seventh row
fa := mat64.Formatted(m, mat64.Prefix("    "))
fmt.Printf("\nm = %v\n\n", fa)
```

### The code

# References

* [Go Concurrency Patterns (Rob Pike)](https://talks.golang.org/2012/concurrency.slide)
