+++
images = ["/2016/10/image.jpg"]
description = ""
categories = ["category"]
tags = ["tag1", "tag2"]
draft = true
title: "Parsing mathematical equation to generate equations flow: Software 1.5 in go"
date: 2017-12-18T16:47:27+01:00
+++

In my last article, I have developped a recurrent neural network in pure go without any third party library.

As an example, I did an implementation of a character based generation (the famous Shakespeare example).
I have tried to tune the hyperparameters, but I haven't been able to reach a very usable text.

Actually, without any randomness in the generation process, the output was recurrent.

The point is that the toy I made is based on a vanilla RNN. And Vanillas RNNs are suffering from the [vanishing gradient problem](https://en.wikipedia.org/wiki/Vanishing_gradient_problem).
This is a well known problem, and one solution is to change the core model for a more robust network called __L__ong __S__hort __T__erm __M__emory network (LSTM for short).

# Implementing an LSTM


# About software 2.0

# Equations are graphs

## Gorgonia

# Good ol' software 1.0

## Lexer/Parser

### goyacc

# Conclusion

