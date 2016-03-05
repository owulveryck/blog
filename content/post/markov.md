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
_Mark Burgess_ explains in his book [in search of certainty](http://http://www.amazon.com/gp/product/1491923075/ref=pd_lpo_sbs_dp_ss_1?pf_rd_p=1944687522&pf_rd_s=lpo-top-stripe-1&pf_rd_t=201&pf_rd_i=1492389161&pf_rd_m=ATVPDKIKX0DER&pf_rd_r=1BRFTEAZ2RRQ8M77MZ0C) that none should consider the command and control anymore.

Actually we grew up with the idea that a computer will do whatever we told him to.
The truth is that it simply don't. If that sounds astonishing to you, just consider the famous bug.
A bug is a little insect that will avoid any programmed behaviour to act as it should.

In a lot of wide spread software, we find _if-then-else_ or _try-catch_ statements.
Of course one could argue that the purpose of this conditional executionis is to deal with different scenarii, which is true, but indeed,
the keyword is _try_...

## Back to the choreography

In the choreography principle, the automation is performed by a set of dancer that acts on their own. Actually, the most logical way
to program it, is to let them know about the execution plan, and assume that everything will run as expected.

What I would like to study is simply that deployement without knowing the deployement plan.
The nodes would try to perform the task, and depending on the event they receive, they learn and enhance their probability of success.

### First problem


First, I'm considering a single node $A$  which can be in three states $\alpha$, $\beta$ and $\gamma$.
Let's $S$ be the set of states such as $S = \left\\{\alpha, \beta, \gamma\right\\}$

#### Actually knowing what's expected

The expected execution is: $ \alpha -> \beta -> \gamma$

therefore, the transition matrix should be:

$$
P=\\begin\{pmatrix\}
0 & 1 & 0 \\\\
0 & 0 & 1 \\\\
0 & 0 & 0
\\end\{pmatrix\}
$$

Let's represent it with GNU-R (see [this blog post](http://www.r-bloggers.com/getting-started-with-markov-chains/) 
for an introduction of markov reprentation with this software)

```R
> library(expm)
> library(markovchain)
> library(diagram)
> library(pracma)
> stateNames <- c("Alpha","Beta","Gamma")
> ExecutionPlan <- matrix(c(0,1,0,0,0,1,0,0,0),nrow=3, byrow=TRUE)
> row.names(ExecutionPlan) <- stateNames; colnames(ExecutionPlan) <- stateNames
> ExecutionPlan
      Alpha Beta Gamma
      Alpha     0    1     0
      Beta      0    0     1
      Gamma     0    0     0
> svg("ExecutionPlan.svg")
> plotmat(ExecutionPlan,pos = c(1,2), 
+         lwd = 1, box.lwd = 2, 
+         cex.txt = 0.8, 
+         box.size = 0.1, 
+         box.type = "circle", 
+         box.prop = 0.5,
+         box.col = "light yellow",
+         arr.length=.1,
+         arr.width=.1,
+         self.cex = .4,
+         self.shifty = -.01,
+         self.shiftx = .13,
+         main = "")
> dev.off()
```
which is represented by:

![Representation](/blog/assets/images/ExecutionPlan.svg)

#### Knowing part of the plan...


Now let's consider a different scenario. I assume now that the only known hypothesis is that we cannot go
from $\alpha$ to $\gamma$ and vice-versa, but for the rest, the execution may refer to this transition matrix:

$
P=\\begin\{pmatrix\}
\frac{1}{2} & \frac{1}{2} & 0 \\\\
\frac{1}{3} & \frac{1}{3} & \frac{1}{3}  \\\\
0 & \frac{1}{2} & \frac{1}{2} 
\\end\{pmatrix\}
$
which is represented this way ![Representation](/blog/assets/images/ExecutionPlan2.svg)


