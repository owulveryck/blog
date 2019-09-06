---
title: "Test and validate the IaC"
date: 2019-09-06T09:41:04+02:00
lastmod: 2019-09-06T09:41:04+02:00
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

# about

### Quick reminder about terraform

### The benefit of TDD

### About terratest

# The goal of the module

I am writing this module to fit a need I have when I develop something related to machine learning.

## The usecase - "_eat your own dog food_"

All the material I need to work as a developer is installed in user space in my home directory.

**Editor:** I use neovim with a couple of plugins, and as its [doc says](https://neovim.io/): neovim _works the same everywhere: one build-type, one command_).

**Python:** I use Anaconda to manage my python environments; therefore, everything is installed in a userspace in my home directory.

**Go:** my workspace (the `GOPATH`) and my go installation (the `GOROOT`) are all installed in my `$HOME` as well.

During the life cycle of the development, I need different computation power. Most of the time, I need a basic CPU to run my editor, the linter and some basic tests.

### Vision statement
This paragraph aim to find a vision statement for the module. The vision statement will guide us toward the process of design an development.

> For the developer who needs elastic CPU/GPU power, the Terraform module will facilitate the creation of a development environment in the cloud,
> and unlike other solution, any change in the environment has no impact on the work and settings of the user.

# Architecture of the module



### Input variables

### Output

# Conclusion
