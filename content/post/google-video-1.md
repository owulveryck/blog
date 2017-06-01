---
categories:
- category
date: 2017-06-01T22:07:56+02:00
description: ""
draft: true
images:
- /2016/10/image.jpg
tags:
- tag1
- tag2
title: Analysing a parodic trailer with Google Cloud Video Intelligence
---

# The sample movie 

As an example, I will use a trailer from a french parody called _"A la recherche de l'utra-sex"_.

_Caution_: this video **is not a porn video** but is indeed **not safe for work** (_#nsfw_)


## Querying Google Cloud Video Intelligence 

{{< highlight js >}}
{
    "inputUri": "gs://video-test-blog/trailer.mp4",
    "features": ["SHOT_CHANGE_DETECTION","LABEL_DETECTION"]
}
{{< /highlight >}}

{{< highlight shell >}}
curl -s -k -H 'Content-Type: application/json' \
      -H 'Authorization: Bearer MYTOKEN' \
      'https://videointelligence.googleapis.com/v1beta1/videos:annotate' \
      -d @demo.json
{
   "name": "us-east1.16784866925473582660"
}
{{< /highlight >}}

{{< highlight shell >}}
curl -s -k -H 'Content-Type: application/json' \
      -H 'Authorization: Bearer MYTOKEN' \
      'https://videointelligence.googleapis.com/v1/operations/us-east1.16784866925473582660'
{{< /highlight >}}

The full result is [here](/assets/video-intelligence/video-analysis-a-la-recherche.json)
{{< highlight js >}}
{
  "response": {
    "annotationResults": [
      {
        "shotAnnotations": [
          {
            "endTimeOffset": "1920048"
          },
          // ...
          {
            "endTimeOffset": "109479985",
            "startTimeOffset": "106479974"
          }
        ],
        "labelAnnotations": [
          {
            "locations": [
              {
                "level": "SHOT_LEVEL",
                "confidence": 0.40029016,
                "segment": {
                  "endTimeOffset": "24360033",
                  "startTimeOffset": "21559963"
                }
              },
              {
                "level": "SHOT_LEVEL",
                "confidence": 0.41241992,
                "segment": {
                  "endTimeOffset": "29039958",
                  "startTimeOffset": "26999971"
                }
              },
              {
                "level": "SHOT_LEVEL",
                "confidence": 0.41364595,
                "segment": {
                  "endTimeOffset": "30639998",
                  "startTimeOffset": "29080023"
                }
              }
            ],
            "languageCode": "en-us",
            "description": "Abdomen"
          },
          // ... 
          {
            "locations": [
              {
                "level": "SHOT_LEVEL",
                "confidence": 0.8738658,
                "segment": {
                  "endTimeOffset": "85080015",
                  "startTimeOffset": "83840048"
                }
              }
            ],
            "languageCode": "en-us",
            "description": "Acrobatics"
          },
        ],
        "inputUri": "/video-test-blog/trailer.mp4"
      }
    ],
    "@type": "type.googleapis.com/google.cloud.videointelligence.v1beta1.AnnotateVideoResponse"
  },
  "done": true,
  "metadata": {
    "annotationProgress": [
      {
        "updateTime": "2017-06-01T20:53:42.679081Z",
        "startTime": "2017-06-01T20:53:30.610811Z",
        "progressPercent": 100,
        "inputUri": "/video-test-blog/trailer.mp4"
      },
      {
        "updateTime": "2017-06-01T20:53:45.472996Z",
        "startTime": "2017-06-01T20:53:30.610811Z",
        "progressPercent": 100,
        "inputUri": "/video-test-blog/trailer.mp4"
      }
    ],
    "@type": "type.googleapis.com/google.cloud.videointelligence.v1beta1.AnnotateVideoProgress"
  },
  "name": "us-east1.9773682320661713538"
}
{{< /highlight >}}

<iframe width="100%" height="400" src="/assets/video-intelligence/result.html"></iframe>
