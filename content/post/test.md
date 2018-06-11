---
title: "Recurrent Neural Network, Serverless with Webassembly and S3"
date: 2018-06-11T10:54:59+02:00
lastmod: 2018-06-11T10:54:59+02:00
draft: true
keywords: []
description: "This article is Bikeshedding!"
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
----
<link rel="stylesheet" type="text/css" href="/css/extra.css">
<script src="/js/wasm_exec.js"></script>
<script src="/js/loader.js"></script>
Download a "knowledge"

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

