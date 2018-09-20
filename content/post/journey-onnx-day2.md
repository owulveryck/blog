---
title: "My journey with ONNX and Go - Running the graph"
date: 2018-09-19T08:53:09+02:00
lastmod: 2018-09-19T08:53:09+02:00
draft: true
keywords: []
description: "This post is the second part of my experiments with ONNX and Go. In this post I am describing how to create a computation graph in Gorgonia (ExprGraph) from a ONNX Model."
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

In the [previous post](/2018/08/14/my-journey-with-onnx-and-go---the-begining.html#building-the-dag), I made an introduction and a POC to interact with ONNX models and Go.

I have decoded the information to reconstruct a graph.
Now I propose to expand the principle and to create a proper execution backend based on Gorgonia.
This post is a bit more technical than the previous one as I don't have any new concept to explain. 

# Decoding the tensor

In machine learning, the basic element of a computation graph is a Tensor.
In ONNX this element is described in the structure [TensorProto](https://godoc.org/github.com/owulveryck/onnx-go#TensorProto). 
A tensor has a shape represented here by the field `Dims` which is an array of int64, is holdind a data type and obviously some data.

Gorgonia has also a notion of tensor. A tensor is an interface. Therefore, creating a Go object from TensorProto that fulfils the Tensor interface of Gorgonia
should be easy.

Let's write a method that taks a `onnx.TensorProto` as input and that returns a `tensor.Tensor` as output

{{< highlight go >}}
func NewTensor(tx *onnx.TensorProto) (tensor.Tensor, error) { ... } 
{{</ highlight >}}

We need to address the thre elements:

* convert the data type to something understandable by Go (and Gorgonia)
* read and process the data to write a tensor backend
* deal with tensor shape.

I will not focus much on tensor shape. Actually ONNX has a notion of dimension which is an array of integer. Every entry represent the size of an axe of the tensor.
This can be converted out-of-the-box into a [`Shape`](https://godoc.org/gorgonia.org/tensor#Shape) element of the `tensor` package.

The data type conversion and the raw data processing is a (little) bit trickier, so let's focus on them.

### Data types

A tensor is composed of elements of certain types. The supported data types are described as constants in ONNX. They can be found [in the documentation of ONNX](https://github.com/onnx/onnx/blob/master/docs/IR.md#standard-data-types) and are represented in [Go constant values](https://godoc.org/github.com/owulveryck/onnx-go#TensorProto_DataType) of our Go API.

On the other hand, the tensor package of gorgonia has also its own declaration of types represented by values of [`Dtypes`](https://godoc.org/gorgonia.org/tensor#Dtype). The list is a set of variables declared [here](https://godoc.org/gorgonia.org/tensor#pkg-variables).

Writing a function to return a `Dtype` from a `TensorProto_DataType` is relatively straightforward: 

{{< highlight go >}}
func Dtype(t *onnx.TensorProto_DataType) (tensor.Dtype, error) {
	switch *t {
	case onnx.TensorProto_FLOAT:
		return tensor.Float32, nil
        //...
{{</ highlight >}}

### Raw Data

ONNX has two way to encode the data of a tensor.
The first is really easy and is a straigh serialization of the basic type. For example, a tensor of type Float32 will have its data set in the `FloatData` field which is of type `[]float32`.

The second one is a bit trickier. ONNX allows to serialize the "raw data" encoded in a sequence of bytes. The documentation says that:

> When this raw_data field is used to store tensor value, elements MUST
> be stored in as fixed-width, little-endian order.
> Floating-point data types MUST be stored in IEEE 754 format.
> Complex64 elements must be written as two consecutive FLOAT values, real component first.
> Complex128 elements must be written as two consecutive DOUBLE values, real component first.
> Boolean type MUST be written one byte per tensor element (00000001 for true, 00000000 for false).
>
> Note: the advantage of specific field rather than the raw_data field is
> that in some cases (e.g. int data), protobuf does a better packing via
> variable length storage, and may lead to smaller binary footprint.
> When this field is present, the data_type field MUST NOT be STRING or UNDEFINED

So our function must handle this special case.
Let's focus on the Float32 type for now. Go has natively everything needed to read this famous `IEEE 754 format` (thanks to the binary and the math packages).

Here is how to read the informations and to transcribe it into a `[]float32`:

{{< highlight go >}}
buf := bytes.NewReader(tx.RawData)
element := make([]byte, 4)
var backing []float32
for {
        var n int
        n, err = buf.Read(element)
        if err != nil || n != 4 {
                break
        }
        uintElement := binary.LittleEndian.Uint32(element)
        backing = append(backing, math.Float32frombits(uintElement))
}
{{</ highlight >}}


## Vizualizing the tensor

Let's take back the MNIST example from the ONNX model zoo.
```
curl https://www.cntk.ai/OnnxModels/mnist/opset_7/mnist.tar.gz | \
tar -C /tmp -xzvf -
```

The model is delivered with three tests. The tests are made of an input tensor and the expected output tensor.
Let's take one of the input tensor, convert it to a Gorgonia tensor and create a picture from it (so see if the data, types and shapes are coherents).
I am using the `image` package of the standard Go distribution and dumping a png file on stdout for commodity:

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

Running the code produces a `0` as expected:

<center>
{{< figure src="/assets/onnx/0.png" height="200%" title="Representation of a zero from a tensor" >}}
</center>

# Creating an ExprGraph

Now that we are able to decode the tensors from the ONNX model, let's go further and create a graph.
In the previous post, we have sliced the parsing function into three parts:

* the processing of the _Initializers_
* the processing of the _Inputs_
* the processing of the _Operators_

(cf [_Building the DAG_](/2018/08/14/my-journey-with-onnx-and-go---the-begining.html#building-the-dag) in the previous post for more information)

### Constraints with the broadcastable operators

# Running the graph

## Sample output

{{< figure src="/assets/onnx/mnist_gorgonia.svg" title="MNIST Model with Gorgonia" >}}

# Conclusion
