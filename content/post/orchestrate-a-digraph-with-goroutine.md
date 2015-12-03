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
They have changed the world a lot. From the simple representation to bayesian network via Markov chains, the applications are numerous.

Today I would like to imagine a graph as a workflow of execution. Every node would be considered as runable. And every  edge would be a dependency.

# The use case 
If we consider this very simple graph (example taken from the french wikipedia page)

<img class="img-responsive" src="https://upload.wikimedia.org/wikipedia/commons/0/07/Grafodirigido.jpg" alt="digraph example"/>

its corresponding adjacency matrix is:

 <img class="img-responsive" src="/assets/images/matrix1.png" alt="Adjacency matrix"/>

its dimension is 8x8

# Let's GO

## The representation of the use case in go

to keep it simple, I won't use a `list` or a `slice` to represent the matrix, but instead i will rely on the [package mat64](https://godoc.org/github.com/gonum/matrix/mat64).

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


# References

* [Go Concurrency Patterns (Rob Pike)](https://talks.golang.org/2012/concurrency.slide)