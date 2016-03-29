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
title: Behaviour Driven Development with Gherkin and Cucumber (an introduction)
topics:
- BDD
type: post
---
#### Opening remarks

All my previous posts were about choreography, deployment, topology, and more recently about an attempt to include _AI_ in those systems.
This post is a bit apart, because I'm facing a new challenge in my work which is to implement BDD in a _CI_ chain. Therefore, I'm using
this blog as a reminder of what I did personally. The following of the _Markov_ saga will come again later.

# Introduction

Wikipedia defines the contract like this:

> A contract is a voluntary arrangement between two or more parties that is enforceable at law as a binding legal agreement.

If law usually describes what you can and cannot do, a contract is more likely to describe what's you are expected to do.

A law's goal is not only to give rules to follow, 
but also to maintain a stability in an ecosystem. 
In IT there are laws, that may be implicit, didactic, empiric, ... but the IT with all its laws should not 
dictate the expected behavior of the customer. But how often have you heard:

> "those computer stuffs are not for me, just get the thing done"

> "we've always done it this way"

There are laws that cannot be changed, but the contract between a customer and its provider could and should evolve.

In IT, like everywhere else where a customer/provider relationship exists, a special need is formalized via specifications.
Specifications are hard to follow, but even more they're hard to evaluate.

<center>
![Babies (xkcd)](http://imgs.xkcd.com/comics/babies.png)
</center>

The __B__ehavior __D__riven __D__evelopment is the assurance that everything have been made respectfully i
with the contract ²that has been established between the parties (customers and providers). 
To do things right, this contract should be established at the very beginning. 

Hence, every single item must be developed with all the _features_ of the contract in mind. And then, it should be
possible to use automation to perform the tests of behaviour, so that the developer can see if the contract is fulfilled, and if, for 
example, no regression has been introduced.

In a continuous integration chain, this is an essential piece that can be use to fully automate the process of delivery.

## Gherkin

To express the specification in a way that can be both human and comuter readable, the easiest way is to use a special dedicated
language. 

Such a language is known as [DSL](https://en.wikipedia.org/wiki/Domain-specific_language) ( Domain Specific Language). 

[Gherkin](https://github.com/cucumber/cucumber/wiki/Gherkin) is a DSL that _lets you describe software's behaviour without dealing how that behaviour
is implemented_

The behaviour is a scenario detailed as a set of _features_. A feature is a human readable english 
(or another human language among 37 implemented languages) text file with a bunch of key words in it (eg: __Given__, __And__, __When__, __Then__,...).
Those words do not only help the writer of the feature to organize its idea, but they are used by the Gherkin processor to localize the
test of the feature in the code. Of course, there is no magic in it: the test must have been implemented manually.

## And here comes Cucumber

The historic Gherkin processor is called Cucumber. It's a Ruby implementation of the Gherkin DSL.
Its purpose is to read a scenario, and to localize the Ruby code that is implementing the all the tests corresponding to the scenario.
Finally it executes the code, and for each feature it simply says ok or ko.

Easy.

Nowadays there are many implementation of Gherkin parser for different languages, but in this post I will stick to the Cucumber.

# Let's play

Let's see how we can implement a basic behaviour driver development with the help of cucumber and Ruby.


## The scenario

## The basic _features_

## Setting up the Ruby environment 

I will use the _Ruby_ implementation of cucumber.
To install it, assuming that we have a ` gem` installed, just run this command

```shell
# gem install cucumber
```

This will load all the required dependencies.
It may also be a good idea to use `bundle` if we plan to do further development of the steps in ruby.

#### The test environment with bundler

```shell
> gem install bundler
> mkdir gherkin-test
> cd gherkin-test
> bundle init
Writing new Gemfile to /home/chronos/user/gherkin/Gemfile
```

#### the _Gemfile_

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

### The implementation of the tests

First, as requested by cucumber, let's initialize a couple of files in the directory to be "cucumber compliant".
Cucumber do have a helpful _init_ function. Let's run it now:

```shell
bundle exec cucumber --init
  create   features
  create   features/step_definitions
  create   features/support
  create   features/support/env.rb
```
