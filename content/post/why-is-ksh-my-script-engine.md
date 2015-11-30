---
author: Olivier Wulveryck
date: 2015-11-30T13:17:41Z
description: Two reasons why I usually use KSH93 as my script engine
draft: true
keywords:
- ksh
- shell
- getopts
- man
tags:
- ksh
- bash
- shell
- getopts
- man
title: KSH93 cool features for scripting
type: post
---

From time to time, I'm envolved into a tolling conversation when any linux kiddie tells me:

_Bash is really the superior shell_ 

I totally disagree, but as I'm getting older, I don't argue anymore.

Anyway, I this post I will expose two arguments, or I should say two reasons, why I usually use `ksh93` to run my scripts.

**Note** I'm really talking about the engine of the script, (the shebang definition). 
My `$SHELL` for interactivity is `zsh` because it's efficient enough
and it has a bunch of really cool features I won't discuss in this post (maybe later)

# Read, loops, forks and efficiency...



```shell
$ cat test                                                                                                         
for i in $(seq 1 500)
do
    echo $i | read a
    echo -ne "$a\r"
done
echo ""
```

```shell
for i in bash zsh ksh                                                                                         
do
    echo "$i =>"
    eval time $i test
done
```

```shell
bash =>

bash test  0,17s user 0,86s system 95% cpu 1,079 total
zsh =>
500
zsh test  0,08s user 0,46s system 82% cpu 0,648 total
ksh =>
500
ksh test  0,07s user 0,46s system 65% cpu 0,819 total
```



On voit que bash ne donne pas le resultat escompté, la commande read a ingurgité la sortie standard de exec, mais comme le shell a vu un pipe, il a bêtement forké... et donc, en retournant au shell courant, on a perdu notre $a....

ksh et zsh sont un peu plus malin, ils ont vu que l'appel après le pipe était une fonction "builtin", donc ils n'ont pas forké... Ce qui, en plus de donner le résultat attentu, fait gagner du temps et des ressources...

```shell
$ for i in bash zsh ksh                                                                                         
do
    echo "$i =>"
    strace -c  $i  test 2>&1 | egrep "clone|calls"
done
```



```shell
bash =>
% time     seconds  usecs/call     calls    errors syscall
56.05    0.067081          67      1001           clone
zsh =>
% time     seconds  usecs/call     calls    errors syscall
71.57    0.057681         115       501           clone
ksh =>
% time     seconds  usecs/call     calls    errors syscall
68.50    0.042059          84       500           clone
```

# The KSH93 Getopts unknown feature

A few month back, I wanted to use the `getopts` builtin in a script. As usual, I did _RTFM_.
Here is the extract of the man page of ksh93 relative to the getopts function:


<pre>
<B>getopts</B> [ <B>-a</B> <I>name</I> ] <I>optstring vname</I> [ <I>arg</I> ... ]

Checks <I>arg</I> for legal options.  If <I>arg</I> is omitted, the positional parameters are used.

An option argument begins with a <B>+</B> or a <B>-</B>.  An option not beginning with <B>+</B> or <B>-</B> or the argument <B>--</B> ends the options.
Options beginning with <B>+</B> are only recognized when <I>optstring</I> begins with a <B>+</B>.

<I>optstring</I> contains the letters that <B>getopts</B> recognizes.
If a letter is followed by a <B>:</B>, that option is expected to have an argument.
The options can be separated from the argument by blanks.
The option <B>-?</B> causes <B>getopts</B> to generate a usage message on standard error.
The <B>-a</B> argument can be used to specify the name to use for the usage message, which defaults to <B>$0</B>.

<B>getopts</B> places the next option letter it finds inside variable <I>vname</I> each time it is invoked.
The option letter will be prepended with a <B>+</B> when <I>arg</I> begins with a <B>+</B>.
The index of the next <I>arg</I> is stored in <FONT SIZE="-1"><B>OPTIND</B>.
</FONT> The option argument, if any, gets stored in <FONT SIZE="-1"><B>OPTARG</B>.  </FONT> 

A leading <B>:</B> in <I>optstring</I> causes <B>getopts</B> to store the letter of an invalid option in <FONT SIZE="-1"><B>OPTARG</B>, </FONT>
and to set <I>vname</I> to <B>?</B> for an unknown option and to <B>:</B> when a required option argument is missing.
Otherwise, <B>getopts</B> prints an error message.
The exit status is non-zero when there are no more options.

<P> There is no way to specify any of the options <B>:</B>, <B>+</B>, <B>-</B>, <B>?</B>, <B>[</B>, and <B>]</B>.

The option <B>#</B> can only be specified as the first option. 
</pre>

This particular sentence, in the middle of the documentation peaked my interest

> The option -? causes getopts to generate a usage message on standard error.

What? We can generate usage with getopts? 

Cool, any script should be documented, but any documentation should not be difficult to implement.

<img class="img-responsive center-block" src="http://imgs.xkcd.com/comics/manuals.png"> 
<center>_https://xkcd.com/1343/_</center>


I did googled and found this 
[web page](http://docstore.mik.ua/orelly/unix3/korn/appb_11.htm) which is an extract from this book [Learning the Korn Shell](http://shop.oreilly.com/product/9780596001957.do)

Sounds cool, let's see how it works in the real life
## An example

### The script
```shell
#!/bin/ksh

ENV=dev
MPATH=/tmp
##
### Man usage and co...

USAGE="[-?The example script v1.0]"
USAGE+="[-author?Olivier Wulveryck]"
USAGE+="[-copyright?Copyright (C) My Blog]"
USAGE+="[+NAME?$0 --- The Example Script]"
USAGE+="[+DESCRIPTION?The description of the script]"
USAGE+="[u:user]:[user to run the command as:=$USER?Use the name of the user you want to sudo to: ]"
USAGE+="[e:env]:[environnement:=$ENV?environnement to use (eg: dev, prod) ]"
USAGE+="[p:path]:[Execution PATH:=$MPATH?prefix of the chroot]"
USAGE+="[+EXAMPLE?$0 action2]"
USAGE+='[+SEE ALSO?My Blog Post: http://blog.owulveryck.info/2015/11/30/ksh93-cool-features-for-scripting]'
USAGE+="[+BUGS?A few, maybe...]"

### Option Checking

while getopts "$USAGE" optchar ; do
    case $optchar in
        u)  USER=$OPTARG
        ;;
        e)  ENV=$OPTARG
        ;;
        p)  PATH=$OPTARG
        ;;
    esac
done
shift OPTIND-1
ACTION=$1
```
### The invocation

Here are two _singing_ examples of the usage output (sorry, I'm tired)
#### _Ballad of a thin man_
```shell
$ ./manheader.ksh --man
NAME
  ./manheader.ksh --- The Example Script

SYNOPSIS
  ./manheader.ksh [ options ]

DESCRIPTION
  The description of the script

OPTIONS
  -u, --user=user to run the command as
                  Use the name of the user you want to sudo to: The default value is owulveryck.
  -e, --env=environnement
                  environnement to use (eg: dev, prod) The default value is dev.
  -p, --path=Execution PATH
                  prefix of the chroot The default value is /tmp.

EXAMPLE
  ./manheader.ksh action2

SEE ALSO
  My Blog Post: http://blog.owulveryck.info/2015/11/30/ksh93-cool-features-for-scripting

BUGS
  A few, maybe...

IMPLEMENTATION
  version         The example script v1.0
  author          Olivier Wulveryck
  copyright       Copyright (C) My Blog
```

#### I'm gonna try _with a little help (from my friends)_

```shell
$ ./manheader.ksh --help
Usage: ./manheader.ksh [ options ]
OPTIONS
  -u, --user=user to run the command as
                  Use the name of the user you want to sudo to: The default value is owulveryck.
  -e, --env=environnement
                  environnement to use (eg: dev, prod) The default value is dev.
  -p, --path=Execution PATH
                  prefix of the chroot The default value is /tmp.
```

And let's try with an invalid option...

```shell
  ./manheader.ksh -t
./manheader.ksh: -t: unknown option
Usage: ./manheader.ksh [-u user to run the command as] [-e environnement] [-p Execution PATH]
```
