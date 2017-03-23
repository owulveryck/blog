---
categories:
- category
date: 2017-03-22T09:15:35+01:00
description: "My first experiencewith GraphQL. I will try to see how it fits the pricing model of AWS as described in an earlier post."
draft: true
images:
- http://graphql.org/img/logo.svg
tags:
- golang
- aws
- graphql
title: Playing with Facebook's GraphQL (for AWS products and offers management)
---

# About GraphQL

GraphQL has been invented by Facebook for the purpose of refactoring their mobile application. Facebook had reached the limits of the standard REST API mainly because:

* Getting that much information was requiring a huge amount of API endpoints
* The versioning of the API was counter-productive regarding Facebook's frequents deployements.

But graphql is not only a query language related to Facebook. GraphQL is not only applicable to social data. 

Of course it is about graphs and graphs represents relationships. But you can represent relationships in all of your business objects.

Actually, GraphQL is all about your application data.

In this post I will try to take a concrete use case. I will first describe the business objects as a graph, then I will try to implement a schema with GraphQL. At the very end I will develop a small GraphQL endpoint to test the use case.

__Caution__ _I am discovering GraphQL on my own. This post reflects my own work and some stuff may be inaccurate or not idiomatic._

## The use case: AWS billing

Let's take a concrete example of a graph representation. Let's imagine that we are selling products related to Infrastructre as a Service (_IaaS_). 

For the purpose of this post, I will use the AWS data model because it is publicly available and I have already blogged about it.
We are dealing with products families, products, offers and prices.

In (a relative) proper english, let's write down a description of the relationships:

* Products
  * A product family is composed of several products
  * A product belongs to a product family
  * A product owns a set of attributes (for example its location, its operating system type, its type...)
  * A product and all its attributes are identified by a stock keeping unit (SKU)
  * A SKU has a set of offers
* Offers
  * An offer represents a selling contract
  * An offer is specific to a SKU
  * An offer is characterized by the term of the offer
  * A term is typed as either "Reserved" or "OnDemand"
  * A term has attributes
* Prices
  * An offer has at least one price dimension
  * A price dimension is caracterized by its currency, its unit of measure, its price per unit, its desription and eventually per a range of application (start and end)

Regarding those elements, I have extracted and represented a "t2.micro/linux in virginia" with 3 of its offers and all the prices associated.

Here is the graphical representation generated thanks to [graphviz' fdp](http://www.graphviz.org/)
![Graph Representation](/assets/graphql/graph.svg)


The goal of GraphQL is to extract a substree of this graph to get part or all information.
As an example, here is a tree representation of the same graph:

![Graph Representation](/assets/graphql/graph_tree.svg)

_Note_: I wrote a very quick'n'dirty parser to get the information which you can find [here](https://gist.github.com/owulveryck/bac700e2f5e5b1af0fffda4e7adb9eed). I wrote an idiomatic one but it is the property of the company I made it for.

# Defining the GraphQL schema

The first thing that needs to be done is to write the [schema](http://graphql.org/learn/schema/) that will define the _query_ type.

I will not go into deep details in here. I will simple refer to this excellent document which is a _résumé_ of the language:
[Graphql shorthand notation cheat sheet](https://github.com/sogko/graphql-schema-language-cheat-sheet/raw/master/graphql-shorthand-notation-cheat-sheet.png)

We can define a product that must contains a list of offers thsi way and a product family like this:

{{< highlight graphql >}}
# Product definition
type Product {
  offers: [Offer]!
  location: String
  type: String
  SKU: String!
  operatingSystem: String
}

# Definition of the product family
type ProductFamily {
  products: [Product]!
}
{{< /highlight >}}

One offer is composed of a mandatory price list. An offer must be of a pre-defined type: _OnDemand_ or _Reserved_.
Let's define this:
{{< highlight graphql >}}
# Definition of an offer
type Offer {
  type: OFFER_TYPE!
  code: String!
  LeaseContractLength: String!
  PurchaseOption: String!
  OfferingClass: String
  prices: [Price]!
}

# All possible offer types
enum OFFER_TYPE {
  OnDemand
  Reserved
}

# Definition of a price
type Price {
  description: String
  unit: String
  currency: String
  price: Float
}
{{< /highlight >}}

At the very end we define the _queries_ 
Let's start by defining a single query. To make it simple for the purpose of the post, Let's assume that we will try to get a whole _product family_.
If we qurey the entire product family, we can be able to display all informations of all product in the family. But let's also consider that we want to limit the family and extract only a certain product identified by its SKU.

The Query definition is therfore:
{{< highlight graphql >}}
# root Query type
type Query {
    ProductFamily(sku: String): productFamily
}
{{< /highlight >}}

## Query

Let's see now how a typical query would look like. To understand the structure of a query, I advise you to read this excellent blog post: [The Anatomy of a GraphQL Query](https://dev-blog.apollodata.com/the-anatomy-of-a-graphql-query-6dffa9e9e747#.jbklz6h17).

{{< highlight graphql >}}
{
  ProductFamily {
    products {
      location
      type
    }
  }
}
{{< /highlight >}}

This query should normally return all the products of the family and display their location and their type.
Let's try to implement this

# Geek time: let's go!

I will use the `go` implementation of GraphQL which is a "simple" translation in go of the [javascript's reference implementation](https://github.com/graphql/graphql-js).

To use it: 

{{< highlight go >}}
import "github.com/graphql-go/graphql"
{{< /highlight >}}

To keep it simple, I will load all the products and offers in memory. In the real life, we should implement an access to whatever database. But that is a strength of the GraphQL model: The flexibility. The backend can be changed later without breaking the model or the API.

## Defining the schema and the query in go

