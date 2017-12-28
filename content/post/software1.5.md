+++
images = ["/assets/lstm/LSTM-cell.png"]
description = ""
draft = true
title = "Parsing mathematical equation to generate computation graphs - First step from software 1.0 to 2.0 in go"
date = 2017-12-18T16:47:27+01:00
author = "Olivier Wulveryck"
+++

In my previous article, I have explained how to code a RNN from scratch in go.
The goal of this work is to use the RNN as a processing unit for different information that I could grab in my day-to-day work.
(for example to find a the root-cause of an incident, or as a helper decision tool for capacity management).

The purpose of this article is to describe a way to code in software 1.0 an execution machine for a software 2.0.

I will first explain the concepts, then I will describe how to implement different parts.

# Considerations about software 1.0 and software 2.0

## What is a software?

It is sequence of bits and bytes that can be computed, and that produces a result (the solution of a problem for example). 

To build a software until now, a compiler is used. Its goal is to turn a "human readable sequence of characters", called code, into the sequence of bytes. 

This sequence of bytes is evaluated and executed by a machine at runtime. Depending of the input, the execution produces (hopefully) the expected output.

The art of programming, is, in essence, the faculty for a human to describe the solution of a problem and to express it in a computer language. 

## What is software 2.0?

I discovered the concept of software 2.0 thanks to [Andrej Karpathy's blog](https://medium.com/@karpathy/software-2-0-a64152b37c35).
The idea is similar to any software: a compiler is used to turn a sequence of code into a sequence of bytes. This sequence is interpreted by a machine.

The difference is that the code is a sequence of mathematical equations (called model). Those equations are composed of variables and "constants". Let's call the constants "the weights".

The compiler is a software 1.0 that is able to transpile the equations into a sequence of bytes that will be evaluated by a machine (note that the compiler itself is a machine).

So what is the difference between 1.0 and 2.0? Is it just a matter of language?

No, the major difference is in the art of programming and the use case. 

For example:

A programmer __cannot__ write an algorithm that will __solve a specific problem__ (eg: I need to recognize a cat on any photo).

So, the programmer will write a set of equations __able to solve a kind of problem__ (recognize objects on any photo). 

The solution to the specific problem will be given by the evaluation of the _equation_ __with__ _its weights_ (a cat is an object that corresponds to the specific weights: {0,1,3,2,45,6,6,5,3,4,6,....}.)

And what makes the software 2.0 so specific? The amount of weights is so important that it cannot be determined manually. They are determined empirically. And a computer is faster than any human in this learning process.

# Example of a software 2.0: Deep learning

Neuron networks are the perfect representation of the software 2.0.
In my last [blog post](/2017/10/29/about-recurrent-neural-network-shakespeare-and-go.html) I have implemented a recurrent neural network in pure go.

My toy is working, I do not have the expected results: the generated text is poor and repetitive (for example it generates: `hello, the the the the the the...`). Vanillas RNNs are suffering from the [vanishing gradient problem](https://en.wikipedia.org/wiki/Vanishing_gradient_problem) which is most likely the root cause of my problems.

One solution is to change the core model for a more robust network called __L__ong __S__hort __T__erm __M__emory network (LSTM for short).

The software 2.0 will be an implementation of the equations described 

_Sidenote about go_: I am a gopher and an Ops. I really like go because I find it easy and fun to do fancy stuffs. But actually go is not the first choice when we talk about machine learning. My goal is to write a kind of portable virtual machine for software 2.0. I will not explain this in details in this post why go, and anything about software 2.0; But the facilities offered by the go language in order to reach my goal.

## LSTM

LSTM are a bit more complex than vanilla RNN. Therefore, a naive go implementation as made for the RNN will be a harder.

As one of my goal is to understand how things deeply works, I have tried to implement the back propagation mechanism manually without any luck.
I have read this post from Karpathy: [Yes you should understand backprop](https://medium.com/@karpathy/yes-you-should-understand-backprop-e2f06eab496b).

The best explanation I have found so far is in [cs231n course from Stanford](http://cs231n.github.io/optimization-2/).
It is a clear explanation of how the process works. And it is obvious that the graph representation helps a lot in the computation of the gradient.

## Equations are graphs

So equations are graphs... Cool, I have always been attracted by graphs. It is a very natural way to understand and express the ideas. This [post](http://gopherdata.io/post/deeplearning_in_go_part_1/) from [Chewxy](https://twitter.com/chewxy) is a perfect illustration of how the expression of a mathematical expression is turned into a graph at a compiler level.

It sounds that implementing the LSTM as a graph will make the task a lot easier. 

# Writing the machinery: software 1.0

Machine learning is about graphs and tensors. It exists some optimized library to transpile the equations into a graph. Tensorflow is one of those.
Tensorflow is highly optimized, but the setup of the working environment may be tricky from times to time.

On top of that, the binding exists for the go language, but their purpose is to run a software 2.0 and not to code the model.
Tensorflow does some things that are too magic for me by now, and it is too much abstract. I want something simpler.

## Gorgonia

Chewxy, the author of the post about equation, is alors the author of the Gorgonia project. 

> Package gorgonia is a library that helps facilitate machine learning in Go. Write and evaluate mathematical equations involving multidimensional arrays easily. Do differentiation with them just as easily.

I have talked to Chewxy on the channel #data-science on #slack. He is really commited, and very active. On top of that I am really attracted by the idea of such a library in go. 
I have decided to give gorgonia a try. 

### Machines, Graphs, Nodes, Values and Backends

In gorgonia an equation is represented by an [`ExprGraph`](https://godoc.org/github.com/gorgonia/gorgonia#ExprGraph). It is the main entry point of Gorgonia.
A graph is composed of [`Nodes`](https://godoc.org/github.com/gorgonia/gorgonia#Node).
A node is any element in the graph. It is a placeholder that will host a [`Value`](https://godoc.org/github.com/gorgonia/gorgonia#Value).

A `Value` is an interface. A [`Tensor`](https://godoc.org/gorgonia.org/tensor#Tensor) is a type of `Value`.

`Tensors` contains elements of the same [`Dtype`](https://godoc.org/gorgonia.org/tensor#Dtype). All those elements are stored in concrete arrays of elements (for example `[]float32`).

To actually compute the graph, Gorgonia is using on of the two implementation of Machines: 

* [`lispMachine`](https://godoc.org/gorgonia.org/gorgonia#NewLispMachine)
* [`tapeMachine`](https://godoc.org/gorgonia.org/gorgonia#NewTapeMachine)

#### Building a graph

To transform a mathematical equation into a graph, we first need to create a graph, then create the Values, assign them to some nodes and add the nodes to the graph.

For example, this equation:

$$z = W \cdot x$$
With 
$$W = \begin{bmatrix}0.95 & 0,8 \\\ 0 & 0\end{bmatrix}, x = \begin{bmatrix}1 \\\ 1\end{bmatrix}$$ 

Is written like this in "gorgonia":

```go
// Create a graph
g := G.NewGraph()

// Create the backend with the inputs
vecB := []float32{1,1}
// Create the tensor and specify its shape
vecT := tensor.New(tensor.WithBacking(vecB), tensor.WithShape(2))
// Create a node of type "vector"
vec := G.NewVector(g,
        tensor.Float32,    // The type of the data encapsulated within the node
        G.WithName("x"),   // The name of the node (optional)
        G.WithShape(2),    // The shape of the Vector
        G.WithValue(vecT), // The value of the node
)
matB := []float32{0.95,0.8,0,0}
matT := tensor.New(tensor.WithBacking(matB), tensor.WithShape(2, 2))
mat := G.NewMatrix(g, 
        tensor.Float32, 
        G.WithName("W"), 
        G.WithShape(2, 2), 
        G.WithValue(matT),
)

// z is a new node of the graph "g".
// It does not contains the actual result because the graph
// has not be computed yet
z, err := G.Mul(mat, vec)
// ... error handling

// create a VM to run the program on
machine := G.NewTapeMachine(g)

// The graph is executed now !
err = machine.RunAll()
// ... error handling
// Now we can print the value of z
fmt.Println(z.Value().Data())
// will display [1.75 0] which is a []float32{}
```



## Good ol' software 1.0

### Lexer/Parser

<center>  
{{< tweet 941817771863584768 >}}
</center>

[goyacc example](https://github.com/golang/tools/tree/master/cmd/goyacc/testdata/expr)

### goyacc

{{< highlight go >}}
// Forward pass as described here https://en.wikipedia.org/wiki/Long_short-term_memory#LSTM_with_a_forget_gate
func (l *lstm) fwd(inputVector, prevHidden, prevCell *G.Node) (hidden, cell *G.Node) {
        // Helper function for clarity
        set := func(ident, equation string) *G.Node {
                res, _ := l.parser.Parse(equation)
                l.parser.Set(ident, res)
                return res 
        } 

        l.parser.Set(`xₜ`, inputVector)
        l.parser.Set(`hₜ₋₁`, prevHidden)
        l.parser.Set(`cₜ₋₁`, prevCell)
        set(`iₜ`, `σ(Wᵢ·xₜ+Uᵢ·hₜ₋₁+Bᵢ)`)
        set(`fₜ`, `σ(Wf·xₜ+Uf·hₜ₋₁+Bf)`) // dot product made with ctrl+k . M
        set(`oₜ`, `σ(Wₒ·xₜ+Uₒ·hₜ₋₁+Bₒ)`)
        // ċₜis a vector of new candidates value
        set(`ĉₜ`, `tanh(Wc·xₜ+Uc·hₜ₋₁+Bc)`) // c made with ctrl+k c >
        ct := set(`cₜ`, `fₜ*cₜ₋₁+iₜ*ĉₜ`)
        set(`hc`, `tanh(cₜ)`)
        ht, _ := l.parser.Parse(`oₜ*hc`)
        return ht, ct
}
{{</ highlight >}}

If you don't have the correct font to display the unicode character, you may find a picture [here](/assets/lstm/uni-code.png)

![image](/assets/lstm/LSTM.png)

# Conclusion

