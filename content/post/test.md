---
title: "Recurrent Neural Network, Serverless with Webassembly and S3"
date: 2018-06-11T10:54:59+02:00
lastmod: 2018-06-11T10:54:59+02:00
draft: true
keywords: []
description: "This article is Bikeshedding! It is about creating a neural network runtime environment and running it in the browser via Wasm in #Golang. It also demonstrates the strict separation of the Neural Net dev kit, the Neural Net runtime and the knowledge (software 2.0)"
tags: []
categories: []
author: "Olivier Wulveryck"

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
----
<link rel="stylesheet" type="text/css" href="/css/extra.css">
<script src="/js/wasm_exec.js"></script>
<script src="/js/loader.js"></script>

During the past weeks, I've had the opportunity to play a bit with Wasm and Go.

All those experiments led me to a write a proof of concepts that can illustrate everything I have said recently about:

* Thinking the deep-learning stack like an Ops (see my post about [NNRE/NNDK](https://blog.owulveryck.info/2018/04/16/considerations-about-software-2.0.html)).
* Capturing the real value of the training process (the knowledge) into a sequence of bits (the lightning talk I gave about it at the [dotAI](https://www.dotai.io/) should be online soon).


# Greetings professor Falken... Shall we play a game?

For a demo, I have developed a simple LSTM that can play the tic-tac-toe game.
I am a fan of this example for AI-related kinds of stuff, it is indeed a "[Madeleine](https://en.wikipedia.org/wiki/Madeleine_(cake)#Literary_reference)" for me.

I will not go into every single detail here, and you can find the code [here](https://github.com/owulveryck/rnnttt/tree/blog).
But the principle is as follow:

An autonomous code based on an LSTM is learning the possible combinations of a winning tic-tac-toe board for token X.

Then, the weights of the LSTM are exported (as a Gob file).

Another code (the NNRE) car read the knowledge file, applies it to the LSTM model and implements the mechanism to play the game.
The interactive part is concurrent and synchronized via channels (did I told you how much I like this model of concurrency :))
That's almost it.

## Wasm ?

This was my first POC. Then I realized that I could run all of this in the browser.

I first created a table and gave every cell an ID: `ttt0 ... ttt8`.

Within my GO/Wasm code, I added an EventListener to trigger a callback when a click on a cell is made:

{{< highlight go >}}
for i, v := range []string{"ttt0", "ttt1", "ttt2", "ttt3", "ttt4", "ttt5", "ttt6", "ttt7", "ttt8"} {
      m := mycb{ v, i, p }
      js.Global.Get("document").Call("getElementById", v).Call("addEventListener", "click", js.NewCallback(m.cb))
}
{{</ highlight >}}


_Note_ `mycb` is a just a wrapper to pass some parameters to the "cb" method which is actually the callback.

When the AI player has made a move, an event is also triggered, and the corresponding letter is placed in the correct cell of the table.

## The knowledge

I really wanted to show that the knowledge was strictly separated from the code.
So I made a little hack to be able to import the knowledge at runtime.
The uploaded file is 

# Go and test it live!

Download a "knowledge" (Thos have been pre-trained with different hyper parameters).

* [Knowldege 1](/tictactoe/tictactoe1.bin)
* [Knowldege 2](/tictactoe/tictactoe2.bin)
* [Knowldege 3](/tictactoe/tictactoe3.bin)

Upload it here: <input type="file" id="knowledgeFile" multiple size="1" style="width:250px" accept=".bin">

Load the WASM file (the file is 25Mo): <button onClick="load();" id="loadButton" style="width:125px;">Load</button>

Wait for the file to be compiled (the run button will become available): <button onClick="run();" id="runButton" style="width:125px;" disabled>Run</button>


<center>
<table style="border:1px solid black;">
  <tr style="height: 50px; border:1px solid black;">
    <td style="text-align: center; vtext-align: middle; width: 50px; border:1px solid black;" id="ttt0"></td>
    <td style="text-align: center; vtext-align: middle; width: 50px; border:1px solid black;" id="ttt1"></td>
    <td style="text-align: center; vtext-align: middle; width: 50px; border:1px solid black;" id="ttt2"></td>
  </tr>
  <tr style="height: 50px; border:1px solid black;">
    <td style="text-align: center; vtext-align: middle; width: 50px; border:1px solid black;" id="ttt3"></td>
    <td style="text-align: center; vtext-align: middle; width: 50px; border:1px solid black;" id="ttt4"></td>
    <td style="text-align: center; vtext-align: middle; width: 50px; border:1px solid black;" id="ttt5"></td>
  </tr>
  <tr style="height: 50px; border:1px solid black;">
    <td style="text-align: center; vtext-align: middle; width: 50px; border:1px solid black;" id="ttt6"></td>
    <td style="text-align: center; vtext-align: middle; width: 50px; border:1px solid black;" id="ttt7"></td>
    <td style="text-align: center; vtext-align: middle; width: 50px; border:1px solid black;" id="ttt8"></td>
  </tr>
</table>
</center>

