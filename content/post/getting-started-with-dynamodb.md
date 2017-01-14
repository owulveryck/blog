---
categories:
- aws
date: 2017-01-13T22:22:46+01:00
description: "In this post, I will explain how to extract, process and store informations from a webservice to a NoSQL database (DynamoDB)"
draft: true
images:
- /assets/images/bigdata/stones-483138_640.png
tags:
- dynamodb
- aws
- golang
title: A feet in NoSQL and a toe in big data
---

The more I work with AWS, the more I understand their models. This goes far beyond the technical principles of micro service.
As an example I recently had an opportunity to dig a bit into the billing process.
I had an explanation given by a colleague whose understanding was more advanced than mine.
In his explanation, he mentionned this blog post: [New price list API](https://aws.amazon.com/blogs/aws/new-aws-price-list-api/).

# Understanding the model
By reading this post and this [explanation](http://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/price-changes.html), I understand that the offers are categorized in families (eg AmazonS3) and that an offer is composed of a set of products.
Each product is caracterized by its SKU's reference ([stock-keeping unit](https://en.wikipedia.org/wiki/Stock_keeping_unit))

## Inventory management

So finally, it is just about inventory management. In the retail, when you say "inventory management", the IT usually replies with millions dollars _ERP_.
And the more items we have, the more processing power we need and then more dollar are involved... and richer the IT specialists are (just kidding).

Moreover enhancing an item by adding some attributes can be painfull and risky

![xkcd](http://imgs.xkcd.com/comics/exploits_of_a_mom.png)

## The NoSQL approach 

Due to the rise of the online shopping, inventory management must be real time.
The stock inventory is a business service. and placing it in a micro service architecture bring constraints: the request should be satisfied in micro seconds.

More over, the key/value concept allows to store "anything" in a value. Therefore, you can store a list of attributes regardless of what the attributes are.

When it comes to NoSQL, there are usuallly two approches to store the data:

* simple Key/Value;
* document-oriented.

At first I did and experiment with a simple key/value store called BoldDB (which is more or less like Redis).
In this approch the value stored was a json representation... A kind of document.
Then I though that it could be a good idea to use a more document oriented service: DynamoDB

# Geek time

In this part I will explain how to get the data from AWS and to store them in the dynamoDB service. The code is written in GO and is just a proof of concept.

## The product informations

A product is described [here](http://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/reading-an-offer.html).
We have:

{{< highlight js >}}
"Product Details": {
   "sku": {
      "sku":"The SKU of the product",
      "productFamily":"The product family of the product",
      "attributes": {
         "attributeName":"attributeValue",
      }
   }
}
{{< /highlight >}}

There are three important entries:

* SKU: A unique code for a product. 
* Product Family The category for the type of product. For example, compute for Amazon EC2 or storage for Amazon S3.
* Attributes: A list of all of the product attributes.

## Creating the "table"

As my goal is for now to create a proof of concept and play with the data, I am creating the table manually.
DynamoDB allows the creation of two indexes per table. So I create a table products with two indexes:

* the SKU
* the product family

![Create Table](/assets/images/bigdata/blog-dynamo-create-table.png)

## Grabbing the data

### Fetching

### Decoding

## Storing the informations

# Execution and conclusion

![Result](/assets/images/bigdata/blog-dynamo-result.png)

{{< gist owulveryck f9665470e8334e8609434feeeddc6071 "putproducts.go" >}}


