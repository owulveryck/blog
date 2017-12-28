+++
images = ["/assets/lstm/LSTM-cell.png"]
description = ""
categories = ["category"]
tags = ["tag1", "tag2"]
draft = true
title = "Parsing mathematical equation to generate computation graphs - First step from software 1.0 to 2.0 in go"
date = 2017-12-18T16:47:27+01:00
+++


# Considerations about software 1.0 and software 2.0

## What is a software?

It is sequence of bits and bytes that can be computed, and the produces a result (the solution of a problem for example). 

To build a software until now, a compiler is used. Its goal is to turn a "human readable sequence of characters", called code, into the sequence of bytes. 

This sequence of bytes is evaluated and executed by a machine at runtime. Depending of the input, the execution produces (hopefully) the expected output.

The art of programming, is, in essence, the faculty for a human to describe the solution of a problem and to express it in a computer language. 

## What is software 2.0?

Not so long ago, I discovered the concept of software 2.0 thanks to [Andrej Karpathy's blog](https://medium.com/@karpathy/software-2-0-a64152b37c35).
The idea is similar to any software: a compiler is used to turn a sequence of code into a sequence of bytes. This sequence is interpreted by a machine.

The difference is that the code is a sequence of mathematical equations (called model). Those equations are composed of variables and "constants". Let's call the constants "the weights".

The compiler is a software 1.0 that is able to transpile the equations into a sequence of bytes that will be evaluated by a machine (note that the compiler itself is a machine).

So what is the difference between 1.0 and 2.0? Is it just a matter of language?

No, the major difference is in the art of programming. 

On certain fields, a programmer __cannot__ write an algorithm that will __solve a specific problem__ (eg: I need to recognize a cat on any photo).

So, the programmer will write a set of equations __able to solve a kind of problem__ (recognize objects on any photo). 

The solution to your specific problem will be given by the evaluation of the _equation_ __and__ _the weights_ (a cat is an object that corresponds to the specific weights: {0,1,3,2,45,6,6,5,3,4,6,....}.)

And what makes the software 2.0 so specific, is the amount of weights that is so important that it cannot be determined manually. They are determined empirically. And a computer is faster than any human in this learning process.



_Sidenote about go_: I am a gopher and an Ops. I really like go because I find it easy and fun to do fancy stuffs. But actually go is not the first choice when we talk about machine learning. My goal is to write a kind a virtual machine for software 2.0. I will not explain this in details in this post why go, and anything about software 2.0; But the facilities offered by the go language in order to reach my goal.

# Context
In my last article, I have developped a recurrent neural network in pure go without any third party library.

As an example, I did an implementation of a character based generation (the famous Shakespeare example).
I have tried to tune the hyperparameters, but I haven't been able to reach a very usable text.

Actually, without any randomness in the generation process, the output was recurrent.

For example: 

```
Hello, The The The The The The ...
```


The point is that the toy I made is based on a vanilla RNN. And Vanillas RNNs are suffering from the [vanishing gradient problem](https://en.wikipedia.org/wiki/Vanishing_gradient_problem).
This is a well known problem, and one solution is to change the core model for a more robust network called __L__ong __S__hort __T__erm __M__emory network (LSTM for short).

# Implementing an LSTM

LSTM are a bit more complex than vanilla RNN. Therefore, a naive go implementation as made for the RNN will be a harder.

As one of my goal is to understand how things deeply works, I have tried to implement the back propagation mechanism manually.
I have read this post from Karpathy: [Yes you should understand backprop](https://medium.com/@karpathy/yes-you-should-understand-backprop-e2f06eab496b).

The best explanation I have found so far is in [cs231n course from Stanford](http://cs231n.github.io/optimization-2/).
It is a clear explanation of how the process works. And it is obvious that the graph representation helps a lot in the computation of the gradient.

I see now why tensorflow is so linked with the machine learning field.  

# Equations are graphs

So equations are graphs... Cool, I have always been attracted by the graphical representations. It is a very natural way to understand and express the ideas. This [post](http://gopherdata.io/post/deeplearning_in_go_part_1/) from [Chewxy](https://twitter.com/chewxy) is a perfect illustration of how the expression of a mathematical expression is turned into a graph at a compiler level.

It sounds that implementing the LSTM as a graph will make the task a lot easier. 

## Gorgonia

Chewxy is the author of the gorgonia project. 

> Package gorgonia is a library that helps facilitate machine learning in Go. Write and evaluate mathematical equations involving multidimensional arrays easily. Do differentiation with them just as easily.

## Writing software

### Considerations about software 1.0 and software 2.0

_Sidenote about go_: I am a gopher and an Ops. I really like go because I find it easy and fun to do fancy stuffs. But actually go is not the first choice when we talk about machine learning. My goal is to write a kind a virtual machine for software 2.0. I will not explain this in details in this post why go, and anything about software 2.0; But the facilities offered by the go language in order to reach my goal.

# Good ol' software 1.0

## Lexer/Parser

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

