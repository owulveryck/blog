---
author: Olivier Wulveryck
date: 2015-11-20T10:09:30Z
description: A tosca lifecycle represented as a digraph
draft: true
tags:
- TOSCA
- Digraph
- Graph Theory
- golang
title: TOSCA lifecycle as a digraph
topics:
- TOSCA
type: post
---

# About TOSCA

The [TOSCA](https://www.oasis-open.org/committees/tc_home.php?wg_abbrev=tosca) acronym stands for 
_Topology and Orchestration Specification for Cloud Applications_. It's an [OASIS](https://www.oasis-open.org) standard.

The purpose of the TOSCA project is to represent an application by its topology and formalize it using the TOSCA grammar.

The [[TOSCA-Simple-Profile-YAML-v1.0]](http://docs.oasis-open.org/tosca/TOSCA-Simple-Profile-YAML/v1.0/csprd01/TOSCA-Simp$le-Profile-YAML-v1.0-csprd01.html) 
current specification in YAML introduces the following concepts.

> - TOSCA YAML service template: A YAML document artifact containing a (TOSCA) service template that represents a Cloud application.
> - TOSCA processor: An engine or tool that is capable of parsing and interpreting a TOSCA YAML service template for a particular purpose. For example, the purpose could be validation, translation or visual rendering.
> - TOSCA orchestrator (also called orchestration engine): A TOSCA processor that interprets a TOSCA YAML service template then instantiates and deploys the described application in a Cloud.
> - TOSCA generator: A tool that generates a TOSCA YAML service template. An example of generator is a modeling tool capable of generating or editing a TOSCA YAML service template (often such a tool would also be a TOSCA processor).
> - TOSCA archive (or TOSCA Cloud Service Archive, or “CSAR”): a package artifact that contains a TOSCA YAML service template and other artifacts usable by a TOSCA orchestrator to deploy an application.

## My work with TOSCA

I do believe that TOSCA may be a very good leverage to port a "legacy application" (aka _born in the datacenter_ application) into a cloud ready application without rewriting it completely to be cloud compliant.
To be clear, It may act on the hosting and execution plan of the application, and not on the application itself.

As an example, the very famous ELK suite may be described in a TOSCA way as written [here](http://docs.oasis-open.org/tosca/TOSCA-Simple-Profile-YAML/v1.0/csprd01/TOSCA-Simple-Profile-YAML-v1.0-csprd01.html#USE_CASE_MULTI_TIER_1).

<img class="img-square img-responsive" src="http://docs.oasis-open.org/tosca/TOSCA-Simple-Profile-YAML/v1.0/csprd01/TOSCA-Simple-Profile-YAML-v1.0-csprd01_files/image037.jpg" alt="ELK representation"/>

While I was learnig GO, I have developped a [TOSCA lib](https://github.com/owulveryck/toscalib) and a [TOSCA processor](https://github.com/owulveryck/toscaviewer) which are, by far, not _idiomatic GO_.

Here are two screenshots of the rendering in a web page made with my tool (and the graphviz product):

<hr/>
*The graph representation of a _Single instance wordpress_*
<img class="img-responsive" src="/assets/images/toscaviewer_template_def.png" alt="Tosca view ofthe single instance wordpress"/>


*The graph representation of a lifecycle of _Single instance wordpress_*
<img class="img-responsive" src="/assets/images/toscaviewer_licecycle_def.png" alt="Lifecycle representation of the single wordpress instance representation"/>
<hr/>

The TOSCA file is parsed with the help of the `TOSCALIB` and then it may fill two adjacency matrix.

- The first one will represent the nodes
- The second one will be focused on the lifecycles of the nodes.

The [graphviz](http://graphviz.org) take care of the (di)graph representation.

What I would like to do now, is a little bit more: I would like to play with the graph and query it
Then I should perform requests on this graph. For example I could ask:

* _What are the steps to go from the state Initial of the application, to the state running_ ?
* _What are the steps to go from stop to delete_
* ...

and that would be **the premise of a TOSCA orchestrator**.

## The digraph go code

I've recently discoverd the [digraph](https://github.com/golang/tools/tree/master/cmd/digraph) tool, that I will use for querying the graphs.
I need to dig a little bit in the code to see how the graph is represented.

# Let's go 

I consider that the yaml file is read and that a [ToscaDefinition](http://godoc.org/github.com/owulveryck/toscalib/#ToscaDefinition) structure has been filled.

```golang
...
var toscaTemplate toscalib.ToscaDefinition
...
err = toscaTemplate.Parse(file)
...
```

