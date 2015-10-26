+++
date = "2015-10-26T10:41:57Z"
draft = false
title = "Developping \"Google Apps\" on my Chromebook"

+++

It is a  week now that I'm playing with my chromebook.
I really enjoy this little internet Terminal.

I "geeked" it a little bit and I installed my favorites dev tools eg:

* [The solarized theme for the terminal](https://gist.github.com/johnbender/5018685)
* `zsh` with [Oh-my-zsh](https://github.com/robbyrussell/oh-my-zsh)
* `tmux` (stared with `tmux -2` to get 254 colors)
* `git`
* `vim`
* a `Go` compiler
* The [`HUGO`](http://gohugo.io/overview/quickstart/) tools to write this blog.


All of it has been installed thanks to the "brew" package manager and following [those instructions](https://github.com/Homebrew/linuxbrew/wiki/Chromebook-Install-Instructions).

## Google Development Environment

I've installed the Google Development Environement as described [here](https://cloud.google.com/appengine/docs/go/gettingstarted/devenvironment).

Python 2.7 is a requirements so I `brewed it` without any noticeable issue.

When I wanted to serve locally my very first Google App developement, I ran into the following error:

```
~ go app serve $GOPATH/src/myapp
...
ImportError: No module named _sqlite3
error while running dev_appserver.py: exit status 1
```

Too bad. I've read that this module should be built with python, but a even a `find /` (I know it's evil) didn't return me any occurence.

So, I have:

* Googled 
* reinstalled sqlite with `brew reinstall sqlite`
* reinstalled python with `brew reinstall python`
* played with brew link, unlink and so
* ...

Still no luck!

I've also tried the compilation with a `verbose` option, and I the log file, there is an explicit message:

```
Python build finished, but the necessary bits to build these modules were not found:
_bsddb  _sqlite3_tkinter
...
To find the necessary bits, look in setup.py in detect_modules() for the modules name.
```


That's where I am now, stuck with a stupid python error. I'd like the folks at google to provide a pure go developement enrironement that would avoid the bootstraping problems.

I'll post an update as soon as I have solved this issue !

*EDIT*:

I've had a look in the `setup.py` file. To compile the sqlite extension, it looks into the following paths:

```
...
sqlite_incdir = sqlite_libdir = None
sqlite_inc_paths = [ '/usr/include',
                     '/usr/include/sqlite',
                     '/usr/include/sqlite3',
                     '/usr/local/include',
                     '/usr/local/include/sqlite',
                     '/usr/local/include/sqlite3',
                   ]
...
```

But in my configuration, the libraries are present in `/usr/local/linuxbrew/*`. Hence, simply linking the include and libs dd the trick
