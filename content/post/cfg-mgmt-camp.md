---
author: Olivier Wulveryck
date: 2016-02-01T09:42:15+01:00
description: Configuration management camp
draft: true
keywords:
- key
- words
tags:
- one
- two
title: cfg mgm camp
topics:
- topic 1
type: post
---


## To reach the next level we must *stop doing configuration management*

## from *scarcity* to abundance


The problem is now:

    * How do we choose and manage software *

## this is the age of *big software*

Actually: knoledge scarce...

## scarcity has shifted from code to *ops*

## the solution is *reusable, onpensource ops*

## encapsulation of a sofware

deb, rpm, ... but encapsulation requires a model

## the modeling language

the modeling language for applicatios
* model the software, not the machines
* model the software, not the configuration files

Introducing juju
...


https://xkcd.com/1319/

## beyond automation: reuse & sharing !
You should stop doing configuration management for software uniq to your organisation



# Valut
hashicorp product

## What is a secret 

Secret vs sensisitv

Secret:
* db credentials
* SSL CA / certificates
* Cloud access key
* wifi password
* source code

Sensitive :
* Phone numbers
* email addresses
* Datacenter location


## Why not config management

* No access control
* No auditing
* No revocation
* No key rolling


## Why not (onlinr) database?
* Not designed for secrets
* limited access controls
* typically plaintext storage
* no auditing or recovation abilities

## how to handle secret sprawl?
* secret material is distributed
* who has access?
* when were serets used?
* what is the attack surface ?
* what do we do in the event of compromise?
