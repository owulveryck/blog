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

# Machine learning and computer vision

Machine learning and tooling around it has increaded and gain in efficiency in the past years. it now "easy" to code a model that can be trained to detect and classify elements from a picture. 

Cloud providers are offering services that can instantly tag and label elements from an image. To achieve the goal of the CCTV, it would be really easy to use, for example, [AWS rekognition](https://aws.amazon.com/rekognition/), train the model, and post a request for each image seen.
This solution has a couple of problems:

* The network bandwidth: you need a network operational and a bandwidth big enough in upload to process the flow of images
* The cost: these services are cheap for thousand images, but consider about 10 fps to process (I don't even dream of 24fps), it is 864000 images a day and 26 million images a month... and considering that 1 million images are 1000 dollar...

I don't even talk about network latency because my CCTV would be pseudo-real-time and the ms of latency can be neglected.


"[Inception](https://research.google.com/pubs/pub43022.html)" is a deep convolutional neural network architecture used to classify images originally developped by Google.

# Geek

I am using the excellent blog post [How to Retrain Inception's Final Layer for New Categories](https://www.tensorflow.org/tutorials/image_retraining).

## Phase 1: recognizing usual people

To keep it simple, I've created a "class" people with the flowers classes. It means that I simple added a directory "people" to my "flowers" for now.

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


