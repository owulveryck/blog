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
title: Trying GraphQL as an API enhancement for AWS products and billing
---

# About graphql

Graph because of App Data Graph. Business objects are connected with each others


To understand correctly the structure of a query, please refer to this excellent blog post: [The Anatomy of a GraphQL Query](https://dev-blog.apollodata.com/the-anatomy-of-a-graphql-query-6dffa9e9e747#.jbklz6h17).

# The use case: AWS billing

The first thing that needs to be done is to define the [schema](http://graphql.org/learn/schema/)

Here is the description of what can be done.

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


The aim of graphql is to extract a substree of this graph to get part or all information.
As an example, here is a tree representation of the same graph:

![Graph Representation](/assets/graphql/graph_tree.svg)

[Graphql shorthand notation cheat sheet](https://github.com/sogko/graphql-schema-language-cheat-sheet/raw/master/graphql-shorthand-notation-cheat-sheet.png)

## Example query

# Defining the GraphQL schema

{{< highlight graphql >}}
type Product {
  offers: [Offer]!
  location: String
  type: String
  SKU: String!
  operatingSystem: String
}

type Offer {
  type: OFFER_TYPE!
  code: String!
  LeaseContractLength: String!
  PurchaseOption: String!
  OfferingClass: String
  prices: [Price]!
}

enum OFFER_TYPE {
  OnDemand
  Reserved
}

type Price {
  description: String
  unit: String
  currency: String
  price: Float
}

type ProductFamily {
  products: [Product]!
}

# root Query type
type Query {
  productFamily: User
}
{{< /highlight >}}


{{< highlight graphql >}}
{
  products(location="EU (London)",type="r3.8xlarge") {
    location
    type
    offers(term_type="Reserved",lease_contact_length="1yr",offering_class="standard",purchase_option="No Upfront"){
      term_type
      lease_contract_length
      offering_class
      purchase_option
      prices(unit="Quantity") {
        description
        unit
        currency
        price_per_unit
      }
    }
  }
}
{{< /highlight >}}
