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

_Note_ : I will consider that the openstack keystone is installed. As I don't want to rewrite an installation procedure as many exists already on the web. For my tests, I'm using an keystone installation from sources in a Ubuntu VM

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

```json
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

The proper keystone configuration is done in the file `keystone.conf`. This configuration file is decomposed into different sections as explained in the documentation.

### The general configuration (Default section)
I will only set the `admin token` randomly as it will be used to create the users, roles and so on.

Let's generate a token with `openssl rand -hex 10` and report it to my configuration:
```
[DEFAULT]
admin_token = 8a0b4eacc6a81c3bc5aa
```

The rest will use all the default values for the General configuration (the [DEFAULT] section). This means that this section may be empty or full of comments.


### The assignment configuration
In this section, we choose the driver for the assignment service.
This purpose of this service is

> [to] provide data about roules and role assignments 
> to the entities managed by the Identity and Resource services

(source [Keystone architecture](http://docs.openstack.org/developer/keystone/architecture.html))

I plan to use a SQL backend instead of a LDAP for my tests, so the configuration may be:
```
[assignment]
driver = sql
```

### The authentication plugin configuration
Keystone supports authentication plugins; those plugins are specified in the `[auth]` section.
In my test, the `password` plugin will be used.

```
[auth]
methods = password
```

### The credentials
The credentials are stored in a sql database as well:

```
[credential]
driver = sql
```

### The DB configuration
For my tests I will use a sqlite database as configured in this section:
```
[database]
sqlite_db = oslo.sqlite
sqlite_synchronous = true
backend = sqlalchemy
connection = sqlite:////var/lib/keystone/keystone.db

```

## Restart the keystone server and play 


```
# service keystone restart
# service keystone status
● keystone.service - OpenStack Identity service
   Loaded: loaded (/lib/systemd/system/keystone.service; enabled; vendor preset: enabled)
   Active: active (running) since Tue 2015-11-17 14:47:06 GMT; 3s ago
  Process: 15505 ExecStartPre=/bin/chown keystone:keystone /var/lock/keystone /var/log/keystone /var/lib/keystone (code=exited, status=0/SUCCESS)
  Process: 15502 ExecStartPre=/bin/mkdir -p /var/lock/keystone /var/log/keystone /var/lib/keystone (code=exited, status=0/SUCCESS)
 Main PID: 15508 (keystone-all)
   CGroup: /system.slice/keystone.service
           ├─15508 /usr/bin/python /usr/bin/keystone-all --config-file=/etc/keystone/keystone.conf --log-file=/var/log/keystone/keystone.log
           ├─15523 /usr/bin/python /usr/bin/keystone-all --config-file=/etc/keystone/keystone.conf --log-file=/var/log/keystone/keystone.log
           ├─15524 /usr/bin/python /usr/bin/keystone-all --config-file=/etc/keystone/keystone.conf --log-file=/var/log/keystone/keystone.log
           ├─15525 /usr/bin/python /usr/bin/keystone-all --config-file=/etc/keystone/keystone.conf --log-file=/var/log/keystone/keystone.log
           └─15526 /usr/bin/python /usr/bin/keystone-all --config-file=/etc/keystone/keystone.conf --log-file=/var/log/keystone/keystone.log

Nov 17 14:47:08 UBUNTU keystone[15508]: 2015-11-17 14:47:08.479 15508 INFO oslo_service.service [-] Started child 15523
Nov 17 14:47:08 UBUNTU keystone[15508]: 2015-11-17 14:47:08.482 15508 INFO oslo_service.service [-] Started child 15524
Nov 17 14:47:08 UBUNTU keystone[15508]: 2015-11-17 14:47:08.486 15508 INFO keystone.common.environment.eventlet_server [-] Starting /usr/bin/keystone-all on 0.0.0.0:5000
Nov 17 14:47:08 UBUNTU keystone[15508]: 2015-11-17 14:47:08.490 15508 INFO oslo_service.service [-] Starting 2 workers
Nov 17 14:47:08 UBUNTU keystone[15508]: 2015-11-17 14:47:08.491 15523 INFO eventlet.wsgi.server [-] (15523) wsgi starting up on http://0.0.0.0:35357/
Nov 17 14:47:08 UBUNTU keystone[15508]: 2015-11-17 14:47:08.493 15508 INFO oslo_service.service [-] Started child 15525
Nov 17 14:47:08 UBUNTU keystone[15508]: 2015-11-17 14:47:08.499 15524 INFO eventlet.wsgi.server [-] (15524) wsgi starting up on http://0.0.0.0:35357/
Nov 17 14:47:08 UBUNTU keystone[15508]: 2015-11-17 14:47:08.502 15508 INFO oslo_service.service [-] Started child 15526
Nov 17 14:47:08 UBUNTU keystone[15508]: 2015-11-17 14:47:08.506 15525 INFO eventlet.wsgi.server [-] (15525) wsgi starting up on http://0.0.0.0:5000/
Nov 17 14:47:08 UBUNTU keystone[15508]: 2015-11-17 14:47:08.510 15526 INFO eventlet.wsgi.server [-] (15526) wsgi starting up on http://0.0.0.0:5000/
```

so far so good... let's check if the DB is here now:

```
# sqlite3 /var/lib/keystone/keystone.db
SQLite version 3.8.11.1 2015-07-29 20:00:57
Enter ".help" for usage hints.
sqlite> .tables
access_token            identity_provider       revocation_event
assignment              idp_remote_ids          role
config_register         mapping                 sensitive_config
consumer                migrate_version         service
credential              policy                  service_provider
domain                  policy_association      token
endpoint                project                 trust
endpoint_group          project_endpoint        trust_role
federation_protocol     project_endpoint_group  user
group                   region                  user_group_membership
id_mapping              request_token           whitelisted_config
sqlite> .quit
```

## Interacting with openstack

A tools called [python-openstackclient](http://docs.openstack.org/developer/python-openstackclient/command-list.html) is available in my ubuntu release and will be used for testing purpose.

The binary provided is `openstack` (`dpkg-query -L python-openstackclient | grep bin`)

### Creating a user

We need to define a couple of environment variables to be able to connect to the keystone server:
```
export OS_TOKEN=8a0b4eacc6a81c3bc5aa # The value of admin_token defined in the keystone.conf
export OS_URL=http://localhost:35357/v2.0 # This is the default value if not overridden by the directive admin_endpoint
export OS_IDENTITY_API_VERSION=3
```

Then we create the user: 
```
openstack user create olivier
'links'
```

Then set its password:
```
openstack user set --password-prompt olivier
User Password:
Repeat User Password:
'users'
```

And see if it's actually here:
```
openstack user list
+----------------------------------+---------+
| ID                               | Name    |
+----------------------------------+---------+
| c80f5244c7d3486fbf4059b7197b4770 | olivier |
+----------------------------------+---------+
```

