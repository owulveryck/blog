+++
date = "2015-10-23T09:54:27+01:00"
draft = true
title = "Simple polling app, cloud native"

+++

In this post I explain how to setup a simple polling app.
This app, written in go, will be hosted on a PAAS, and I've chosen the [Google App Engine](https://cloud.google.com/appengine/docs) for convenience.

# Born in the cloud

I'm 38, and I'm qualified as a person "_born in the datacenter_". Developping an app _born in the cloud_ is not an evidence for me.
When i'm conceiving an application, or just imaginig a service, I'm still thinking of the technical infrastructure, keeping in mind the technical solution for hosting, resiliency and clustering.

* I'm thinking of the need of a frontend for proxyfing the application and anticipate the scalability. 
* I'm thinking about the storage: which FS, which backup solution, snapshots, SAN, ... 
* I'm evaluating the hosting solution and the application isolement (container, jail, chroot)... 

Anyway, I'm thinking the hosting of the application in the Datacenter way, and any developpement is compliant with this architecture.


