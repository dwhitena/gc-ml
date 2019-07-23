# Project Ideas and Data Sets
Awesome! You've chosen to train some of your own ML/AI models using new data sets. During this time, you are more than welcome to try your hand at some regression/classification problem you are facing at work, assuming you have the right data. However, we thought it may be useful to compile a set of well defined problems where there is public data available. Those are listed below. 

Find a problem that looks interesting, assemble a team of 3-4 people, and _Go_ for it! Try using some of the methods we have introduced earlier in the day, or search for some other model that might be relevant. If you are having trouble figuring out where to start, ask the instructors and TAs. If you are looking for more relevant Go tooling, see [here](https://github.com/gopherdata/resources/blob/master/tooling/README.md).

## Deep Learning for Sentiment Analysis

Sentiment analysis is a method that tries to determine whether a piece of text is generally _positive_ or _negative_ in tone and/or content. Certain implementations do this by predicting a score/number that is indicative of sentiment, while other implementions just return binary indications of sentiment. Gather some text data from one of the sources below and try to determine the sentiment of that text data. Some cool application/project ideas would be to:

- Write a Go program that streams in tweets from Twitter and aggregates or updates a measure of sentiment on a certain topic or hashtag.
- Answer some questions about the content on Hacker News. What is the overall sentiment of comments on the front page Hacker news articles? Which ones have the most positive sentiment? 

### Datasets:

* [twitter - api](https://developer.twitter.com/en/use-cases/analyze)
* [twitter - Max-plank institute](http://twitter.mpi-sws.org/)
* [hacker news comments](https://console.cloud.google.com/marketplace/details/y-combinator/hacker-news?filter=solution-type%3Adataset&id=5227103e-0eb9-4744-872b-325a8df50bee) 
* [stack  exchange](https://console.cloud.google.com/marketplace/details/stack-exchange/stack-overflow?filter=solution-type:dataset&id=46a148ff-896d-444c-b08d-360169911f59)

### Tools:
* [Machinebox's TextBox](https://blog.machinebox.io/introducing-textbox-natural-language-processing-inside-a-docker-container-bdb57a2a3e64) 

## Deep Learning for Image/Object Classification

If you are following the news, you have probably heard a bunch about deep learning methods in computer vision (think deep fakes or FaceApp). So why not try your hand at doing a little bit of image processing? Some cool application/project ideas would be to:

- Write a Go program that could detect 

### Datasets:

* [mnist](https://www.kaggle.com/c/digit-recognizer/data)
* [object detection](https://www.kaggle.com/c/open-images-2019-object-detection/data)
* [malaria detection](https://ceb.nlm.nih.gov/repositories/malaria-datasets/)
* [object detection](https://ai.google/tools/datasets/open-images-extended-crowdsourced/)
### Tools:
* [Machinebox's <insert>](https://blog.machinebox.io/introducing-textbox-natural-language-processing-inside-a-docker-container-bdb57a2a3e64) 

## Linear and Logistic Regression

### Datasets:
* [census income data](https://archive.ics.uci.edu/ml/datasets/census+income)
* [iris classification](https://archive.ics.uci.edu/ml/datasets/Iris)
* [Bureau of Labor income statistics](https://console.cloud.google.com/marketplace/details/bls-public-data/bureau-of-labor-statistics?filter=solution-type:dataset&id=e632a715-857e-4c41-8257-da123607ea89)
* [NCAA basketball statistics](https://console.cloud.google.com/marketplace/details/ncaa-bb-public/ncaa-basketball?filter=solution-type%3Adataset&id=f262fa22-2021-44c6-a628-15eab8237de5)

### Tools:
* Workshop examples of linear regression
* Workshop examples of logistic regression
* goml
* [golearn](https://github.com/sjwhitworth/golearn)

