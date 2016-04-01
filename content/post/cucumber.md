---
author: Olivier Wulveryck
date: 2016-03-31T23:39:35+01:00
description: Some notes about Behaviour driver development, gherkin and Cucumber.
  The example describes here will test a service on an AWS's EC2 instance.
draft: true
keywords:
- EC2
- BDD
- Gherkin
- Cucumber
- Ruby
title: Behaviour Driven Development with Gherkin and Cucumber (an introduction)
topics:
- BDD
type: post
---
#### Opening remarks

All my previous posts were about choreography, deployment, topology, and more recently about an attempt to include _AI_ in those systems.
This post is a bit apart, because I'm facing a new challenge in my job which is to implement BDD in a _CI_ chain. Therefore, I'm using
this blog as a reminder of what I did personally. The following of the _Markov_ saga will come again later.

# Introduction

Wikipedia defines the word contract like this:

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

To express the specification in a way that can be both human and computer readable, the easiest way is to use a special dedicated
language. 

Such a language is known as [DSL](https://en.wikipedia.org/wiki/Domain-specific_language) ( Domain Specific Language). 

[Gherkin](https://github.com/cucumber/cucumber/wiki/Gherkin) is a DSL that _lets you describe software's behaviour without dealing how that behaviour
is implemented_

The behaviour is a scenario detailed as a set of _features_. A feature is a human readable English 
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
The idea here is not to test a Ruby development, but instead to use ruby to validate a shell script.
That's the main reason why I stick to Ruby (instead of GO which I know better). The Go implementation 
([GoDoc](https://github.com/DATA-DOG/godog), [GoConvey](https://github.com/smartystreets/goconvey), ...) relies 
on `go test` and therefore are related to a pure GO development. 
Of course I could do a complete GO development to encapsulate my scripts, but that's not the point; for my purpose, a scripting
language is a better choice.

Ruby is a scripting language and all the tests implemented here are neither dependent on the Ruby test framework nor on [RSpec](http://rspec.info/).

I will write a script that will deploy an EC2 instance via vagrant-aws and install an openvpn instance on it.

## The scenario

#### The customer point of view
With my role of customer, the feature I'm expecting is:

* Given the execution of the program, and waiting for it to be successful
* Then I may be able to watch netflix US from France.

The feature may be:

```gherkin
Feature: I want a program that
  will simply allows me to watch netflix US

  Scenario: I want to watch netflix
     Given I am on my chromebook
     And I have access to the shell
     When I want to watch netflix
     And I launch gonetflix.sh
     And it displays ready
     Then I open a navigator windows on http://www.netflix.com
     And I can watch Grey's anatomy (which is not available in france)
```

#### The architect point of view
As an architect the implementation I'm thinking of is

* start an EC2 instance (I will not create it in this post)
* register it to my DNS (with blog-test.owulveryck.info)
* install OpenVPN
* configure OpenVPN to make it accessible via blog-test.owulveryck.info 

#### The developer point of view
And as a developer, I'm thinking about using [vagrant-aws](https://github.com/mitchellh/vagrant-aws) to perform the tasks.

## The basic _feature_

I will describe here a single feature, just for testing purpose.

## Setting up the Ruby environment 

I will use the _Ruby_ implementation of cucumber.
To install it, assuming that we have a ` gem` installed, just run this command

```shell
# gem install cucumber
```

This will load all the required dependencies.
It may also be a good idea to use `bundle` if we plan to do further development of the steps in ruby.

#### The test environment with bundler

The whole development will run with the help of bundler (and RVM).
See this [post](http://dev.owulveryck.info/blog/2016/03/31/rvm-from-a-usb-stick-on-a-chromebook/) for more explanation on
how I set it up on my Chromebook.

```shell
> mkdir /media/removable/Lexar/tools/vpn-blog
> cd /media/removable/Lexar/tools/vpn-blog
> rvmrc --create 2.2.0@vpn-blog
> source .rvmrc
> gem install bundler -v 1.5.2 
> bundle init
Writing new Gemfile to /home/chronos/user/gherkin/Gemfile
```

#### the _Gemfile_

Let's add the cucumber, vagrant (as installed in a previous [post](http://dev.owulveryck.info/blog/2016/03/31/rvm-from-a-usb-stick-on-a-chromebook/) ),
and vagrant-aws dependencies in the Gemfile:

```shell
> cat Gemfile
source "https://rubygems.org"

gem "vagrant", :path => "/media/removable/Lexar/tools/vagrant"
gem "vagrant-aws"
gem "bundler", "1.5.2"
gem "cucumber"
```

and then _install_ the bundle:

```shell
> bundle _1.5.2_ install
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
...
Bundle complete! 1 Gemfile dependency, 9 gems now installed.
Use `bundle show [gemname]` to see where a bundled gem is installed.
```

And now let's run cucumber within the bundle:

```shell
> bundle _1.5.2_ exec cucumber
No such file or directory - features. You can use `cucumber --init` to get started.
```

### The skeleton of the test

First, as requested by cucumber, let's initialize a couple of files in the directory to be "cucumber compliant".
Cucumber do have a helpful _init_ function. Let's run it now:

```shell
bundle _1.5.2_ exec cucumber --init
  create   features
  create   features/step_definitions
  create   features/support
  create   features/support/env.rb
```

### The actual implementation of the scenario

What I need to do is to implement the scenario. Not the test scenario, the real one;
the one that will actually allows me to launch my ec2 instance, configure and start openvpn.

As I said before, I will use vagrant-aws to do so.

__Note__ vagrant was depending on _bsdtar_, and I've had to install it manually from source:

(`tar xzvf libarchive-3.1.2.tar.gz && ... && ./configure --prefix=/usr/local && make install clean`)

#### Installling vagrant-aws plugin

The vagrant-aws plugin has been installed by the bundler because I've indicated it as a dependency in the Gemfile.
But, I will have to have it as a requirement in the Vagrantfile because I'm not using the "official vagrant" and that
I am running in a bundler environment:

> Vagrant's built-in bundler management mechanism is disabled because
> Vagrant is running in an external bundler environment. In these
> cases, plugin management does not work with Vagrant. To install
> plugins, use your own Gemfile. To load plugins, either put the
> plugins in the `plugins` group in your Gemfile or manually require
> them in a Vagrantfile.

#### Installing the base box 

The documentation says that the quickest way to get started is to install the dummy box. 
That's what I did:

```shell
$ bundle _1.5.2_ exec vagrant box add dummy https://github.com/mitchellh/vagrant-aws/raw/master/dummy.box
...
==> box: Successfully added box 'dummy' (v0) for 'aws'!
```

#### The Vragrantfile

The initial vagrantfile looks like this:

```ruby
require "vagrant-aws"
Vagrant.configure("2") do |config|
  config.vm.box = "dummy"

  config.vm.provider :aws do |aws, override|
    aws.access_key_id = "YOUR KEY"
    aws.secret_access_key = "YOUR SECRET KEY"
    aws.session_token = "SESSION TOKEN"
    aws.keypair_name = "KEYPAIR NAME"

    aws.ami = "ami-7747d01e"

    override.ssh.username = "ubuntu"
    override.ssh.private_key_path = "PATH TO YOUR PRIVATE KEY"
  end
end
```