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

**TL;DR:** Thank you for passing by. This article is, as usual, geek oriented. But if you are not a geek, and/or you are in hurry, you can jump to the conclusion: _[Any real application?](any-real-application)_

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

The idea is to get a video stream and grab pictures from this stream in order to activate a neural network.

I will present different objects in front of my webcam, and their name will be displayed on the screen.

The architecture is client server: The Chrome is the eye of my bot, it communicate with the brain (a webservice in go that is running a pre-trained tensorflow neural network) via a websocket.

**The rest of this paragraph is geek/javascript, if you're not interested you can jump to the next paragraph about the brain implementation called _[Cortical](#the-brain-cortical)_**

## getUserMedia

I am using the Web API [MediaDevices.getUserMedia()](https://developer.mozilla.org/en-US/docs/Web/API/MediaDevices/getUserMedia) to open the webcam and get the stream.

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

According to Wikipedia's definition, Websocket is _a computer communications protocol, providing full-duplex communication channels over a single TCP connection_.
The full duplex mode is important in my architecture. 

Let me explain why with a simple use case:

Imagine that your eye captures a scene and sends it to the brain for analysis. In a classic RESTfull architecture, the browser (the eye) would perform a POST request.
The brain would reply with a process ID, and the eye would poll the endpoint every x seconds to get the processing status.

This can be tedious in case of mutiple stimuli.

Thanks to the websocket, the server can send the query, and the server will send an event back once the processing is done.
Of course this is stateless in a sort, as the query is lost once the browser is closed.

Another use case would be to get a stimulus from another "sense". For example, imagine that you want to "warn" the end user that he has been mentioned in a tweet. The brain can be in charge of polling
twitter, and it would send a message through the websocket in case of event.

### Connecting to the websocket

A websocket URI is prefixed by `ws` or `wss` if the communication is encrypted (aka https).
This code allows a connection through ws(s).

{{< highlight js >}}
var ws
// Connecting the websocket
var loc = window.location, new_uri;
if (loc.protocol === "https:") {
  new_uri = "wss:";
} else {
  new_uri = "ws:";
}
new_uri += "//" + loc.host + "/ws";
ws = new WebSocket(new_uri);
{{</ highlight >}}

### Messages

Web socket communication is message oriented. A message can be sent simply by calling the function `ws.send(message)`. Websockets are supporting texts and binary messages.
But for this test only text messages will be used (images will be encoded in base64).

The browser implementation of a websocket in javascript is event based. 
When the server sends a message, an interruption is fired and the `ws.onmessage` call is triggered.

This code will display the message received on the console:

{{< highlight js >}}
ws.onmessage = function(event) {
  console.log("Received:" + event.data);
};
{{</ highlight >}}

### Sending pictures to the websocket: actually seeing

I didn't find a way to send the video stream to the brain via the websocket. Therefore I will do what everybody does: create a canvas and "take" a picture from the video:

The method [toDataURL()](https://developer.mozilla.org/en-US/docs/Web/API/HTMLCanvasElement/toDataURL) will take care of encoding the picture in a well known format (png or jpeg).

{{< highlight js >}}
function takeSnapshot() {
  var context;
  var width = video.offsetWidth
  , height = video.offsetHeight;

  canvas = canvas || document.createElement('canvas');
  canvas.width = width;
  canvas.height = height;

  context = canvas.getContext('2d');
  context.drawImage(video, 0, 0, width, height);

  var dataURI = canvas.toDataURL('image/jpeg')
  //...
};
{{</ highlight >}}

To make the processing easier in the brain, I will serialize the video into a json object and sending it via the websocket:

{{< highlight js >}}
var message = {"dataURI":{}};
message.dataURI.content = dataURI.split(',')[1];
message.dataURI.contentType = dataURI.split(',')[0].split(':')[1].split(';')[0]
var json = JSON.stringify(message);
ws.send(json);
{{</ highlight >}}

## Bonus: ear and voice

It is relativly easy to make chome speak out loud the message received. This snippet will speak out loud the message Received:

{{< highlight js >}}
function talk(message) {
  var utterance = new SpeechSynthesisUtterance(message);
  window.speechSynthesis.speak(utterance);
}
{{</ highlight >}}

Therefore, simply addind a call to this function in the "onmessage" event of the websocket will trigger the voice of Chrome. 

Listening to what is said is just a little bit trickier. It is done by a call to the `webkitSpeechRecognition();` method. 
This [blog post](https://developers.google.com/web/updates/2013/01/Voice-Driven-Web-Apps-Introduction-to-the-Web-Speech-API) explains in details how this works.

The call is also event based. What's important is that, in chrome, by default, it will use an API call to the Google's engine. Therefore the recognition won't work offline.

When the language processing is done by chrome, 5 potential sentences are stored in a json array.
The following snippet will take the most relevant one and send it to the brain via the websocket:


{{< highlight js >}}
recognition.onresult = function(event) { 
  for (var i = event.resultIndex; i < event.results.length; ++i) {
    if (event.results[i].isFinal) {
      final_transcript += event.results[i][0].transcript;
      ws.send(final_transcript);
    }
  }
};
{{</ highlight >}}

_Now that we have setup the senses, let's make a "brain"_

# The _brain_: **Cortical**

Now, let me explain what is, according to me, the **most interresting part** of this post.

## Concurrency

[Rob Pike - 'Concurrency Is Not Parallelism'](https://www.youtube.com/watch?v=cN_DpYBzKso) this is a *must see!*

## Sample cortex

### Locally with tensorflow

### In the cloud with GCP

### In the cloud with AWS

# Going deeper

# Any real application?

As usual the code can be found on [github under the tag v0.1](https://github.com/owulveryck/socketcam/releases/tag/v0.1)

