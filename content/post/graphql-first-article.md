---
categories:
- category
date: 2017-03-22T09:15:35+01:00
description: ""
draft: true
images:
- https://avatars0.githubusercontent.com/u/12972006
tags:
- tag1
- tag2
title: Trying GraphQL as an API enhancement for AWS products and billing
---

<script type="text/javascript" src="https://ssl.gstatic.com/trends_nrtr/962_RC10/embed_loader.js"></script> <script type="text/javascript"> trends.embed.renderExploreWidget("TIMESERIES", {"comparisonItem":[{"keyword":"graphql","geo":"","time":"2015-09-14 2017-03-23"}],"category":0,"property":""}, {"exploreQuery":"date=2015-09-14%202017-03-23&q=graphql","guestPath":"https://trends.google.com:443/trends/embed/"}); </script> 

To understand correctly the structure of a query, please refer to this excellent blog post: [The Anatomy of a GraphQL Query](https://dev-blog.apollodata.com/the-anatomy-of-a-graphql-query-6dffa9e9e747#.jbklz6h17).

The first thing that needs to be done is to define the [schema](http://graphql.org/learn/schema/)

Here is the description of what can be done.

* A product family is composed of several products
* A product belongs to a product family
* A product owns a set of attributes
* A product and all its attributes are identified by a stock keeping unit (SKU)
* A SKU has a set of offers
* An offer represents a selling contract
* An offer has a term of offer
* An offer term is either "Reserved" or "OnDemand"
*

![Graphql shorthand notation cheat sheet](https://github.com/sogko/graphql-schema-language-cheat-sheet/raw/master/graphql-shorthand-notation-cheat-sheet.png)

{{< highlight graphql >}}
# Define the item inteface
interface Item {
    SKU: String!
}

# User type O
implements Entity interface
type Element implements Entity {
  id: ID!
  name: String
  age: Int
  balance: Float
  is_active: Boolean
  friends: [User]!
  homepage: Url
}

# root Query type
type Query {
  me: User
  friends(limit: Int = 10): [User]!
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
