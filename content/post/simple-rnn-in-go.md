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

_a 100 of lines of python are reading all Shakespeare's plays;, it learns his stysle, and then generates a brand new play from scratch._ 

Of course, when you are not data-scientist (and I am not), this looks pretty amazing (and a bit magical).

Back home, I have told my friend how amazing it was. I have downloaded the code from [github](https://github.com/martin-gorner/tensorflow-rnn-shakespeare), installed tensorflow, and played my Shakespeare to show them.
They told me: 

- _and you know how this works?_ 
- _Well..._ let's be honnest, I had only a vague idea.


It was about something called "Recurrent Neural Networks". 
I dived into the internet... 100 lines of python shouldn't be hard to understand and reproduce... It took me months to be able to write this post.
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
The vocabulary of the example is composed of 4 letters: `H`, `E`, `L` and `O`. 

The goal is to train the RNN network in order to make it predict the next letter.

Therefore, if I give an `H` as input the fully trained RNN, it will return an `E`,

Then, the `E` will become the input, and the output will be an `L`.

This `L` will become the new input. Here is a difficulty: after an `L`, there can be 

* another `L` or 
* an `O`; 

This is what make RNN suitable for this case: RNN has a memory!
Then, it will most probably choose a second `L`, based, not only on the last letter, but also on the previous `H` and `E` it has seen.

If correctly trained, the RNN should be able to produce an `O`.

## A classification problem

In practice, this is a [classification problem](https://en.wikipedia.org/wiki/Statistical_classification); Consider that every letter in the alphabet is a class.
Given a sequence of letter as input, the mechanism should predict to which class it belongs. This class is be the next letter to be displayed.

For example: 

- `h` belongs to class `e`
- `h e` belongs to class `l`
- `h e l` also belongs to classe `l`
- `h e l l` belongs to class `o`

The goal of the network is to give a probability for each class given the input and the context.
So, every letter will be given a value between 0 and 1 by the algorithm.

If we formalize that in an array, the ideal situation would be:

<html>
<table border=1 align=center>
<tr>
<th>context</th><th>input</th>
<th>Probability that the class is h</th>
<th>Probability that the class is e</th>
<th>Probability that the class is l</th>
<th>Probability that the class is o</th>
</tr>
<tr><td></td><td>h</td><td>0</td><td>1</td><td>0</td><td>0</td></tr>
<tr><td>h</td><td>e</td><td>0</td><td>0</td><td>1</td><td>0</td></tr>
<tr><td>h e</td><td>l</td><td>0</td><td>0</td><td>1</td><td>0</td></tr>
<tr><td>h e l</td><td>l</td><td>0</td><td>0</td><td>0</td><td>1</td></tr>
</table>
</html>

In pratcice, we may have something slightly different (this is an example, do not try to interpret the values):

<html>
<table border=1 align=center>
<tr>
<th>context</th><th>input</th>
<th>Probability that the class is h</th>
<th>Probability that the class is e</th>
<th>Probability that the class is l</th>
<th>Probability that the class is o</th>
</tr>
<tr><td></td><td>h</td><td>0.1</td><td>0.8</td><td>0.05</td><td>0.05</td></tr>
<tr><td>h</td><td>e</td><td>0.1</td><td>0.07</td><td>0.8</td><td>0.03</td></tr>
<tr><td>h e</td><td>l</td><td>0.05</td><td>0.05</td><td>0.5</td><td>0.4</td></tr>
<tr><td>h e l</td><td>l</td><td>0.05</td><td>0.05</td><td>0.4</td><td>0.5</td></tr>
</table>
</html>

We have encoded the output into an array; in mathematics, such array is called a vector.

On the same principle, we can encode the input letters into a _1-of-k_ vector (1 in the cell corresponding to character, 0 elsewhere).

<pre> 
<code>  
    h e l o
h = 1 0 0 0
e = 0 1 0 0
l = 0 0 1 0
o = 0 0 0 1
</code>
</pre>

The purpose of the prediction is to apply a mathematical function to an input vector in order to produce an output vector that will allow us to classify the output (and to to predict the next character).
The RNN does not know natively an equation able to predict the correct values. Instead, it knows a mathematical model (a mathematical function) that contains a lot of parameters or variables. With correct values, those parameters applied to the mathematical model should allow to reach the goal.
Adapting the parameters is called the _training process_. By giving the RNN a lot of data, and the expected output classes, we will allow the RNN to adjust its internal parameters.
At each step, the difference between the output and the expected result is evaluated; it is call "the loss". The purpose of the adaptation, it to reduce the loss at every step.

_It is not the purpose of this article to give details about the mathematical functions involved. Now, let's dig into the code._

# Let's geek

My goal is to generate a code that will be able to gerenate a Shakespeare play as described in Karpathy's blog post.
His implementation in python is [here](https://gist.github.com/karpathy/d4dee566867f8291f086); you can find my implemetation [here](https://github.com/owulveryck/min-char-rnn).

# Code organisation

To fully understand what is related to the RNN, and what is more related to the example about character recognition, I have created a seperate package for the RNN.
For the same reason, I have tried to keep parameters as private as possible within the objects.

## The `rnn` package

This package contains the model of the RNN. It is independant of the example (min-char); therefore it should probably be suitable to another classification problem.

I am using the `mat64.Dense` structure to represent the matrices. To represent a column vector, I am simply using `[]float64` elements. 

### The RNN object

The RNN structure holds the three matrices that represent the weights to be adapted:

* Wxh
* Whh
* Why

On top of that, the RNN store two "vectors" that represent the biais.
The hidden vector is not stored within the structure. Actually, only the last hidden vector evaluated in the process of feedforward/backpropagation is stored
within the structure.
Not storing the hidden vector within the structure, allows to use the same "step" function in the sampling process as well as the training process.

### RNN's step

RNN's step method is the proper implementation of the RNN as described by _Karpathy_.
As explained before, the hidden state is not part of the RNN structure, therefore it is an output of the step function:

{{< highlight go >}}
func (rnn *RNN) step(x, hprev []float64) (y, h []float64) {
	h = tanh(
		add(
			dot(rnn.wxh, x),
			dot(rnn.whh, hprev),
			rnn.bh,
		))
	y = add(
		dot(rnn.why, h),
		rnn.by)
	return
}
{{</ highlight >}}

You see here that the step function of my RNN takes two vector as input: 

* a vector that represents the currently evaluated item
* a hidden vector that stores the memory of the passed elements

It returns two vector:

* the evaluated output in term of vector
* a new and updated hidden vector

_Note_ : For clarity, I have declared a couple of math helpers such as `dot`, `tanh` and `add`

### The `Train`  method

This method is returning two channels and triggers a goroutine for training the network.

{{< highlight go >}}
func (rnn *RNN) Train() (chan<- TrainingSet, chan float64) {
    ...
}
{{</ highlight >}}

The first channel is a feeding channel for the RNN. It receives a `TrainingSet` that is composed of:

* an input vector
* a target vector 

The training goroutine will read the channel, and get all the TrainingSet.
It will evaluates the input of the training set and use the target to adapt the parameters

The second channel is a non blocking channel. It is used to transfer the loss evaluation.

### Forward processing

The forward processing takes a batch of inputs (and array of array) and a batch of outputs.
It runs the step as many times as needed and stores the hidden vectors in a temporary array, then the values are used for the back propagation.

{{< highlight go >}}
func (rnn *RNN) forwardPass(xs [][]float64, hprev []float64) (ys, hs [][]float64) {
    ...
}
{{</ highlight >}}

### Back propagation through time

The back propagation is evaluating the gradient. With this evaluation, we can adapt the parameters. 

### Adapting the parameters via "AdaGrad"

The method that has been used by Karapathy is the Adaptative gradient.
The adaptative gradient needs a memory; therefore I have declared a new object for the adagrad with a simple Apply method.
The `apply` method takes the neural network as a parameter as well as the previously evaluated derivated.
It then updated the RNN

## Codecs

This RNN implementation is enough to generate the Shakespeare. 
In order to work with any character (= any symbol), the best way to go is to use the concept of [rune](https://blog.golang.org/strings).
The first implementation of the min-char-rnn I made, was using this package and a couple of method to 1-of-k encode and decode the rune I was reading from a text file.

This was ok, but I was stuck within the character based neural network.
As I exaplained before, the RNN package is working with vectors, and have no knowledge about characters, pictures, bytes or whatever.

So to continue with this level of abstraction, I have declared a codec interface that will be fullfilled with a "character based" implementation.

### The codec interface

The codec interface describes the required methods any object must implement in order to use the RNN.

The most important methods are:

{{< highlight go >}}
Decode([][]float64) io.Reader
Encode(io.Reader) [][]float64
{{</ highlight >}}

Actually, those methods are dealing with arrays of vectors on one side, and with `io.Reader` on the other side.
Therefore, it can use any input type, from a text representation to a data flow over the network.

The other methods are simply helpful (and I should rework that anyway, because _Pike_ loves the one-function-interfaces, and _Pike_ knows!)

### The char implementation of the codec interface

As explained before, the char implementation consists of a couple of methods that reads a file and encode it.
It can also decode a generated output.

## The main tool

# Example

# Conclusion

