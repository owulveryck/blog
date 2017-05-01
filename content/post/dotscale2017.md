---
categories:
- category
date: 2017-04-24T20:51:22+02:00
description: ""
draft: true
images:
- https://pbs.twimg.com/profile_images/659732730784129025/TlfVUxEl.png
tags:
- tag1
- tag2
title: Retour sur les conférences dotScale, edition 2017.
---

# A propos de la dotScale

Lundi 24 avril se tenait la cinquième édition de la conférence dotScale qui se décrit elle-même comme étant la _tech conférence Européenne sur la "scalabilité"_
Cette conférence se veut être le "[TED](https://www.ted.com/)" des informaticiens.
Elle se déroule dans le prestigieux théâtre de Paris. 
Les orateurs se succèdent pour des présentations d'une vingtaine de minutes maximum. Les conditions sont donc réunies pour mener les geeks et autres professionnels au plus profond des sujets évoqués. 
D'ailleurs, les ordinateurs sont bannis (c'est une conférence _wireless-less_).
Le rythme est cadré et maîtrisé par les organisateur. Bref, on se laisse guider.

Cette année encore les orateurs étaient prestigieux. L'avantage de cette conférence, c'est que les orateurs ne sont pas des vendeurs. Ils viennent simplement pour partager leur vision et leur expérience de la scalabilité.

## La "tradition"

Inspirée de "TED", chaque conférence "dot" débute par une tradition. Il s'agit de se présenter rapidement à son premier voisin inconnu. Je ne vais pas m'étendre sur le bienfait d'un tel acte, mais c'est quelque chose qui vaut la peine d'être fait.

# Les conférences

## Neha Narkhede - _The rise of real time_

La première conférence fût donnée par Neha Narkhede.

L'introduction part d'un constat.

Les données gérées par les entreprises ont souvent été regroupées en bloques de travail pour qu'ils puissent être traités au mieux (en comptabilité, on parle par exemple "d'arrêter les comptes"). 
Ce mode de fonctionnement, copié par l'informatique a donné lieu aux traitements "batchs"
Le problème est que ce mode de traitement n'est, de nos jours, plus du tout adapté au flux de données de plus en plus important.

Alors que faire ? Une version plus rapide ou plus efficace de ce qui a été fait jusqu'à lors ? Probablement pas.

L'idée est désormais de traiter le flux de données et non les données elles-mêmes.

Neha Narkhede est co-créatrice du projet [Apache Kafka](http://kafka.apache.org/) qui est un système de messaging realtime. C'est le système utilisé par linkedin et il est capable de traiter 1,4 milliard de message par jour. Elle explique dans les minutes qui lui restent, comment kafka permet d'urbaniser les services en faisant transiter des messages qui transportent de l'information.
Puis elle présente son produit [confluent](https://www.confluent.io/) qui est une plateforme de streaming de données qui se base sur kafka. 

Le produit semble intéressant, mais l'idée l'est plus encore: traiter un flux de données à tous les niveaux, de l'applicatif aux logs.

En attendant la video de la conférence qui devrait se trouver sur [www.dotconferences.io/talks](www.dotconferences.com/talks).

{{< speakerdeck b962268ed0724e7eb64d2d79c0c9fac2 >}}

## Adrian Cole - _Observability 3 ways / Logging, Metrics, Tracing_

On reste dans le flux de données et le realtime avec la présentation d'Adrian Cole (Lead Developer du projet Zipkin).
Il est 11h, c'est la seconde conférence, et on constate déjà que les problèmes de scalabilité de "2017" tournent autour de la quantité de données qui peuvent transiter entre les éléments d'une application distribuée.

[Zipkin](http://zipkin.io/) est un système de tracing distribué. Son rôle est de collecter des données horodatées pour analyser les problèmes de architecture microservices.

Adrian présente ici 3 techniques différentes de collectes qui permettent aux ingénieurs d'avoir la vision totale de l'activité d'une plateforme de production: 

* Logging 
* Metrics
* tracing

On retrouvera des détails de ces principes sur le [blog de Peter Bourgon](https://peter.bourgon.org/blog/2017/02/21/metrics-tracing-and-logging.html)

Dans sa présentation, Adrian expliquera comment mettre en place ces principes et avec quels outils. 

{{< speakerdeck a3d9e5a710634bca8b4180e26ca50f73 >}}

## Mitchell Hashimoto - _Scaling security / Grow without compromising security_

Mitchell Hashimoto est le fondateur de la société Hashicorp. Hashicorp est l'éditeur d'un très grand nombre d'outils reconnus pour leur utilité dans le cloud et le monde devops.
On pourra citer par exemple:

* Vagrant
* Terraform
* Consul
* ...

Le sujet exposé aujourd'hui concerne la sécurité. Mitchell commence par replacer la scalabilité dans un contexte DevOps. C'est le _Pourquoi_.
Ainsi il explique que le modèle traditionnel fait passer le travail du développeur à l'ops, et de l'ops à la production.
Le DevOps a permis une paralellisation des tâches. Ainsi Dev et Ops travaillent de consors pour livrer en production.

Mais quid de la sécurité. Selon Mitchell Hashimoto, trop souvent la sécurité se place en "poseur de veto" entre les "DevOps" et la mise en production.

{{< speakerdeck 0decb31f9c374c21a5c0899fa89e39d1 >}}

## Ulf Adams - _Build and test performance at scale_

Quand on parle de systèmes "scalables", c'est souvent pour évoquer les environnements d'exécution des applications. Les systèmes distribués sont conçus pour permettre la tolérance aux pannes ainsi que l'efficacité (équilibrage de charge, CDN, ...)

Ulf Adams expose ici une autre problématique de scalabilité: la "construction" de l'application. 

En effet, Google dispose du plus gros écosystème logiciel de la planète. Il est de plus connu que Google utilise un [repository monolithique](https://cacm.acm.org/magazines/2016/7/204032-why-google-stores-billions-of-lines-of-code-in-a-single-repository/fulltext) pour versionner ses sources. Construire et jouer les tests unitaires de cet écosystème est de ce fait un défi.

Pour répondre à cette problématique, Google a développé un logiciel spécifique: _Blaze_. En 2015, Google a rendu open-source le coeur de ce logiciel sous le nom de [Bazel](https://bazel.build/).

Ulf expose succinctement comment Bazel détecte les changements et se charge de ne reconstruire *que ce qui est nécessaire* de manière concurrente.

Pour ma part, je ne connaissais pas Bazel. Ainsi donc la conférence a été une excellente introduction sur l'outil.

## Marco Slot - _Scaling out (Postgre)SQL_

Marco Slot pose le constat suivant:

Quand il s'agit de base de données et de scalabilité, on ne dissocie jamais les deux termes suivant: No et SQL.
En effet les bases de données NoSQL sont nées pour répondre au besoin de scalabilité des entreprises.

Cependant, il est tout à fait possible de rendre pleinement scalable une architecture SQL classique. Ainsi, en expliquant l'architecture modulaire de PostGres et en la faisant converger avec un modèle arithmétique, Marco explique comment il est possible de mettre en place une architecture "SQL" résiliente et scalable.

La demonstration est réussie

{{< speakerdeck 80cc7e9225914845921786440affee1d >}}

## James Cammarata - _Keep calm and automate all the things_

James Cammarata est le "lead maintener" du projet Ansible.
Il fait ici l'apologie de l'automatisation (mais est-il encore besoin de le faire). Là où James insiste, c'est qu'il faut s'astreindre à respecter les règles de l'art en matière d'automatisation:

* Tout automatiser (tout!)
* Utiliser les fonction natives des outils lorsque c'est possible, et c'est souvent possible (éviter les appels à `sudo rm` dans les playbooks par exemple)
* Préfixer les variables pour éviter le désastre du `sudo rm -rf ${foo}/${bar}`

En bref:

<center>
![Keep calm and automate all the things](/assets/images/keep-calm-and-automate-all-the-things-small.jpg)
</center>

Autre info : le produit Tower est planifié pour être rendu opensource par la société RedHat qui possède le projet Ansible.

## Aish Raj Dahal - _Chaos management during a major incident_

{{< speakerdeck a01da34216124ae9bbd9f589ea9bb480 >}}

## Andrew Shafer - _DevOps_

Andrew Shafer est un "Jedi" du DevOps. C'est également un présentateur excellent.

## Clay Smith - _Searching for the server in serverless_ 

Clay Smith pose la question de la latence dans les architecture "serveur-less".
Il expose ainsi les recherches qu'il a faites autour de [Lambda (aws)](https://aws.amazon.com/fr/lambda/details/). Il montre de prime abord comment il a réussi à obtenir des informations sur l'environnement d'exécution du code (accès via SSH depuis le lambda vers une autre machine avec tunneling). 


Clay expose un ensemble de graphiques qui permettent affiche clairement un signal. Ce signal lui permet d'émettre certaines hypothèses quant-à l'algorithme possible d'attribution des environnements lambda sur AWS.

Bien qu'elle semble peut utile de prime abord, la démonstration mets en exergue certaines problématiques du "serverless" à savoir le temps de démarrage des applications qui est non constant.

## David Mazieres - _ _



## Benjamin Hindman - _ _ 

# Conclusion

Ces conférences reflètent le savoir-faire et les bonnes pratiques actuelles de grandes sociétés ou de produits "à la mode".
Beaucoup de concepts font écho avec les concepts énoncés dans l'ouvrage [Site Reliability Engineering - How Google Runs Production Systems](https://landing.google.com/sre/).

Il y a donc beaucoup de choses intéressantes à tirer de ces conférence, bien au-delà des aspects techniques. Tout comme les conférences TED, elles permettent une ouverture d'esprit ainsi qu'un regard critique sur notre activité au quotidien, que l'on soit, Ops, Dev ou Consultant.

Pas besoin d'être une startup pour profiter de ces conseils, et les problématiques que rencontrent les conférenciers sont en grande partie les mêmes que celles que rencontrent nos clients.
