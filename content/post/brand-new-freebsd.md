---
author: Olivier Wulveryck
date: 2016-01-24T15:25:54+01:00
description: The setup of my new Freeebsd
draft: true
keywords:
- FreeBSD
tags:
title: Setting up my new BSD box
topics:
- topic 1
type: post
---
# About the server

I have subscribed for a now dedicated BSD box.
The provider is OVH, and this box is hosted in Canada.

It's been a few years since I first subscribed a box a OVH, and I really enjoy their service.

The main usage for this box is a `geek box`. I use it to exeperiment some stuffs, I cannot experiment into my home PC.
Therefore I usually create a bunch of `jails`, and each jail is dedicated to a task

Here are the informations of my new box:

```shell
~ uname -a
FreeBSD localhost 10.2-RELEASE-p9 FreeBSD 10.2-RELEASE-p9 #0: Thu Jan 14 01:32:46 UTC 2016
root@amd64-builder.daemonology.net:/usr/obj/usr/src/sys/GENERIC  amd64
```

## ZFS

My root is a ZFS pool named `zroot`

# Basic installation

This box is a `10.2 release` therefore it uses the "new" `pkg` tool instead of the legacy `pkg_*` tools.

## Configuration of `pkg`


## ZSH

### Installation
```shell
pkg install zsh
```
### Oh-my-zsh

Install `git` 

```shell
pkg install git
```

```shell
sh -c "$(curl -fsSL https://raw.github.com/robbyrussell/oh-my-zsh/master/tools/install.sh)"
```

# Setting up openvpn

## Installation
```shell
pkg update
pkg install openvpn
```

## Configuration
```shell
mkdir /usr/local/etc/openvpn
cp /usr/local/share/examples/openvpn/sample-config-files/server.conf /usr/local/etc/openvpn/server.conf 
```

The default options are used, and only user and groups ares set to nobody for security reasons

```shell
...
# It's a good idea to reduce the OpenVPN
# daemon's privileges after initialization.
#
# You can uncomment this out on
# non-Windows systems.
user nobody
group nobody
...
```

### Generating the keys

```shell
cp -r /usr/local/share/easy-rsa /usr/local/etc/openvpn/easy-rsa
```

### Generating a client certificate

I will generate an openvpn configuration for my chromebook

```shell
cd /usr/local/etc/openvpn/easy-rsa/
./build-key chromebook
```

# Rescue...

Of course, I forgot one rue in my pf.conf and therefore I could not access to my box anymore

## Maganer
Boot into rescue mode

```shell
rescue-bsd# zpool import zroot
rescue-bsd# zpool list
internal error: failed to initialize ZFS library
```

That's because I did import the zroot into /

```shell
zpool import -o altroot=/mnt zroot
```
