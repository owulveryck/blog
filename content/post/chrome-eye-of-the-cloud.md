---
categories:
- category
date: 2017-05-16T21:43:46+02:00
description: ""
draft: true
images:
- https://lh3.googleusercontent.com/nYhPnY2I-e9rpqnid9u9aAODz4C04OycEGxqHG5vxFnA35OGmLMrrUmhM9eaHKJ7liB-=w300
tags:
- tag1
- tag2
title: Chrome, the eye of the cloud - Computer vision with deep learning and only 2Gb of RAM
---

During the month of may, I have had the chance to attend to the Google Next event in London and the dotAI in Paris. In both conferences I learned a lot about machine learning. 

What those great speakers have taught me is that you shall not reinvent the wheel in AI. Actually a lot of research is done and there are very good implementation of the latest efficient algorithm.

*The tool* that every engineer that wants to try AI must know is [tensorflow](https://www.tensorflow.org/). Tensorflow is a generic framework that has been developed by Google's Machine Intelligence research organization. The tool has been open-sourced last year and has reached the v1.0 earlier this year.

## So what makes tensorflow so great?

### Bindings
First of all, it has bindings so it can be used within various programming languages such as:

* python
* c++
* java
* go

But to be honest, mainly python and c++ are described in the documentation. And to be even more honest I think that python is the language that you should use to prototype applications.

### ML and neuron network examples

Tensorflow is easy to use for machine learning, and a lot of deep-learning implementation are available.
Actually it is very easy to download a trained model and use it to recognize some pictures for example.

### Built-in computation at scale

Tensorflow's model has a built-in way to perform distributed computation. It is really important as machine learning is usually a very intensive task in term of computation.

### GCP's ML engine

Tensorflow is the engine used by Google for their service called ML engine.
That means that you can write your function locally and run them serverless on the cloud.
You only pay for what you have effectively consumed.
That means for example that you can train a neuron network on GCP (so you don't need GPU. TPU, or whatever computing power) and transfer your model locally,

For example, this is how the mobile app "google translate" works. A pre-trained model is downloaded on your phone, and the live translation is done locally.

![Image](http://technews.wpengine.netdna-cdn.com/wp-content/uploads/2015/01/www.lanacion.com_.ar_.jpg)

_Note_ The other ML services from GCP such as cloud vision, translate, or image search, are "just" API that query a neuron network with a model trained by google.

# So What?

I want to play with image recognition.
Actually I already did a test with AWS's rekognition service ([See this post](/2016/12/16/image-rekognition-with-a-webcam-go-and-aws..html)). But the problem were:

* I relied on a low-level webcam implementation, therefore the code was not portable
* I had no preview of what my computer was looking at 
* I could not execute it on any mobile app for a demo 

As I am using a Chromebook for a while, I found a solution: Using a Javascript API and the Chrome browser to access the camera. Then the pictures can be transfered to a backend via a websocket.
The backend would do the ML and reply with whatever information via the websocket. I can then display the result or even use the voice api of Chrome to tell the result loud.

# Chrome as the eye of the computer

The idea is to activate a video stream and process pictures from this stream to active a neuron network.

I will present different objects in from of my webcam, and their name will be displayed on the screen.

The architecture is client server: The Chrome is the eye of my bot, it communicate with the brain (a webservice in go that is running a pre-trained tensorflow neural network) via a websocket.

## getUserMedia

The Web API [MediaDevices.getUserMedia()](https://developer.mozilla.org/en-US/docs/Web/API/MediaDevices/getUserMedia) is used to open a webcam stream.

This API is compatible with chrome on desktop *and* mobile on Android phone (but not on iOS). This means that I will be able to use a mobile phone as an "eye" of my bot.

See the [compatibility matrix here](https://developer.mozilla.org/en-US/docs/Web/API/MediaDevices/getUserMedia#Browser_compatibility)

Here is the code to get access to the camera and display the video stream:

_html_
{{< highlight html >}}
<body>
  <video autoplay id="webcam"></video>
</body>
{{</ highlight >}}

_Javascript_
{{< highlight js >}}
// use MediaDevices API
// docs: https://developer.mozilla.org/en-US/docs/Web/API/MediaDevices/getUserMedia
if (navigator.mediaDevices) {
    // access the web cam
    navigator.mediaDevices.getUserMedia({video: true})
      // permission granted:
      .then(function(stream) {
          video.src = window.URL.createObjectURL(stream);
      })
      // permission denied:
      .catch(function(error) {
          document.body.textContent = 'Could not access the camera. Error: ' + error.name;
      });
}
{{</ highlight >}}

## Websockets

### Sending pictures to the websocket

### Getting info from the websocket and displaying it

## Other senses: hear and voice

# The _brain_: **Cortical**

## Sample cortex
### Locally with tensorflow

### In the cloud with GCP

### In the cloud with AWS

# Going deeper

As usual the code can be found on [github under the tag v0.1](https://github.com/owulveryck/socketcam/releases/tag/v0.1)

