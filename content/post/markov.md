---
author: Olivier Wulveryck
date: 2016-02-29T20:55:01+01:00
description: 
draft: true
keywords:
- key
- words
tags:
- one
- two
title: Is there a Markov model hidden in the choreography?
topics:
- topic 1
type: post
---

# Introduction

In my last post I introduced the notion of choreography as a way to deploy an manage application.
It could be possible to implement self-healing, elasticity and in a certain extent
self awareness.

To do so, we must not rely on the _certainty_ and the _determinism_ of the automated tasks.
_Mark Burgess_ explains in his book [in search of certainty](http://TODO) that none should consider the command and control anymore.

Actually we grew up with the idea that a computer will do whatever we told him to.
The truth is that it simply don't. If that sounds astonishing to you, just consider the famous bug.
A bug is a little insect that will avoid any programmed behaviour to act as it should.

In a lot of wide spread software, we find _if-then-else_ or _try-catch_ statements.
Of course one could argue that the purpose of this conditionals execution is to deal with different scenarii, which is true, but indeed...


## Back to the choreography

IN the choreography principle, the automation is performed by a set of dancer that acts on their own. Actually, the most logical way
to program it, is to let them know about the execution plan, and assume that everything will run as expected.

Let's consider a simple finite state machine with two nodes A and B, which can be in two states $\alpha$ and $\beta$.

The states of the couple (A,B) at time _t1, t2 and t3_ is

$$(A,B)_t1 = (\alpha,\alpha)$$

$$(A,B)_t2 = (\beta,\alpha)$$

$$(A,B)_t3 = (\beta,\beta)$$ 

Sounds easy to code as an execution plan.

$$A:\alpha -> A:\beta -> B:\alpha -> B:\beta$$

We notice that in the execution plan, there is no way to be in this state $(A,B)_t2 = (\alpha,\beta)$. Actually if for any reason state $\beta$ is buggy,
it may fail at any moment, and put A back in state $\alpha$.

What should then be the behaviour ot node B?

## From certainty to probability

In the previous example, we notice that spontaneously, our mindset is command and control based.
We didn't think at first of the probability of a failure; we may have in the code, if we were good coders, but in the theory that was not
suppose to happen.

But there is a _probability_ that it could happen. By now, we are totally unable to quantify this probability, but it's non-nil.

The only certainty we can get is the current states of the node. We cannot presume anything on how they came to ths state.

This is exactly a hidden markov process.

## HMM

According to _Rabiner_ (see [A tutorial on Hidden Markov Models and selected applications in speech recognition](http://TODO)):

> the first problem one faces is deciding what the states in the model correspond to, and the deciding how many states should be in the model.

## The obeservation sequence $O$

When $a \ne 0$, there are two solutions to $\(ax^2 + bx + c = 0\)$ and they are
$$x = {-b \pm \sqrt{b^2-4ac} \over 2a}.$$

$$\sum_{i=0}^n i^2 = \frac{(n^2+n)(2n+1)}{6}$$ _

and the distribution $\lambda$


$$\sum_{all Q} P(O|Q_i,\lambda) $$


$$\begin{bmatrix}a & b\\\c & d\end{bmatrix}$$




