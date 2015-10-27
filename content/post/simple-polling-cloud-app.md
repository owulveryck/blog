+++
date = "2015-10-23T09:54:27+01:00"
draft = true
title = "Simple polling app, a cloud native part 1"

+++

In this series of posts I'll explain how to setup a simple polling app, the cloud way.
This app, written in go, will be hosted on a PAAS, and I've chosen the [Google App Engine](https://cloud.google.com/appengine/docs) for convenience.

I will not explain in this post how to setup the Development environment as it is described [here](https://cloud.google.com/appengine/docs/go/gettingstarted/devenvironment)

# The principles of the application

The application is composed of two web pages that will display :

* the question "will you participate" and a form input where you will be able to write your name and three buttons "yes", "no" and ""maybe"".
```
Will you participate 
+------------------+  +-----+ +-----+ +-------+
|  Your name       |  | YES | |  NO | | Maybe |
+------------------+  +-----+ +-----+ +-------+
```

* a simple table with two columns:

* One will hold the participant name
* The other one will display its response

```
+---------------------+-------+
|  John doe           | YES   |
+---------------------+-------+
|  Johnny Vacances    | NO    |
+---------------------+-------+
|  Foo Bar            | YES   |
+---------------------+-------+
|  Toto Titi          | NO    |
+---------------------+-------+
|  Pascal Obistro     | YES   |
+---------------------+-------+
```

# Setting up the development environment

First, we will create a directory that will host the sources of our application in our `GOPATH/src`.
_Note_: For convenience I've created a github repo named "google-app-example" to host the complete source.

```
~ mkdir -p $GOPATH/src/github.com/owulveryck/google-app-example
~ cd $GOPATH/src/github.com/owulveryck/google-app-example
~ git init
~ git remote add origin https://github.com/owulveryck/google-app-example
```


## Hello World!

Let's create the hello world first to validate the whole development chain.
As written in the doc, create the two files `hello.go` and `app.yaml`.
Obviously the `simple-polling.go` file will hold the code of the application. Let's focus a bit on the _app.yaml_ file.
The documentation of the _app.yaml_ file is [here](https://cloud.google.com/appengine/docs/go/config/appconfig). The goal of this file is to specifiy the runtime configuration of the engine.
This simple file replace the "integration" task for an application typed "born in the datacenter"

Here is my app.yaml
```yaml
application: simple-polling
version: 1
runtime: go
api_version: go1

handlers:
- url: /.*
  script: _go_app
```

And then we try our application with the command `goapp serve $GOPATH/src/github.com/owulveryck/google-app-example/`
which should display something similar to:
```
INFO     2015-10-26 21:02:10,295 devappserver2.py:763] Skipping SDK update check.
INFO     2015-10-26 21:02:10,468 api_server.py:205] Starting API server at: http://localhost:52457
INFO     2015-10-26 21:02:12,011 dispatcher.py:197] Starting module "default" running at: http://localhost:8080
INFO     2015-10-26 21:02:12,014 admin_server.py:116] Starting admin server at: http://localhost:8000
```

Then I can open my browser (or curl) and point it to http://localhost:8080 to see my brand new "Hello World!" displayed

```
~ curl http://localhost:8080
Hello, world!
```

# The web templates

