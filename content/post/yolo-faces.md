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
        model_data/yolo2.h5
```

[The h5 version](https://drive.google.com/file/d/1O4BF8m3WrrHTIHnqFtl2oghaw_esRaYn/view)

```python
from keras.models import load_model
keras_model= load_model('../../FACES/keras/yolo.h5')
from keras.utils import plot_model
plot_model(keras_model, to_file='model.png')
```

<figure>
  <img src="/assets/yolofaces/keras-preview.png"
  <figcaption>
      <h4>Keras representation of the tiny yolo v2 graph</h4>
  </figcaption>
</figure>


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
onnx.save(onnx_model, '../FACES/yolo2.onnx')
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
