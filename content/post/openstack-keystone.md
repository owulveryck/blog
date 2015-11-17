---
author: Olivier Wulveryck
date: 2015-11-17T10:05:42Z
description: Playing with openstack keystone
draft: true
keywords:
- openstack
- keystone
- authentication
tags:
- openstack
- keystone
- authentication
- REST
title: Playing with (Openstack) Keystone
topics:
- topic 1
type: post
---

In the cloud computing, alongside of the hosting monsters such as amazon or google, there is the [Openstack Platform](https://www.openstack.org).

Openstack is not a single software, it is more a galaxy of components aim to control the infrastructure, such as hardware pools, storage, network.
The management can then be done via a Web based interface or via a bunch of RESTful API.

I would like to evaluate its identity service named [keystone](http://docs.openstack.org/developer/keystone/) and use it as a AuthN and AuthZ backend for my simple_iaas example.

_Note_ : I will consider that the openstack keystone is installed. As I don't want to rewrite an installation procedure as many exists already on the web.

# My goal

My goal is to have a webservice that will protect the scopes of my IAAS. 
I may declare two users:

- One may list the nodes via a GET request
- The other one may also create and destroy nodes via POST and DELETE request

# Let's go 

I won't use any external web server. Instead I will rely on the builtin Eventlet based web server.

The documentation says it is deprecated, indeed I will use it for testing purpose, so that will do the job.

## The WSGI pipeline configuration

To be honest, I don't know anything about the python ecosystem. And as it is my blog, I will write anything I've learned from this experience... 

So:

- WSGI is a gateway interface for python, and my understanding is that it's like the good old CGI we used in the beginning of this century;
- Is is configured by a ini file based on [Paste](http://pythonpaste.org/) and especially _Paste Deploy_ which is a system made for loading and configuring WSGI components.

The WSGI interface is configured by a ini file as written in the [Openstack keystone documentation](http://docs.openstack.org/developer/keystone/configuration.html).
This file is called `keystone-paste.ini`. I won't touch it and use the provided one. It sounds ok and when I start the service with `keystone-all` I can see in the logs:

```logs
2015-11-17 10:05:04.918 7068 INFO oslo_service.service [-] Starting 2 workers
2015-11-17 10:05:04.920 7068 INFO oslo_service.service [-] Started child 7082
2015-11-17 10:05:04.922 7068 INFO oslo_service.service [-] Started child 7083
2015-11-17 10:05:04.925 7082 INFO eventlet.wsgi.server [-] (7082) wsgi starting up on http://0.0.0.0:35357/
2015-11-17 10:05:04.927 7068 INFO keystone.common.environment.eventlet_server [-] Starting /usr/bin/keystone-all on 0.0.0.0:5000
2015-11-17 10:05:04.927 7068 INFO oslo_service.service [-] Starting 2 workers
2015-11-17 10:05:04.930 7068 INFO oslo_service.service [-] Started child 7084
2015-11-17 10:05:04.934 7083 INFO eventlet.wsgi.server [-] (7083) wsgi starting up on http://0.0.0.0:35357/
2015-11-17 10:05:04.936 7068 INFO oslo_service.service [-] Started child 7085
2015-11-17 10:05:04.940 7085 INFO eventlet.wsgi.server [-] (7085) wsgi starting up on http://0.0.0.0:5000/
2015-11-17 10:05:04.941 7084 INFO eventlet.wsgi.server [-] (7084) wsgi starting up on http://0.0.0.0:5000/
2015-11-17 10:17:01.005 7085 INFO keystone.common.wsgi [-] GET http://localhost:5000/
```

which sounds ok and a `curl` call to the endpoint reply at least something:

```shell
$ curl -s http://localhost:5000/v3 | jsonformat
{
  "version": {
    "id": "v3.4",
    "links": [
      {
        "href": "http://localhost:5000/v3/",
        "rel": "self"
      }
    ],
    "media-types": [
      {
        "base": "application/json",
        "type": "application/vnd.openstack.identity-v3+json"
      }
    ],
    "status": "stable",
    "updated": "2015-03-30T00:00:00Z"
  }
}
```

## The keystone configuration
