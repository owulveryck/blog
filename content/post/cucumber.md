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

# Setup

I will use the _ruby_ implementation of cucumber.
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

we should add the cucumber dependency in the Gemfile:

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

And now we can run cucumber within the bundle:

```shell
> bundle exec cucumber
No such file or directory - features. You can use `cucumber --init` to get started.
```

# Our first BDD

First, as requested by cucumber, we need to initialize a couple of files in the directory to be "cucumber compliant".
Cucumber do have a helpful _init_ function. Let's run it now:

```shell
bundle exec cucumber --init
  create   features
  create   features/step_definitions
  create   features/support
  create   features/support/env.rb
```
