---
title: "A simple face detection utility from Python to Go"
date: 2019-08-16T21:25:30+02:00
lastmod: 2019-08-16T21:25:30+02:00
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

<!--more-->

```
$ git clone https://github.com/azmathmoosa/azFace
$ cd azFce && tree | grep -v png
.
├── LICENSE
├── README.md
├── darknet.exe
├── data
│   └── labels
│       └── make_labels.py
├── net_cfg
│   ├── azface.data
│   ├── azface.names
│   └── tiny-yolo-azface-fddb.cfg
├── pthreadGC2.dll
├── pthreadVC2.dll
├── weights
│   └── tiny-yolo-azface-fddb_82000.weights
└── yolo_cpp_dll.dll
```

`git clone https://github.com/allanzelener/yad2k.git`

`conda env create -f environment.yml`



```bash
./yad2k.py \
        ../azFace/net_cfg/tiny-yolo-azface-fddb.cfg \
        ../azFace/weights/tiny-yolo-azface-fddb_82000.weights \
        ../FACES/keras/yolo2.h5
```

[The h5 version](https://drive.google.com/file/d/1O4BF8m3WrrHTIHnqFtl2oghaw_esRaYn/view)

```python
from keras.models import load_model
keras_model= load_model('../FACES/keras/yolo.h5')
keras_model.summary()
```

```
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

```python
from keras.models import load_model
import numpy as np
keras_model= load_model('../FACES/keras/yolo.h5')

output = keras_model.predict(np.zeros((1,416,416,3)))
np.save("../FACES/keras/output.npy",output)
```

```python
import onnxmltools
import onnx
from keras.models import load_model

keras_model= load_model('../FACES/keras/yolo.h5')
# Convert it! The target_opset parameter is optional.
import keras2onnx
onnx_model = keras2onnx.convert_keras(keras_model, name=None, doc_string='', target_opset=None, channel_first_inputs=None)
# onnx_model = onnxmltools.convert_keras(keras_model, target_opset=7)
# Save the ONNX model
onnx.save(onnx_model, '../FACES/yolo.onnx')
```

[Link to the model](https://github.com/owulveryck/gofaces/raw/master/model.onnx)

## Model visualisation

https://lutzroeder.github.io/netron/

<figure>
  <img src="/assets/yolofaces/netron-extract.png" link="/assets/yolofaces/netron.png">
  <figcaption>
      <h4>Netron representation of the tiny yolo v2 graph</h4>
  </figcaption>
</figure>

A copy of the full representation is [here](/assets/yolofaces/netron.png)

# Entering the Go world

```go
import (
        "github.com/owulveryck/onnx-go"
        "github.com/owulveryck/onnx-go/backend/x/gorgonnx"
        "gorgonia.org/gorgonia/encoding/dot"
        "gorgonia.org/tensor"
)

func main() {
        b, _ := ioutil.ReadFile("../FACES/yolo.onnx")
        backend := gorgonnx.NewGraph()
        model := onnx.NewModel(backend)
        model.UnmarshalBinary(b)

        t := tensor.New(
                tensor.WithShape(1, 416, 416, 3), 
                tensor.Of(tensor.Float32))
        _ = model.SetInput(0, t)

        exprGraph, _ := backend.GetExprGraph()
        b, _ := dot.Marshal(exprGraph)
        fmt.Println(string(b))
}
```

which gives:
<figure>
  <img src="/assets/yolofaces/onnx-gorgonia-preview.png" >
  <figcaption>
      <h4>Gorgonia representation of the tiny yolo v2 graph</h4>
  </figcaption>
</figure>

```go
func main(){
        //...
        must(backend.Run())
        outputT, err := model.GetOutputTensors()
        must(err)
        file, err := os.Open("../FACES/keras/output.npy")
        must(err)
        defer file.Close()

        expectedOutput := new(tensor.Dense)
        must(expectedOutput.ReadNpy(file))

        if ! eqInDelta(expectedOutput,outputT[0],5e-5) {
                log.Fatal("tensors differs")
        }
}
```
the comparison function is [here](https://gist.github.com/owulveryck/3d15c0eb9cf7dea6518116ec0a5be581#file-compare_tensor-go)

## Setting an input...


## Post-processing the output...


## Result
