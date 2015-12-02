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

If we consider this very simple graph:

* A -> B
* A -> C
* B -> C

The workflow is simply: 

* A can run
* B can run if A is done
* C can run if A and B are done

# The adjacency matrix



# References

* [Go Concurrency Patterns (Rob Pike)](https://talks.golang.org/2012/concurrency.slide)

