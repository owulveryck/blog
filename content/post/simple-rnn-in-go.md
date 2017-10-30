---
images: ["https://upload.wikimedia.org/wikipedia/en/6/63/Queen_A_Kind_Of_Magic.png"]
description: "You may know how enthusiast I am about machine learning. A while ago I discovered recurrent neural networks. I have read that this 'tool' allow to predict the future! Is this a kind of magic? I have read a lot of stuffs about the 'unreasonable effectiveness' of this mechanism. The litteracy that gives deep explanation exists and is excellent. There is also plehtora of examples, but most of them are using python and a calcul framework. To fully undestand how things work (as I am not a data-scientist), I needed to write my own tool 'from scratch'. This is what this post is about: a more-or-less 'from scratch' implementation of a RNN in go that can be used to applied to a lot of examples"
categories: ["category"]
tags: ["tag1", "tag2"]
draft: true
title: "YAPARNN! - Yet another post about RNN. But this one is made of go"
date: 2017-10-29T07:17:33+01:00
type: post
---

# Why would I want to play with RNN?

## Shakespeare and I, encounter of the third type

A couple of months ago, I have attented to the Google Cloud Next 17 event in London.
Among the talks about SRE, and keynotes, I've had the chance to listen to Martin Gorner's excellent introduction: [TensorFlow and Deep Learning without a PhD, Part 2](https://www.youtube.com/watch?v=fTUwdXUFfI8). If you don't want to look at the video, here is a quick summary:  

_a 100 of lines of python are reading all of Shakespeare's playsr;, it learns his stysle, and then generates a brand new play from scratch._ 

Of course, when you are not data-scientist (and I am not), this looks pretty amazing (and a bit magical).

Back home, I have told my friend how amazing it was. I have downloaded the code from [github](https://github.com/martin-gorner/tensorflow-rnn-shakespeare), installed tensorflow, and played my Shakespeare to show them.
They told me: _and you know how this works? Well...._

It was about something called "Recurrent Neural Networks". 
Then I dived into the internet... 100 lines of python shouldn't be hard to understand and reproduce... It took me months to be able to write this post.
I hope I will be as helpful to you as it is to me.

## The RNN and I, first episode of a time-serie 

From an explanation to another, I've learned that RNN could, by nature, predict time series.

- _Wait, does it mean that it can predict the future?_,
- Well, kind of... 

We are still in the area of what is called "supervised learning". Therefore, based on what it has learned, the algorithm can predict what will come next, only if it is something that has already been seen. 
Let me take an example and consider the lottery (everybody ask me about this):

To win, you need to own a ticket with a sequence of numbers that corresponds to the one that will be choosen randomly at the next lottery draw.
If you can predict the future, you could predict which sequence will be choosen.

The RNN is supervised learning, therefore, it can only predict things based on stuffs it has already seen. So If every week the draw is made of "1 2 3 4 5 6", the RNN can lear, and tell us that the next draw will be: "1 2 3 4 5 6".
Obiously this is useless; now let's consider a more complex sequence: 

Week | sequence
-----|---------
1    | 1 2 3 4 5 6
2    | 2 3 4 5 6 1
3    | 3 4 5 6 1 2
4    | 4 5 6 1 2 3
5    | 5 6 1 2 3 4
6    | 6 1 2 3 4 5
7    | 1 2 3 4 5 6
     | ...

What will be the winning sequence of week 8? 

"2 3 4 5 6 1". Cool, you are rich! 
How did you do? You have memorized the sequence. RNN does exactly the same.

- So, I can predict the next lottery? 
- No, because there is no sequence in the lottery, it is pure randomness.

In other words there is no "recurrence" in the drawing, therefore "recurrent" neural networks cannot be applied. 
 
Anyway, beside the lottery, a lot of events are, in essence, recurrent.
But, the recurrency model is not always easy to detect.

This is where a RNN could help us a lot in our professional lifes.

For example, on certain systems, you can have failures "every now-and-then". Even if you don't find the root cause, it could be useful to predict the next failure. 
If you have enough data about the past failures, the RNN could learn the pattern, and tell you when the next failure will occur.

Ok, this is it, I need to learn how to do that. This article was really inspiring: [how to do time series prediction using RNNs, Tensorflow and Cloud ML engind)](http://dataconomy.com/2017/05/how-to-do-time-series-prediction-using-rnns-tensorflow-and-cloud-ml-engine/). But as a huge majority of articles, it is related to python and a framework (here tensorflow). It has two major drawbacks:

* We need to fully understand how use the framework;
* As it is python related, and I am not fluent in python, building and deploying efficient tools could take me some time;

I have then decided to fully implement a RNN from scratch in GO with a simple goal: to understand what I was writing.

_Whatever is well conceived is clearly said, And the words to say it flow with ease._
(_Ce que l'on conçoit bien s'énonce clairement, et les mots pour le dire arrivent aisément._)

__Nicolas Boileau__

# The initial example

All the following example is basically an adaptation of Andrej Karpathy's post: [The Unreasonable Effectiveness of Recurrent Neural Networks](http://karpathy.github.io/2015/05/21/rnn-effectiveness/).

I strongly encourage you to read the post. Indeed, I will give you a couple of explanation of the principle.
The goal is to write and train a RNN with a certain amount of text data.

Then, once the RNN is trained, we ask it to generate a new text based on what it has learned.

## How does it work?

Consider the "HELLO" example as described in Karpathy's post.
The vocabulary of the example is composed of 4 letters: h, e, l and o.

The goal is to train the RNN network in order to make it predict the next letter.

Therefore, if I give an _H_ as input the fully trained RNN, it will return an _E_,

Then, the _E_ will become the input, and the output will be an _L_.

This _L_ will become the new input. After an _L_, I can have another _L_ or an _O_; but the RNN has a memory,
in essence, it remember the last letters. Then, it will most probably choose a second _L_, based on the _H_ and _E_.

Now we have another _L_ as input, but the context has changed, and the RNN should be able to produce an _O_.

## A classification problem

This is a classification problem. 
