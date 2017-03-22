---
categories:
- category
date: 2017-03-22T09:15:35+01:00
description: ""
draft: true
images:
- /2016/10/image.jpg
tags:
- tag1
- tag2
title: graphql first article
---

To understand correctly the structure of a query, please refer to this excellent blog post: [The Anatomy of a GraphQL Query](https://dev-blog.apollodata.com/the-anatomy-of-a-graphql-query-6dffa9e9e747#.jbklz6h17).

{{< highlight js >}}
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
