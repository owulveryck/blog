---
categories:
- Ideas
date: 2017-02-03T20:57:30+01:00
description: "A non geek article (for now) to talk about the dinner of the philosophers in the cloud"
draft: true
images:
- /2016/10/image.jpg
tags:
- tupleSpace
- topology
- linda
title: Linda, 31yo, with 5 starving philosophers...
---

It ain't no secret to anyone actually knowing me: I have always been fan of automation.

Nowadays the IT landscape has shifted to new disciplines such as "configuration management", "continuous integration", "self-deployment" and so on.
Those activities are driven by automation. Shell-scripting is not enough anymore. You have to learn new specific languages to deploy specific applications.
Cloudformation, Puppet, Ansible, Chef, Terraform, Cfengine... every single tool come along with its DSL and its idiomatic way of doing things.

The idea besides those tools is to reduce the need of human hand to deploy a new service.
Less human hand, less error, better time-to-market (and obviously less human cost.

Don't get me wrong: I think it is fantastic to rethink the IT, and those tools, for the one I know, are very good in doing that.
But that's still not enough.

The next step is, according to me (and to a bunch of colleagues), to give the design of an application to a computer program and to let it integrate, implement and monitor it.

Going straight from the architecture to the runtime. A true NoOps application management.


What I like in the engineering is the ability for a man to imagine and conceive tools.

> __The hand is the tool of tools__ - _Aristotle_.

![Philosophy at configuration manageent camp 2017](/assets/images/philosophy-cfgmgmtcamp2017.jpg)

# [Tuple Spaces (or, Good Ideas Don't Always Win)](https://software-carpentry.org/blog/2011/03/tuple-spaces-or-good-ideas-dont-always-win.html)

The title of this section is taken from [this blog post](https://software-carpentry.org/blog/2011/03/tuple-spaces-or-good-ideas-dont-always-win.html) which is indeed a good introduction on the tuple-space and how to use them.



# Meet Linda

Linda is a "coordination language" developed by Sudhir Ahuja at AT&T Bell Laboratories in collaboration with David Gelernter and Nicholas Carriero at Yale University in 1986 ([cf wikipedia](https://en.wikipedia.org/wiki/Linda_(coordination_language)))


## Back to the philosophers dinning problem

The philosophers dinnig problem is simply described in page 452 of the paper [Linda in context](http://www.inf.ed.ac.uk/teaching/courses/ppls/linda.pdf) from Nicholas Carriero and David Gelernter.



I have extracted the C-Linda implementation of this problem and copied it here. I have also copied the Go implementation I made to show the similarities.

#### The C linda implenentation
{{< highlight c >}}
Phil(i)
  int i;
{
    while(1) {
      think();
      in("room ticket");
      in("chopstick", i) ;
      in("chopstick", (i+l)%Num) ;
      eat();
      out("chopstick", i);
      out("chopstick", (i+i)%Num);
      out("room ticket");
    }
}
{{< /highlight >}}


{{< highlight c >}}
initialize()
{
  int i;
  for (i = 0; i < Hum; i++) C
    out("chopstick", i);
    eval(phil(i));
    if (i < (Num-1)) 
      out("room ticket");
  }
}
{{< /highlight >}}

#### The Go-linda implenentation

{{< highlight go >}}
for i := 0; i < num; i++ {
    ld.Out(chopstick(i))
    ld.Eval([]interface{}{phil, i})
    if i < (num - 1) {
        ld.Out(ticket{})
    }
}
{{< /highlight >}}


{{< highlight go >}}
func phil(i int) {
    p := philosopher{i}
    fmt.Printf("Philosopher %v is born\n", p.ID)
    for {
        p.think()
        fmt.Printf("[%v] is hungry\n", p.ID)
        ld.In(ticket{})
        ld.In(chopstick(i))
        ld.In(chopstick((i + 1) % num))
        p.eat()
        ld.Out(chopstick(i))
        ld.Out(chopstick((i + 1) % num))
        ld.Out(ticket{})
    }
}
{{< /highlight >}}
