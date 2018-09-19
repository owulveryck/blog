---
title: "My journey with ONNX and Go - Running the graph"
date: 2018-09-19T08:53:09+02:00
lastmod: 2018-09-19T08:53:09+02:00
draft: true
keywords: []
description: ""
tags: []
categories: []
author: ""

# You can also close(false) or open(true) something for this content.
# P.S. comment can only be closed
comment: true
toc: true
autoCollapseToc: false
# You can also define another contentCopyright. e.g. contentCopyright: "This is another copyright."
contentCopyright: false
reward: false
mathjax: false
---

<!--more-->

In the previous post, I have made an introduction and a POC to interact with ONNX models and Go.

I have decoded the information to reconstruct a graph.
Now I propose to expand the principle and to create a proper execution backend based on Gorgonia.
This post is a bit more technical than the previous one as I don't have any new concept to explain. 

# Decoding the tensor

In machine learning, the basic element of a computation graph is a Tensor.
In ONNX this element is described in the structure [TensorProto](https://godoc.org/github.com/owulveryck/onnx-go#TensorProto). 
A tensor has a shape represented here by the field `Dims` which is an array of int64, is holdind a data type and obviously some data.

Gorgonia has also a notion of tensor. A tensor is an interface in Gorgonia, therefore, creating a Go object from TensorProto that fulfils the Tensor interface of Gorgonia
should be easy.

Let's write a method that taks a onnx.TensorProto as input and that returns a tensor.Tensor as output

{{< highlight go >}}
func NewTensor(tx *onnx.TensorProto) (tensor.Tensor, error) { ... } 
{{</ highlight >}}

## Testing the tensor

```
curl https://www.cntk.ai/OnnxModels/mnist/opset_7/mnist.tar.gz | \
tar -C /tmp -xzvf -
```

{{< highlight go >}}
b, _ := ioutil.ReadFile("/tmp/mnist/test_data_set_0/input_0.pb")
sampleTestData := new(onnx.TensorProto)
sampleTestData.Unmarshal(b)
t, _ := NewTensor(sampleTestData)
width := t.Shape()[2]
height := t.Shape()[3]
im := image.NewGray(image.Rectangle{Max: image.Point{X: width, Y: height}})
for w := 0; w < width; w++ {
        for h := 0; h < height; h++ {
                v, _ := t.At(0, 0, w, h)
                im.Set(w, h, color.Gray{uint8(v.(float32))})
        }
}
enc := png.Encoder{}
enc.Encode(os.Stdout, im)
{{</ highlight >}}
<center>


![0](/assets/onnx/0.png)
</center>

# Creating an ExprGraph

# Running the graph

## Sample output

{{< figure src="/assets/onnx/mnist_gorgonia.svg" title="MNIST Model with Gorgonia" >}}

# Conclusion
