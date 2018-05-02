---
title: "Considerations about software 2.0"
date: 2018-04-16T10:54:23+02:00
lastmod: 2018-04-16T10:54:23+02:00
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
mathjax: false
---

In my last post I've described the concept of software 2.0.
I have evaluated the idea of implementing a parser of equation (written in unicode) to give a strict separation of the software 1.0 and the software 2.0.

In this post I will go further in the description of this concept with the help of the famous "char RNN" example.

# The use case

The software 1.0 is, for me, a kind of virtual machine. This machine's goal is to execute the software 2.0.
The concept is similar to what we find in the Java world. A runtime environment is able to execute some bytecode. Our bytecode is composed of sequences of float (the weight matrix) and a bunch of equations.
There is a separation of the SDK (which will have a training method), and a NNRE (Neural Network Runtime Environment) which can load and apply the software 2.0.

# Sample implementation

The use case I have chosen is an implementation of a char based LSTM (it's a well known example that has been promoted by Anderij Karpathy).

The code has been implemented as a go package that can be used to build a SDK and/or a NNRE.
The LSTM "kernel" has been copied from the wikipedia page

![img](https://wikimedia.org/api/rest_v1/media/math/render/svg/8a0eddfb6f592041ea04bd26526b52ba1cec192c)

Its implementation is not fully externalized yet and is hardcoded within the LSTM:

![content](/assets/lstm/lstm_implem.png)

See the [code here](https://github.com/owulveryck/lstm/blob/1581884e9d2de83e1150c04fb815637351082b7a/lstm.go#L39-L46)




https://www.ted.com/talks/lera_boroditsky_how_language_shapes_the_way_we_think




