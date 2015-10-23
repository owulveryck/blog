+++
date = "2015-10-23T09:54:27+01:00"
draft = true
title = "Simple polling app, cloud native"

+++

In this post I explain how to setup a simple polling app, the cloud way.
This app, written in go, will be hosted on a PAAS, and I've chosen the [Google App Engine](https://cloud.google.com/appengine/docs) for convenience.

I will not explain in this post how to setup the Development environment as it is described [here](https://cloud.google.com/appengine/docs/go/gettingstarted/devenvironment)

# The principles of the application

The application is a single page that will display a header with a question "will you participate" and a form input where you will be able to write your name and three buttons "yes", "no" and ""maybe"".
The body of the page will contain a simple table with two columns:

* One will hold the participant name
* The other one will display its response

```
Will you participate 
+------------------+  +-----+ +-----+ +-------+
|  Your name       |  | YES | |  NO | | Maybe |
+------------------+  +-----+ +-----+ +-------+

-----------------------------------------------

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



