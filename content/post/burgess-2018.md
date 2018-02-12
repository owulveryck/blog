+++
images = ["/2016/10/image.jpg"]
description = ""
categories = ["category"]
tags = ["tag1", "tag2"]
draft = true
title = "Burgess 2018"
date = 2018-02-12T20:49:15+01:00
+++

# A brief history of configuration (managing sprawl)

[The video](https://www.youtube.com/watch?v=Aorwdg2dRm0&t=653s)

<center>
{{< tweet 909118531362533376 >}}
</center>

At the very begining we were doing stuffs manually.
Then CFEngine is born; the problem was it had to address special needs because every student has special needs


There is a difference between what is happening quickly and what is happening slowly.
Things that are changing slowly are fairly stablem a strong base, a fundation.
Things that changes quickly are for special needs.

When doing configuration management, usually we don't take lessons from the past, and just consider that what has been done so far was wrong.

## How do we keep stuffs in order and how does that change over time?

For simplicity we try to normalize things; take the example of the database and the normalization.
When eveything is homogenous is it easy to manage.
But because of business needs and in order to add value, we must add specific stuffs.
This is when we add some tweaking and when things starts to diverge from the norm

We constantly face this conflict of interest: Keeping things "cheap" (make them fit into the database with a rigid schema) or going schemaless and adapt to a very special set of circonstances.

## How do we get from anarchy to order ?

### The illusion of scale

Let's take this example ![bidonville](https://en.wikipedia.org/wiki/Shanty_town#/media/File:Soweto_township.jpg)
This is a system that has been build manually. There is a sort of configuration management, but it is not managed automatically.

If you take a urbanized system, it looks a lot of simple becaue the number of degrees of freedom has been lowered.

But it is an illusion od scale.
Within the buildings, the "mess" still exists. The random wiring is hidden into the walls.
There has been a modularization that didn't made things easier, We've chosen to hide the complexity.
The whole point of configuration management is therefore not about reducing the complexity or making it disapear; it is about hidig it.


The same idea applies into procedural or functionnal programming: Hiding the complexity an algorithm within functions and get rid of the goto hell

### The obsession of tidiness

We think if we make something tidy (reducing something to a reasonable amount of function without spaghetti code), we have improved it.
But the rainforest are perfect example of sucessfull non-tidy systems. Complexity is eveywhere.
<center>
![Rainforest](https://static1.squarespace.com/static/578728c62994ca06a20bbeb7/t/5a6d47e99140b70ed9406822/1517111281462/b50973972515480367b5eb41139ce2af.jpg?format=2500w)
</center>
Those are the most sucessfull and the more ribusts systems of the planets. They are actually running for million years.

So why do we apply tidiness to our computers systems?
Because computer systems are explicitly under attack. In the rainforest, the actors are "just" trying to survive.
Humans are trying to tear down the systems.
Humans are constantly trying to kill each other.

There fore the modularization is also made in order to seperate elements and build walls around them to protect them.

_Consideration_ : By building wall and isolating, we make stuffs less robusts

What make the modularized system compelling is that we can define templates for creating things. We can make the system (look) more *deterministics*  by hidding complexity.
By reducing the degrees of freedom and absorbing the mess into infrastructure so we can't see it anymore, we can concentrate on the few degrees of freedom that stays on top and that are asociated  with business purpose.

This is the pattern that technology follows over time.

## Lifetime of technology - the stages of hubris



 
