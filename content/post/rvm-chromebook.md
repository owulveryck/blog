---
author: Olivier Wulveryck
date: 2016-03-31T10:23:02+02:00
description: How to setup RVM on an external drive on a chromebook
draft: true
keywords:
- ruby
- rvm
- chromebook
tags:
title: RVM on an USB stick on a chromebook
topics:
- topic 1
type: post
---

# Introduction

For testing purpose, I wanted to play with vagrant-aws and more generally with ruby on my chromebook.

Vagrant does not support _rubygems_ as installation method anymore ([see Mitchell Hashimoto's post](http://mitchellh.com/abandoning-rubygems))
and of course, there is no binary distribution available for the chromebook.

So I have to install it from the sources.

The [documentation](https://github.com/mitchellh/vagrant/wiki/Installing-Vagrant-from-Source) says:

* Do __NOT__ use the system Ruby - use a Ruby version manager like rvm, chruby, etc

Alright, anyway I don't want to mess with my system and break homebrew, so using RVM seems to be a good idea.

## Installing RVM

The RVM installation is relativly easy; simply running `curl -sSL https://get.rvm.io | bash` does the trick.
And then those commands make ruby 2.3.0 available via rvm:

```
$ source ~/.rvm/scripts/rvm  
$ rvm install 2.3.0
```

The supid trick here is that everything is installed in my $HOME directory, and as my chromebook is short on disk space: FS full !

Too bad.

## Using a USB stick

So my idea is to install the RVM suite onto a USB stick (because with me I don't have any SDHC card available).

### Preparing the stick

At first, the USB stick must be formatted in extendX (ext4) in order to be able to use symlinks, correct ownership etc.
I didn't find any way to format the device withing chromeos. Therefore I've used another Linux box to `mkfx.ext4` it.

__Note__: I've found that avoiding spaces in the volume name was a good idea; the command I've used to initialize the FS was
`mkfs.ext4 -L Lexar /dev/sdb1` (Lexar is the brand of the key, easy to remember).


Once connected on the chromebook, it's automatically mounted on `/media/removable/Lexar`.
The problem are the options: 

```shell
/dev/sda1 on /media/removable/Lexar type ext4 (rw,nosuid,nodev,noexec,relatime,dirsync,data=ordered)
```

the most problematic is `noexec` because I want to install executables in it.

So what I did was simply:

`sudo mount -o remount /dev/sda1 /media/removable/Lexar`

and that did the trick.

## Installing RVM on the USB

I will install rvm into `/media/removable/Lexar/rvm`. In order to avoid any ownership and permission problem I did:

```shell
mkdir /media/removable/Lexar/rvm
chown chronos:chronos /media/removable/Lexar/rvm
```

And then I created a simple `~/.rvmrc` file as indicated in the documentation with this:

```shell
$ cat ~/.rvmrc                                          
$ export rvm_path=/media/removable/Lexar/rvm
```

I also included this in my `~/.zshrc`

```shell
if [ -s "$HOME/.rvmrc"   ]; then
    source "$HOME/.rvmrc"
fi # to have $rvm_path defined if set
if [ -s "${rvm_path-$HOME/.rvm}/scripts/rvm"   ]; then
    source "${rvm_path-$HOME/.rvm}/scripts/rvm"
fi
```

## Installing rvm

the command I executed were then:

```
$ curl -sSL https://get.rvm.io | bash
$ source /media/removable/Lexar/rvm/scripts/rvm
$ rvm autolibs enable
$ rvm get stable
$ rvm install 2.3.0
```

And that did the trick

```
$ rvm list

rvm rubies

=* ruby-2.3.0 [ x84_64 ]

# => - current
# =* - current && default
#  * - default
```

## Testing with vagrant

### Cloning the vagrant sources

```shell
$ sudo mkdir /media/removable/Lexar/tools
$ sudo chown chronos:chronos /media/removable/Lexar/tools
$ cd /media/removable/Lexar/tools
$ git clone https://github.com/mitchellh/vagrant.git
```

### Preparing the rvm file for vagrant

To use the ruby 2.3.0 (that I've installed before) with vagrant, I need to create a rvmrc in the vagrant directory:

```
$ cd /media/removable/Lexar/tools/vagrant
$ rvm --rvmrc --create 2.3.0@vagrant
```

### Installing bundler

The bundler version that is supported by vagrant must be <= 1.5.2 as written in the `Gemfile`. So I'm installing version 
1.5.2

```shell
$ cd /media/removable/Lexar/tools/vagrant
$ source .rcmrv
$ gem install bundler -v 1.5.2
```
