---
author: Olivier Wulveryck
date: 2015-11-30T13:17:41Z
description: KSH getopts as a documentation of the script
draft: true
keywords:
- ksh
- shell
- getopts
- man
tags:
- ksh
- shell
- getopts
- man
title: KSH getopts, the unknown builtin documentation tool
type: post
---

[B.11. Using getopts](http://docstore.mik.ua/orelly/unix3/korn/appb_11.htm)

## Extract of the man page

The `man ksh93` does explain the `getopts` builtin and its associate options:

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
