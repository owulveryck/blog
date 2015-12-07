---
author: Olivier Wulveryck
date: 2015-12-07T08:48:20Z
description: A post about 
draft: true
tags:
- golang
- orchestrator
title: OaaS orchestrator as a service - part 1
topics:
- OaaS
type: post
---

In a [previous post](http://blog.owulveryck.info/2015/12/02/orchestrate-a-digraph-with-goroutine-a-concurrent-orchestrator/) I have setup and orchestrator that takes a digraph
as input (via its adjacency matrix).

In this post I will implement an API, so the orchestrator will be transformed into a web service.

