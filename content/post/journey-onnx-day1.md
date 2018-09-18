---
title: "My journey with ONNX and Go - The begining"
date: 2018-08-14T20:41:30+02:00
lastmod: 2018-08-14T20:41:30+02:00
draft: true
keywords: []
description: "This is the very begining of my journey with ONNX and Go"
tags: []
categories: []
author: ""

# You can also close(false) or open(true) something for this content.
# P.S. comment can only be closed
comment: false
toc: false
autoCollapseToc: false
contentCopyright: false
reward: false
mathjax: false
---

This year has started with a lot of deep thoughts about the software 2.0.
My conclusion (which is slightly different from [Andrej Karpathy's consideration](https://medium.com/@karpathy/software-2-0-a64152b37c35)) is that a software 2.0 is a combination of a Neural network model **and** its associated weights.
This is a concept; now the question is: how to materialize the concept? What artifact represents a software 2.0.

I emit several ideas and tried one of them: to serialize the mathematical model and the weights.
The main drawback of this idea is that it is not easy to write down and to parse any mathematical equation.
The best way to express a model is, as of today, via its computation graph (this is what most ML frameworks are doing). 

Therefore, switching from a mathematical representation to the computation graph representation might lead to a  good way to express the artifact of a software 2.0.

# Quick word about ONNX

Describing computation graph in straightforward. A computation graph is a [Direct Acyclic Graph (DAG)](https://en.wikipedia.org/wiki/Directed_acyclic_graph). Each node  of the graph represents a tensor or an operator.
The challenge is to find a domain specific language (DSL) to describe a graph in a way that it is agnostic of its implementation.

This is the promise of ONNX.
[ONNX](https://onnx.ai/) stands for Open Neural Network eXchange (format). The purpose of this project is to establish an open standard for exporting/importing ML models.

<center>
{{< figure src="https://github.com/owulveryck/onnx-go/raw/master/vignettes/imgs/ONNX_logo_main.png" >}}
</center>

In this post, I will describe the first step I have made in order to be able to read (and hopefully) execute an ML model encoded via ONNX into the Go ecosystem.

# From the protobuf definition to a Go structure

## What are protocol buffers

According to the [website](https://developers.google.com/protocol-buffers/)

> Protocol buffers are Google's language-neutral, platform-neutral, extensible mechanism for serializing structured data – think XML, but smaller, faster, and simpler. You define how you want your data to be structured once, then you can use special generated source code to easily write and read your structured data to and from a variety of data streams and using a variety of languages.

Protocol buffer is a binary format (once compiled it cannot be read by a human). It is a way to serialize messages. For short, a protobuf file describes an API contract. 

I will not go deeper in the protobuf description here. In my humble opinion, it is a very good way to express an API when implementing a machine-to-machine communication. Better than JSON because of it simplicity, efficiency and the ability to validate a schema natively.

The main definition file for ONNX (the API contract) is hosted [here](https://github.com/onnx/onnx/blob/master/onnx/onnx.proto3) and is named `onnx.proto3`.
This file is used to generate bindings to other languages.

In order to create a bridge between the protobuf binary format and the Go ecosystem, the first thing to do is to generate the Go API. This will allow to read a ONNX file and to transpile it into a Go compatible object.

To do this, you need a compiler called protoc. I am also using the alternative compiler [gogoprotobuf](https://github.com/gogo/protobuf) which add some useful features (such as fast Mashaller/Unmarshaler methods).

_Note:_ For clarity, I will not describe how to install the `protoc` binary

Simply running `protoc --gofast_out=. onnx.proto3` will generate a file [onnx.pb.go](https://github.com/owulveryck/onnx-go/blob/master/onnx.pb.go) which is usable out-of-the box.

## onnx-go 

After some discussions with the [official team](https://github.com/onnx/onnx/pull/1328), we agreed that, before reaching a certain maturity, it was best to host it on my personal github account. So you can find the definition file [here](https://godoc.org/github.com/owulveryck/onnx-go).

The corresponding Godoc is hosted [here](https://godoc.org/github.com/owulveryck/onnx-go)

This package on its own is enough to read a ONNX format. 

### Testing the package

The ONNX organization has setup a [model repository](https://github.com/onnx/models). From this repository, let's extract the basic MNIST example.

`curl https://www.cntk.ai/OnnxModels/mnist/opset_7/mnist.tar.gz | tar -C /tmp -xzvf -`

Now, let's write a simple program that will read the ONNX file and decode it into the [`ModelProto` Object](https://godoc.org/github.com/owulveryck/onnx-go#example-ModelProto) (which is the top level object in the ONNX file).

Then create a very simple Go program to read and dump the model:

{{< highlight go >}}
import (
        "io/ioutil"
        "log"
        onnx "github.com/owulveryck/onnx-go"
        "github.com/y0ssar1an/q"
)

func main() {
        b, err := ioutil.ReadFile("/tmp/mnist/model.onnx")
        if err != nil {
                log.Fatal(err)
        }
        model := new(onnx.ModelProto)
        err = model.Unmarshal(b)
        if err != nil {
                log.Fatal(err)
        }
        q.Q(model)
}
{{</ highlight >}}

_Note_: I am using the [`q`](https://github.com/y0ssar1an/q) package to dump the content as the output is verbose. The result is present in the file `$TMPDIR/q`

# From the Go structure to a Graph

# Graphs

Computation graphs are made up of a DAG of nodes, which represent what is commonly called a "layer" or "pipeline stage" in machine learning frameworks.

## Gonum

![Graph mnist](/assets/onnx/mnist.png)

# Switching to Gorgonia

## Decoding the tensor



<center>
![0](/assets/onnx/0.png)
</center>

## Creating an ExprGraph

## Running the graph


# Conclusion
