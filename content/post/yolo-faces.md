---
title: "A simple face detection utility from Python to Go"
date: 2019-08-16T21:25:30+02:00
lastmod: 2019-08-16T21:25:30+02:00
draft: true
keywords: ["onnx","DDD","Keras","Go","Neural Net","YOLO"]
description: "This post describes how to build a face detection tool with neural net. The full conception is described, from the design to the implementation."
tags: []
categories: []
author: ""

# You can also close(false) or open(true) something for this content.
# P.S. comment can only be closed
comment: true
toc: true
autoCollapseToc: true
# You can also define another contentCopyright. e.g. contentCopyright: "This is another copyright."
contentCopyright: false
reward: false
mathjax: false
---

In this article, I explain how to build a tool to detect faces in a picture.
The goal of this post is:

* to build the business model thanks to a neural network;
* to adapt the network to the specific domain of face detection by changing its knowledge;
* to use the resulting domain with a go-based infrastructure;
* to code a little application in Go to communicate with the outside world.

This article can be considered as a sort of how-to use a keras model in Go.
Therefore, in this post I will use the following technologies:

* Python / Keras
* ONNX
* Go

**Note**: Some of the terms such as _domain_, _application_, and _infrastructure_ refer to the concepts from Domain Driver Design (DDD) or the hexagonal architecture. For example, do not consider the infrastructure as boxes and wires, but see it as a service layer. The infrastructure represents everything that exists independently of the application. 

**Disclaimer**: I am using those concepts to illustrate what I do; This is not a proper DDD design nor a true hexagonal architecture.


The architecture of our tool can be described by three layers:
<center>
<figure>
  <img src="/assets/yolofaces/archi1.png" >
  <figcaption>
      <h4>Overall picture of the architecture</h4>
  </figcaption>
</figure>
</center>

The basic principle is that every layer is a "closed area"; therefore it is accessed through API and every layer can be tested independently.
Therefore each layer is described in a paragraph of this article.

The "actor" here is a simple cli tool. It is the main package of the application (and in go the main package is the package `main`); In the rest of the article, I will reference it as "**the actor**.

# Implementing the business logic with a neural network

The core functionality of the tool is to detect faces on a picture.
I will use a neural network to achieve this. The model I have chosen is 
[Tiny YOLO v2](https://pjreddie.com/darknet/yolov2/) what is able to perform real time object detection.

> This model is designed to be small but powerful. It attains the same top-1 and top-5 performance as AlexNet but with 1/10th the parameters. It uses mostly convolutional layers without the large fully connected layers at the end. It is about twice as fast as AlexNet on CPU making it more suitable for some vision applications.

I am using the "tiny" version which is based off of the Darknet reference network and is much faster but less accurate than the normal YOLO model.

The model is just an “envelope.” It needs some training to be able to detect some objects. The objects it can detect is dependant of its knowledge. The weights tensors represent its knowledge.
To detect faces, we need to apply the model to the picture with a knowledge (some weights) able to recognize faces.


### Getting the weights
By luck, an engineer named [Azmath Moosa](https://github.com/azmathmoosa) has trained the model and released a tool called [azface](https://github.com/azmathmoosa/azFace).
The project is available on GitHub in LGPLv3 but, it does not contain the sources of the tool (only a Windows binary and some DLL are present). However what I am interested in is not the tool as I am building my own. What I am seeking now is the weights; and the weights are present in the repository as well.

_Disclaimer_: my tool is for academic purpose, I am confident that his tool is much better.

Let's clone the repository:

`$ git clone https://github.com/azmathmoosa/azFace`

The weights are this heavy file of 61Mb: `weights/tiny-yolo-azface-fddb_82000.weights`.

### Combining the weights and the model

Now, we need to combine the knowledge and the model. Together, they constitute the core functionality of our domain.

A business logic should be as independant as possible of any framework. The best way to represent the neural network is to be as close as possible as 
its definition; The original implementation of the yolo model (from "darknet") is written in C; There are other reimplementation in Tensorflow, Keras, Java, ...

I will use [ONNX](https://onnx.ai/) as a format for the business logic; so I will be independant of a framework, and I will be able to use the power of different _infrastructures_.

To create the ONNX format, I will use Keras with the tools: 

* [`yad2k`](https://github.com/allanzelener/yad2k.git) to create a keras model;
* [`keras2onnx`](https://pypi.org/project/keras2onnx/) to encode it into ONNX.


The workflow is:

```
                          yad2k                   keras2onnx               
darknet config + weights -------->  keras model --------------> onnx model
```

let's create a keras model from the config and the weights of `azface`
```bash
./yad2k.py \
        ../azFace/net_cfg/tiny-yolo-azface-fddb.cfg \
        ../azFace/weights/tiny-yolo-azface-fddb_82000.weights \
        ../FACES/keras/yolo2.h5
```

This generates a pre-trained [h5 version](https://drive.google.com/file/d/1O4BF8m3WrrHTIHnqFtl2oghaw_esRaYn/view) of the tiny yolo v2 model, able to find faces.

Let's analyze it:
```python
from keras.models import load_model
keras_model= load_model('../FACES/keras/yolo.h5')
keras_model.summary()
```

```txt
_________________________________________________________________
Layer (type)                 Output Shape              Param #   
=================================================================
input_1 (InputLayer)         (None, 416, 416, 3)       0         
_________________________________________________________________
conv2d_1 (Conv2D)            (None, 416, 416, 16)      432       
_________________________________________________________________
...
_________________________________________________________________
conv2d_9 (Conv2D)            (None, 13, 13, 30)        30750     
=================================================================
Total params: 15,770,510
Trainable params: 15,764,398
Non-trainable params: 6,112
_________________________________________________________________
```

Sounds ok!

### Generate the onnx file

The onnx version is generated with the tool keras2onnx with this script:
```python
import onnxmltools
import onnx
import keras2onnx
from keras.models import load_model

keras_model= load_model('../FACES/keras/yolo.h5')
onnx_model = keras2onnx.convert_keras(keras_model, name=None, doc_string='', target_opset=None, channel_first_inputs=None)
onnx.save(onnx_model, '../FACES/yolo.onnx')
```

I have uploaded the result [here](https://github.com/owulveryck/gofaces/raw/master/model.onnx)

#### Model visualisation

It is interesting to see the result. I am using the tool `netron` which have a [web version](https://lutzroeder.github.io/netron/).
The result looks like this:

<center>
<figure>
  <img src="/assets/yolofaces/netron-extract.png" link="/assets/yolofaces/netron.png" width="50%">
  <figcaption>
      <h4>Netron representation of the tiny yolo v2 graph</h4>
  </figcaption>
</figure>
</center>

I made a copy of the full representation which is available [here](/assets/yolofaces/netron.png)

#### Preparing the test of the infrastructure

In order to validate our future infrastructure, let's prepare a simple test:
Apply the model on a zero value and save the result. I will do the same once the final infrastructure is up and compare the result.

```python
from keras.models import load_model
import numpy as np
keras_model= load_model('../FACES/keras/yolo.h5')

output = keras_model.predict(np.zeros((1,416,416,3)))
np.save("../FACES/keras/output.npy",output)
```

Now, let's move to the infrastructure and application part.

# Infrastructure: Entering the Go world

No surprise here: the infrastructure I am using is made of [`onnx-go`](https://github.com/owulveryck/onnx-go) to decode the onnx file,
and [Gorgonia](https://github.com/gorgonia/gorgonia) to execute the model.
This solution is an efficient solution for a tool; all the dependencies I needed to build the pre-trained model are not needed anymore. This gives the end-user of the tool a much better experience.

### The Service Provider Interface (SPI)

We've seen that the neural network is represented by its model. The SPI should implement a model to fulfill the contract and understand the ONNX Intermediate Representation (IR). [Onnx-go](https://github.com/owulveryck/onnx-go)'s [`Model`](https://godoc.org/github.com/owulveryck/onnx-go#Model) object is a Go structure that acts as a receiver of the neural network model.

The other service required is a computation engine that understands and executes the model. This function is assumed by [Gorgonia](https://github.com/gorgonia/gorgonia).

The **actor** will use those services. A basic implementation in Go is (note the package is `main`):

```go
import (
        "github.com/owulveryck/onnx-go"
        "github.com/owulveryck/onnx-go/backend/x/gorgonnx"
)

func main() {
        b, _ := ioutil.ReadFile("../FACES/yolo.onnx")
        backend := gorgonnx.NewGraph()
        model := onnx.NewModel(backend)
        model.UnmarshalBinary(b)
}
```


To use the model, we need to interact with its inputs and output.
The model takes a tensor as input. To set it onnx-go provides a helper function: `SetInput`.
The outputs are obtained via a call to `GetTensorOutput()`

```go
t := tensor.New(
        tensor.WithShape(1, 416, 416, 3), 
        tensor.Of(tensor.Float32))
model.SetInput(0, t)
```

The **actor** could use those methods, but, as the goal of the application is to analyze pictures, the application will encapsulate them to provide a richer user experience for the actor (the actors will probably not want to mess up with tensors).

#### Testing the infrastructure 

We can now test the infrastructure to see if the implementation is ok. We set an empty tensor, compute it with Gorgonia and compare the result with the one
saved previously:

I wrote a small `test` file in the go format; for clarity I will not copy/paste it here, but host on a [gist](https://gist.github.com/owulveryck/3d15c0eb9cf7dea6518116ec0a5be581#file-yolo_test-go). 

```text
# go test
PASS
ok      tmp/graph       1.054s
```

_Note_: The ExprGraph used by gorgonia can also be represented visually with graphviz. This code generates the _dot_ representation:

```go
exprGraph, _ := backend.GetExprGraph()
b, _ := dot.Marshal(exprGraph)
fmt.Println(string(b))
}
```

(the full graph is [here](/assets/yolofaces/yolo-gorgonia.png))

<center>
<figure>
  <img src="/assets/yolofaces/onnx-gorgonia-preview.png" width="50%">
  <figcaption>
      <h4>Gorgonia representation of the tiny yolo v2 graph</h4>
  </figcaption>
</figure>
</center>

The infrastructure is ok, and is implementing the SPI! Let's move to the application part!

# Writing the application in Go

## The API

Let's start with the interface of the application. I create a package `gofaces` that will hold the logic of the application.
It will be a layer that will add some facilities to communicate with the outside world. This package can be instanciated by a anything from a simple cli to 
a webservice.

### Input

#### GetTensorFromImage

This function takes an image as input; The image is transfered to the function with a stream of bytes (`io.Reader`). This let the possibility for the end user
to use a regular file, to get the file from stdin, or to build a webservice and get the file via http.
This function returns a tensor usable with the yolo faces model; it also returns any error if it cannot process the file.

_Note_ the full signature of the `GetTensorFromImage` function can be found on [GoDoc](https://godoc.org/github.com/owulveryck/gofaces#GetTensorFromImage) 

We can take back the **actor** implementation and add the input picture (I skip the errors checking for clarity):

```go
func main() {
        b, _ := ioutil.ReadFile("../FACES/yolo.onnx")
        // Instanciate the infrastructure
        backend := gorgonnx.NewGraph()
        model := onnx.NewModel(backend)
        // Loading the business logic (the neural net)
        model.UnmarshalBinary(b)
        // Accessing the I/O through the API
        inputT, _ := gofaces.GetTensorFromImage(img)
        model.SetInput(0, inputT)
}
```

The model can be executed with a call to [`backend.Run()`] because Gorgonia fulfills the [`ComputationBackend`](https://godoc.org/github.com/owulveryck/onnx-go/backend#ComputationBackend) interface.

### Output

#### Bounding boxes

The model outputs a tensor. This tensors holds all of the informations to extract bounding boxes. 
Getting the bounding boxes is the responsibility of the application. Therefore, the package `gofaces` defines a [`Box`](https://godoc.org/github.com/owulveryck/gofaces#Box) structure.  
A box contains a set of [`Elements`](https://godoc.org/github.com/owulveryck/gofaces#Element)

#### Get the bounding boxes


TODO 
* [`ProcessOutput`](https://godoc.org/github.com/owulveryck/gofaces#ProcessOutput)
* [`Sanitize`](https://godoc.org/github.com/owulveryck/gofaces#Sanitize)

# Final result

You can find the code of the application in my [`gofaces`](https://github.com/owulveryck/gofaces) repository.

I am using a famous meme as input (you can find the image [here](/assets/yolofaces/meme.jpg))
<center>
<figure>
  <img src="/assets/yolofaces/meme.jpg" width="30%">
</figure>
</center>


```shell
cd $GOPATH/src/github.com/owulveryck/gofaces/cmd
go run main.go \
        -img /tmp/meme.jpg \
        -model ../model/model.onnx
```

gives the following result
```text
[At (187,85)-(251,147) (confidence 0.20):
        - face - 1
]
```

So it has detected only one face; It is possible to play with the confidence threshold to detect other faces.
I have found that it is not possible to detect the face of the _lover_; probably because the picture do not show her full face.

### Getting an output picture


```shell
YOLO_CONFIDENCE_THRESHOLD=0.1 go run main.go \
        -img /tmp/meme.jpg \
        -output /tmp/mask2.png \
        -model ../model/model.onnx
convert \
        /tmp/meme.jpg \
        /tmp/mask2.png \
        \( -resize 418x \) \
        -compose over -composite /tmp/result2.png
```

<center>
<img src="/assets/yolofaces/mask2.png" width="30%" style="border-width: 1px;border-color: black;border-style: solid;">
<img src="/assets/yolofaces/result2.png" width="30%">
</center>


# Conclusion
