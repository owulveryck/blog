---
author: Olivier Wulveryck
date: 2016-03-25T13:39:35+01:00
description: Some notes about Behaviour driver development, gherkin and Cucumber
draft: true
keywords:
- ruby
- BDD
- gherkin
- cucumber
title: How is Behaviour Driven Development is working with Gherkin and Cucumber
topics:
- BDD
type: post
---

# Introduction

Wikipedia defines the contract like this:

> A contract is a voluntary arrangement between two or more parties that is enforceable at law as a binding legal agreement.

If law usually describes what you can and cannot do, a contract is more likely to describe what's you are expected to do.

A law's goal is not only to give rules to follow, 
but also to maintain a stability in an ecosystem. 
In IT there are laws, that may be implicit, didactic, empiric, ... but the IT with all its laws should not 
dictate the expected behavior of the customer. But how often have you heard:

> "those computer stuffs are not for me"

> "It has always been this way"

There are laws that cannot be changed, but the contract between a customer and its provider could and should evolve.

In IT, like everywhere else where a customer/provider relationship exists, a special need is formalized via specifications.
Specifications ard hard to follow, but even more they're hard to evaluate.

<center>
![Babies (xkcd)](http://imgs.xkcd.com/comics/babies.png)
</center>

The __B__ehavior __D__riven __D__evelopment is 

# Setup

I will use the _Ruby_ implementation of cucumber.
To install it, assuming that we have a ` gem` installed, just run this command

```shell
# gem install cucumber
```

This will load all the required dependencies.
It may also be a good idea to use `bundle` if we plan to do further development of the steps in ruby.

### The test environment with bundler

```shell
> gem install bundler
> mkdir gherkin-test
> cd gherkin-test
> bundle init
Writing new Gemfile to /home/chronos/user/gherkin/Gemfile
```

### the _Gemfile_

Let's add the cucumber dependency in the Gemfile:

```shell
> cat Gemfile
source "https://rubygems.org"

gem 'cucumber'
```

and the _install_ the bundle:

```shell
> bundle install
Resolving dependencies...
Using builder 3.2.2
Using gherkin 3.2.0
Using cucumber-wire 0.0.1
Using diff-lcs 1.2.4
Using multi_json 1.7.9
Using multi_test 0.1.2
Using bundler 1.11.2
Using cucumber-core 1.4.0
Using cucumber 2.3.3
Bundle complete! 1 Gemfile dependency, 9 gems now installed.
Use `bundle show [gemname]` to see where a bundled gem is installed.
```

And now let's run cucumber within the bundle:

```shell
> bundle exec cucumber
No such file or directory - features. You can use `cucumber --init` to get started.
```

# Our first BDD

First, as requested by cucumber, let's initialize a couple of files in the directory to be "cucumber compliant".
Cucumber do have a helpful _init_ function. Let's run it now:

```shell
bundle exec cucumber --init
  create   features
  create   features/step_definitions
  create   features/support
  create   features/support/env.rb
```
