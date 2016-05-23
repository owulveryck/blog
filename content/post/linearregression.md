---
author: Olivier Wulveryck
date: 2016-05-20T12:50:59+02:00
description: A post about machine learning and an application on a simple case I've met in my job. 
  Here is the use case \:regarding three different technical solutions, 
  after the evaluation by group of human of different features, can a Bot "think" on its own and evaluate which one offers then best ratio cost/features.
  And therefore, can it tell any manager which solution to choose.
draft: true
tags:
- machine learning
- octave
- linear regression
title: Which solution should I choose? Don't think too much and ask a bot!
topics:
- topic 1
type: post
---

# Let me tell you a story: the why!

A year ago, one of those Sunday morning where spring starts to warm up the souls, I went, as usual to my favorite bakery.
The family tradition is to come back with a bunch of "Pains au Chocolat" (which, are, you can trust me, simply excellent).

- _hello sir, I'd like 4 of your excellent "pains au chocolat" please_
- _I'm sorry, I don't have any pains au chocolat nor any croissant anymore_
- _what? How is it possible ?_
- _everything has been sold._
- _too bad..._

I think to myself: _why didn't you made more?_. He may have read my thought and told me:

- _I wish I could have foreseen_
 
When I left his shop, his words were echoing... I wish I could have foreseen... We have self-driving cars, we have the Internet, 
we are a technology advanced civilization. 
Facebook recognize your face among billions as soon as you post a photo... It must be possible to foresee...

This is how I started to gain interest in machine learning

At first I started to read some papers, then I learn (a very little bit) about graph theory, Bayesian networks, Markov chains.
But I was not accurate and I felt I was missing some basic theory.

That's the main reason why, 8 weeks ago, I signed in a course about ["Machine learning" on Coursera](https://www.coursera.org/learn/machine-learning). 
This course is given by [Andrew Ng](http://www.andrewng.org/) from [Stanford University](https://www.stanford.edu/)

# So what?

The course is not finished yet, but after about 8 weeks, I've learn a lot about what we call "machine learning".

The main idea of the machine learning is:

* to feed some code with a bunch of data (who said big data was useless)
* to code or encode some mathematical formulas that could represent the data
* to implement any algorithm that optimize the formulas by minimizing the error made by the machine on the evolving data sets.

To make it simple: machine learning is feeding a "robot" with data and teach him how to analyse the errors so it can make decisions on its own.

Scary isn't it?

In this post I will explain

