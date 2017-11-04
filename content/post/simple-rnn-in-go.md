---
images: ["https://upload.wikimedia.org/wikipedia/en/6/63/Queen_A_Kind_Of_Magic.png"]
description: "You may know how enthusiast I am about machine learning. A while ago I discovered recurrent neural networks. I have read that this 'tool' allow to predict the future! Is this a kind of magic? I have read a lot of stuffs about the 'unreasonable effectiveness' of this mechanism. The litteracy that gives deep explanation exists and is excellent. There is also plehtora of examples, but most of them are using python and a calcul framework. To fully undestand how things work (as I am not a data-scientist), I needed to write my own tool 'from scratch'. This is what this post is about: a more-or-less 'from scratch' implementation of a RNN in go that can be used to applied to a lot of examples"
draft: true
title: "About Recurrent Neural Network, Shakespeare and GO"
date: 2017-10-29T07:17:33+01:00
type: post
---

# Shakespeare and I, encounter of the third type

A couple of months ago, I have attended the Google Cloud Next 17 event in London.
Among the talks about SRE, and keynotes, I've had the chance to listen to Martin Gorner's excellent introduction: [TensorFlow and Deep Learning without a PhD, Part 2](https://www.youtube.com/watch?v=fTUwdXUFfI8). If you don't want to look at the video, here is a quick summary:

_a 100 of lines of python are reading all Shakespeare's plays; it learns his style, and then generates a brand new play from scratch._ 

Of course, when you are not data-scientist (and I am not), this looks pretty amazing (and a bit magical).

Back home, I have told my friend how amazing it was. I have downloaded the code from [github](https://github.com/martin-gorner/tensorflow-rnn-shakespeare), installed tensorflow, and played my Shakespeare to show them.
In essence, here is what they told me:

- _Amazing, and you know how this works?_ 
- _Well..._ let's be honest, I had only a vague idea.

It was about something called "Recurrent Neural Networks". 
I dived into the internet... 100 lines of python shouldn't be hard to understand, and to reproduce ?
Was-it? Actually, it took me months to be able to write this post, without any previous knowledge, it was not that easy.

So here is why I finally wrote this article. I want to be sure that I have understood the structure and the possibilities offered by recurrent neural network.
I aslo wanted to see whether building a RNN powered tool was doable easily.

This document is divided into two:

* the first part is about recurrent neural networks in general;
* the second part is about a toy I made in GO to play with RNNs

The goal of this text is not to talk about the mathematics behind the neural nerworks.
Of course, I may talk about vectors, but I will not talk about non-linearity or hyperbolic functions. 

I hope you will as much enthusiast as I am. Anyway Do not hesitate to give me any feedback or correction that may improve my work.

# First part: The RNN and I, first episode of a time-serie 

The Web is full of resources about machine learning. You can easily find great articles, very well illustrated about neural networks.
I've read a lot...

The more I was reading, the more excited I was. 

For example, from an explanation to another, I've learned that RNN could, by nature, predict time series.
(cf [how to do time series prediction using RNNs, Tensorflow and Cloud ML engine](http://dataconomy.com/2017/05/how-to-do-time-series-prediction-using-rnns-tensorflow-and-cloud-ml-engine/)).

- _Wait, does it mean that it can predict the future?_,
- Well, kind of...

It is still in the area of "supervised learning". Therefore, the algorithm learns eventsi. Based on this, the algorithm can predict what may come next; but only if it is something it has already seen. 
Let me take an example. Consider a lottery game (everybody ask me about this):

To win, you need to own a ticket with a sequence of numbers that corresponds to the one that will be chosen randomly at the next lottery draw.
If RNN you can predict the future, it should, basically, be able to predict it.

The RNN must learn about the sequences and then it applies its knowledge. So If every week the draw is made of "1 2 3 4 5 6", the RNN will learn, and tell that the next draw will be: "1 2 3 4 5 6".

Obviously this is useless; now let's consider a more complex sequence:

Week | sequence
-----|---------
1    | 1 2 3 4 5 6
2    | 2 3 4 5 6 1
3    | 3 4 5 6 1 2
4    | 4 5 6 1 2 3
5    | 5 6 1 2 3 4
6    | 6 1 2 3 4 5
7    | 1 2 3 4 5 6
8    | ? ? ? ? ? ?

Question: What will be the winning sequence of week 8? 

"2 3 4 5 6 1". Cool, you are rich! 
How did you do? You have memorized the sequence. RNN does exactly the same.

- So, it **can** predict the next lottery? 
- No, because there is no sequence in the lottery, it is pure randomness.

In other words there is no "_recurrence_" in the drawing, therefore "_recurrent_" neural networks do not be applied. 
 
Anyway, beside the lottery, a lot of events are, in essence, recurrent.
The point is that the recurrency model is usually not obvious and therefore not easy to detect. This is the famous "feeling" of the experts. 

For example you may recognize those dialogs:

- Will the system crash?
- based of what I see and what I know, I [don't] think so.

- Will the sales increase on sunday?
- based on the current market situation and on my experience, it may.

This is where a RNN could shine and enhance our professional lives.

In a pure IT context, for example, you have failures "every now-and-then". Even if you don't find the root cause, it could be useful to predict the next failure. 
If you have enough data about the past failures, the RNN could learn the pattern, and tell you when the next failure will occur.

<center>
{{< tweet 844561153229541376 >}}
</center>

### Experimenting with RNN

I needed a simple tool to do some experimentation.
A huge majority of articles that deals with ML are using python and a framework (here tensorflow).
To me, it has two major drawbacks:

* I need to fully understand how use the framework;
* as it is python related, and I am not fluent in python, building **and deploying** efficient tools could take me some time;

About the second point, let me be a bit more specific. I have seen a lot of samples that could do very beautiful stuffs based on fake data.
Playing with every day data usually implies to rewrite the tool, from scratch... Therefore, I have have decided to fully implement a RNN engine from scratch in GO. The goal is simple: to understand what I am writing (I am "fluent" in go, that have save me days of debugging).

_Whatever is well conceived is clearly said, And the words to say it flow with ease._
(_Ce que l'on conçoit bien s'énonce clairement, et les mots pour le dire arrivent aisément._)

__Nicolas Boileau__

## The initial example

All the following example is basically an adaptation of Andrej Karpathy's post: [The Unreasonable Effectiveness of Recurrent Neural Networks](http://karpathy.github.io/2015/05/21/rnn-effectiveness/).

I strongly encourage you to read the post. Indeed, I will give you a couple of explanation of the principle.
The goal is to write and train a RNN with a certain amount of text data.

Then, once the RNN is trained, we ask the tool to generate a new text based on what it has learned.

### How does it work?

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

### A classification problem

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

# Second part: Let's geek

From now on, things are related to the implementation; feel free to skip this part and jump straight to the conclusion if your are not interested in coding.

My goal is to generate a code that will be able to generate a Shakespeare play as described in Karpathy's blog post.
His implementation in python is [here](https://gist.github.com/karpathy/d4dee566867f8291f086); you can find my implementation [here](https://github.com/owulveryck/min-char-rnn).

**edit**: at first, it was a simple transcript from python to go, but the tool has been enhanced and is now a more generic tool that is able to use a RNN as a processing units for any function that is able to encode and decode bytes into a vector.

# The rnn package

To fully understand what is related to the RNN, and what is more related to the example about character recognition, I have created a separate package for the RNN.
For the same reason, I have tried to keep parameters as private as possible within the objects.

The package should be independent of the example (min-char); therefore it should probably be suitable to another classification problem.

I am using the [`mat64.Dense`](https://godoc.org/github.com/gonum/matrix/mat64) structure to represent the matrices. To represent a column vector, I have chosen to use simple `[]float64` elements (for more info: [Go Slices: usage and internals](https://blog.golang.org/go-slices-usage-and-internals#clices)). 

### The RNN object

The RNN structure holds the three matrices that represent the weights to be adapted:

* Wxh
* Whh
* Why

On top of that, the RNN store two "vectors" that represent the [bias](https://stackoverflow.com/questions/2480650/role-of-bias-in-neural-networks) for the hidden layer and for the output layer.
The hidden vector is not stored within the structure. Actually, only the last hidden vector evaluated in the process of feedforward/backpropagation is stored.
Not storing the hidden vector within the structure, allows to use the same "step" function in the sampling process as well as the training process.

### RNN's step

RNN's step method is the proper implementation of the RNN as described by _Karpathy_.
As explained before, the hidden state is not part of the RNN structure, therefore it is an output of the step function:

```go
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
```

You see here that the step function of my RNN takes two vector as input: 

* a vector that represents the currently evaluated item (remember, it is the representation of the `H`, `E`, `L` and `O` in the previous example)
* a hidden vector that stores the memory of the passed elements

It returns two vector:

* the evaluated output in term of vector (again it is the representation of the  `H`, `E`, `L` and `O`) 
* a new and updated hidden vector 

_Note_ : For clarity, I have declared a couple of math helpers such as `dot`, `tanh` and `add` that are out of the scope of the explanation.

### The Train method

This method is returning two channels and triggers a goroutine that does the job of training.

```go
func (rnn *RNN) Train() (chan<- TrainingSet, chan float64) {
    ...
}
```

The first channel is a feeding channel for the RNN. It receives a `TrainingSet` that is composed of:

* an input vector
* a target vector 

The goroutine will read the channel, and get all the training data.
It will evaluates the input of the training set and use the target to adapt the parameters

The second channel is a non blocking channel. It is used to transfer the loss evaluated at each pass for information purpose.

### Forward processing

The forward processing takes a batch of inputs (an array of array) and a sequence of outputs.
It runs the step as many times as needed and stores the hidden vectors in a temporary array, then the values are used for the back propagation.

```go
func (rnn *RNN) forwardPass(xs [][]float64, hprev []float64) (ys, hs [][]float64) {
    ...
}
```

### Back propagation through time

The back propagation is evaluating the gradient. Once the evaluation is done, we can adapt the parameters thanks to the computed gradients. 

### Adapting the parameters via "AdaGrad"

The method that has been used by Karapathy is the Adaptive gradient.
The adaptive gradient needs a memory; therefore I have declared a new object for the adagrad with a simple Apply method.
The `apply` method  of the `adagrad` object takes the neural network as a parameter as well as the previously evaluated gradients.
e
Once this process is done, the RNN is trained and is usable. 

### Prediction 

I have implemented a `Predict` method that applies the same method. But the difference is that it starts with an empty memory (the hidden vector is zeroed), takes a sample text as input and generate the output without evaluating the gradient nor adapting the parameters.

This RNN implementation is enough to generate the Shakespeare, and it works

## Enhancement of the tool: implementing codecs

In order to work with any character (= any symbol), the best way to go is to use the concept of [rune](https://blog.golang.org/strings).
The first implementation of the min-char-rnn I made, was using this package and a couple of method to 1-of-k encode and decode the rune I was reading from a text file.

This was ok, but I was stuck within the character based neural network.
As I explained before, the RNN package is working with vectors, and have no knowledge about characters, pictures, bytes or whatever.

So to continue with this level of abstraction, I have declared a codec interface. The character based example will be a simple implementation that will fulfill the codec interface.
It will allow to implement whatever codec to use my RNN (imagine a log parser, an image encoder/decoder, a webservice [_insert whatever fancy idea here_]...)

### The codec interface

The codec interface describes the required methods any object must implement in order to use the RNN.

The most important methods are:

```go
Decode([][]float64) io.Reader
Encode(io.Reader) [][]float64
```

Actually, those methods are dealing with arrays of vectors on one side, and with `io.Reader` on the other side.
Therefore, it can use any input type, from a text representation to a data flow over the network (and if you are _gopher_, you know how cool `io.Reader`s are!)

The other methods are simply helper funcs I use to train the network. (and I should rework that anyway, because _Pike_ loves the one-function-interfaces, and _Pike_ knows!)

I will simply explain a particular method: 

```go
ApplyDist([]float64) []float64
```

This method is a post processing of the output vector. Actually, the returned vector is made of normalized probabilities of event. In a classification mechanism, one element must be chosen. Obviously, it shall choose the one with the best probability. But, in the case of the char example, we can add some randomness by choosing randomly and applying a certain distribution (I have implemented a [Bernouilli distribution](https://godoc.org/github.com/gonum/stat/distuv#Categorical) for the char codec that is selectable by setting ` CHAR_CODEC_CHOICE=soft` in the environment). It also let the possibility to get the raw normalized probabilities by implementing a no-ops func.

### The char implementation of the codec interface

As explained before, the char implementation consists of a couple of methods that reads a file and encode it.
It can also decode a generated output.

It is that simple. It also serve as an example for whatever new codec implementation

## The main tool

The main tool is just the glue between all the packages. 
It can be used to train the network or to generate an output. The parameters are tweakable via environment variables (actually each package deals with its own env variables).

# Example

Here is an example of the generated shakespeare. The network is not fully trained, and I did not try to optimize the input, therefore, you don't see a proper play (yet).

# Conclusion

This is only the beginning of what I am planing with RNN. I have now understood the principle. The tool offers me very good opportunities to develop some sample that could be useful in my everyday life.
For example, I am planing to write a codec that would parse log files of an application. This application would generate an output that could be decoded into a status of the platform, such as red, green and blue.
With the correct data about when warnings or failures occurred, it could be doable to predict the next failure... before it happens.

The code is on my github. It needs tweaking and deep testing. I have learned as well that testing a neural network was not as easy as testing a web app.
I have implemented a backup and restore mechanism, therefore you can retrain a model, or use a pre-trained model.

I have also uploaded a binary distribution and a pre-trained model if you want to play with it on your PC on [github](https://github.com/owulveryck/min-char-rnn).
