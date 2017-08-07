---
date: 2017-07-07T21:06:46+02:00
description: "Imagine a CCTV at home that would trigger and alert when it detects a movement. Ok, this is easy. Imagine a CCTV that would trigger an alert when it detects a human. A little bit trickier. Now imagine a CCTV that would trigger an alert when it sees someone who is not from the family."
draft: true
images:
- /assets/images/tensorflowserving-4.png
title: A "Smart" CCTV with Tensorflow, and Inception? On a rapsberry pi?
---

Imagine a CCTV at home that would trigger and alert when it detects a movement. 

Ok, this is easy. 

Imagine a CCTV that would trigger an alert when it detects a human (and not the cat). 

A little bit trickier. 

Now imagine a CCTV that would trigger an alert when it sees someone who is not from the family...

__Disclaimer__: This article will not cover everything. I may post a second article later (or not). As you may now, I am doing those experiments during the night as all of this is not part of my job. I hope that I will find time to actually conclude the experiment. If you are a geek and you want to test that with me, feel free and welcome to contact me via the comments or via twitter [@owulveryck](https://twitter.com/owulveryck).
In this article, I will describe the method. I will also retrain a neural network to detect people. I will also use a GO static binary to run it live and evaluate the performances. By the end, I will try a static cross compilation to run it on a raspberry pi, but as my rpi is by now out-of-order, I will test it on qemu.

# Machine learning and computer vision

Machine learning and tooling around it has increased and gain in efficiency in the past years. it now "easy" to code a model that can be trained to detect and classify elements from a picture. 

Cloud providers are offering services that can instantly tag and label elements from an image. To achieve the goal of the CCTV, it would be really easy to use, for example, [AWS rekognition](https://aws.amazon.com/rekognition/), train the model, and post a request for each image seen.

This solution presents a couple of problems:

* The network bandwidth: you need a reliable network and a bandwidth big enough to upload the flow of images

* The cost: these services are cheap for thousand images, but consider about 1 fps to process (I don't even dream of 24fps), it is 86400 images a day and 2.6 million images a month... and considering that 1 million images are 1000 dollar...

I don't even talk about network latency because my CCTV would be pseudo-real-time and the ms of latency can be neglected.

The best solution would be to run the computer vision locally. There are several methods to detect people. The most up-to-date-and-accurate one is based on machine learning and precisely on neural network.

The (really simplified) principle is: 

To describe a mathematical model full of parameters (a kind of giant equation). This model will take a lot of labeled pictures as input (this picture contains a person, this one not, etc..). Every time a picture is processed, it tries to find the best parameter for the equation such as:

if this is a person, f(x) = 1 otherwise f(x) = 0.

The pictures used to feed the model is called the training set. 
You also use a test set (same kind of pictures), that is used to check whether your model generalized well and actually converge to your goal.

<center>
{{< figure src="https://imgs.xkcd.com/comics/machine_learning.png" link="https://xkcd.com/1838/" caption="XKCD 1838" >}}
</center>

## Tools

### Tensorflow

I have already told about tensorflow. Tensorflow is not a ML library. Tensorflow is a mathematical library. It self-describes itself as an  _open source software library for numerical computation using data flow graphs. Nodes in the graph represent mathematical operations, while the graph edges represent the multidimensional data arrays (tensors) communicated between them._


### Inception

"[Inception](https://research.google.com/pubs/pub43022.html)" is a deep convolutional neural network architecture used to classify images originally developed by Google.

Inception is exceptionally accurate for computer vision. It can reach 78% accuracy in "Top-1" and 93.9% in "Top-5". That means that if you feed the model with a picture of sunglasses, you have 93.9% chance that the algorithm detects sunglasses amongst the top 5 results.

On top of that, Inception is implemented with Tensorflow, and well documented. Therefore, it easy "easy" to use it, to train it and "to retrain it".
Actually, training the model is a very long process (several days on very efficient machines with GPU). 

<center>
{{< figure src="https://raw.githubusercontent.com/tensorflow/models/master/inception/g3doc/inception_v3_architecture.png" likn="https://github.com/tensorflow/models/tree/master/inception" caption="Inception v3 architecture" >}}
</center>

# Geek

I am using the excellent blog post [How to Retrain Inception's Final Layer for New Categories](https://www.tensorflow.org/tutorials/image_retraining).

## Phase 1: recognizing usual people

To keep it simple, I've created a "class" people with the flowers classes. It means that I simple added a directory "people" to my "flowers" for now.

{{< tweet 861831985437827072 >}}

```
[~/flower_photos]$ ls -lrt
total 696
-rw-r----- 1 ubuntu ubuntu 418049 Feb  9  2016 LICENSE.txt
drwx------ 2 ubuntu ubuntu  45056 Feb 10  2016 tulips
drwx------ 2 ubuntu ubuntu  36864 Feb 10  2016 sunflowers
drwx------ 2 ubuntu ubuntu  36864 Feb 10  2016 roses
drwx------ 2 ubuntu ubuntu  57344 Feb 10  2016 dandelion
drwx------ 2 ubuntu ubuntu  36864 Feb 10  2016 daisy
drwxrwxr-x 2 ubuntu ubuntu  77824 Jul  7 14:26 people
```

### Getting a training set full of people

I download pictures of people from [http://www.image-net.org/api/text/imagenet.synset.geturls?wnid=n07942152](http://www.image-net.org/api/text/imagenet.synset.geturls?wnid=n07942152)

{{< highlight shell >}}
curl -s  "http://www.image-net.org/api/text/imagenet.synset.geturls?wnid=n07942152" | \
sed 's/^M//' | \
while read file
do
  curl -m 3 -O $file
done
{{</ highlight >}}

Then I remove all "non image" files:

{{< highlight shell >}}
for i in $(ls *jpg)
do
    file $i | egrep -qi "jpeg|png" || rm $i 
done
{{</ highlight >}}

## Learning phase

I am following the tutorial to [Using the retrained model](https://www.tensorflow.org/tutorials/image_retraining#using_the_retrained_model). Therefore I am sure that my installation is working ok.

I had one "issue" because I am using python3. When I executed 

`bazel-bin/tensorflow/examples/image_retraining/retrain --image_dir ~/flower_photos/` 

it failed with a message about `ModuleNotFoundError: No module named 'backports'`. I Googled and found the solution in this [issue](https://github.com/tensorflow/serving/issues/489#issuecomment-313671459).

At the end of the training (which took 12 minutes on a c4.2xlarge spot instance on AWS) I have two files that holds the previous informations.

```
...
2017-07-07 19:22:53.667219: Step 3990: Cross entropy = 0.111931
2017-07-07 19:22:53.728059: Step 3990: Validation accuracy = 93.0% (N=100)
2017-07-07 19:22:54.287266: Step 3999: Train accuracy = 98.0%
2017-07-07 19:22:54.287365: Step 3999: Cross entropy = 0.148188
2017-07-07 19:22:54.348603: Step 3999: Validation accuracy = 91.0% (N=100)
Final test accuracy = 92.7% (N=492)
Converted 2 variables to const ops.
...

(customenv) *[r1.2][~/sources/tensorflow]$ ls -lrth /tmp/output_*
-rw-rw-r-- 1 ubuntu ubuntu  47 Jul  7 19:22 /tmp/output_labels.txt
-rw-rw-r-- 1 ubuntu ubuntu 84M Jul  7 19:22 /tmp/output_graph.pb
```

### Getting a training set 


# Cross compiling on rqaspberry pi3 

Download a release:

`wget https://github.com/meinside/libtensorflow.so-raspberrypi/releases/download/v1.2.0/libtensorflow_v1.2.0_20170619.tgz`

Install the toolchain

`sudo apt install sudo apt install gcc-arm-linux-gnueabihf`

Compile

`export CC=arm-linux-gnueabihf-gcc`
`CC=arm-linux-gnueabihf-gcc-5 GOOS=linux GOARCH=arm GOARM=6 CGO_ENABLED=1 go build  -o myprogram -ldflags="-extld=$CC"`
