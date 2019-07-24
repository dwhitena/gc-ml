# Project Ideas and Data Sets

Awesome! You've chosen to train and evaluate some of your own ML/AI models using new data sets. We have compiled a set of problems with corresponding public data available (listed below).

1. Assemble a group of 3-4 people (try to introduce yourself to some new gopher friends)
2. Pick a project/problem
3. Explore the data 
4. Train/evaluate/integrate your model

_Hint: If you don't now where to start check out [this step by step guide for ML projects](https://github.com/ageron/handson-ml/blob/master/ml-project-checklist.md)_

_Hint: If you are looking for more relevant Go tooling, see [here](https://github.com/gopherdata/resources/blob/master/tooling/README.md)._

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
* [Machine Box's Textbox](https://docs.veritone.com/#/developer/machine-box/boxes/textbox) - Machine Box provides state of the art machine learning technology inside a Docker container which you can run, deploy and scale. Their textbox offering allows you to do sentiment analysis of text with a deep learning model. Moreover, they have a [Go SDK](https://github.com/machinebox/sdk-go) so you can easily integrate this into your Go applications.

## Deep Learning for Image/Object Classification

If you are following the news, you have probably heard a bunch about deep learning methods in computer vision (think deep fakes or FaceApp). So why not try your hand at doing a little bit of image processing? Some cool application/project ideas would be to:

- Write a Go program that can detect your face in an image.
- Use MachineBox's Tagbox to automatically tag the content of images. Then try to add your own custom tags and images to update the model.

### Datasets:

* [mnist](https://www.kaggle.com/c/digit-recognizer/data)
* [open images object detection](https://www.kaggle.com/c/open-images-2019-object-detection/data)
* [malaria detection](https://ceb.nlm.nih.gov/repositories/malaria-datasets/)
* [crowdsourced object detection data](https://ai.google/tools/datasets/open-images-extended-crowdsourced/)
### Tools:
* [Machinebox's Objectbox](https://docs.veritone.com/#/developer/machine-box/boxes/objectbox) - Machine Box provides state of the art machine learning technology inside a Docker container which you can run, deploy and scale. Their object box offering allows you to do object detection using a deep learning model. Moreover, they have a [Go SDK](https://github.com/machinebox/sdk-go) so you can easily integrate this into your Go applications.
* [Machinebox's Classificationbox](https://docs.veritone.com/#/developer/machine-box/boxes/classificationbox) - Classificationbox allows you to classify text, images, structured and unstructured data using a deep learning model. 
* [Machinebox's Facebox](https://docs.veritone.com/#/developer/machine-box/boxes/facebox-overview) - Facebox allows you to teach and recognize faces in images or photographs using a deep learning model.
* [Machinebox's Tagbox](https://docs.veritone.com/#/developer/machine-box/boxes/tagbox) - Tagbox allows you to teach and automatically understand the content of images using a deep learning model. 

## Linear and Logistic Regression

These methods might not be as hyped as the above methods, but they are certainly useful and a little easier to implement. You could take some of the workshop code from earlier in the day and adapt it to these new data sets! Some cool application/project ideas would be to:

- Use census data to try and predict whether a person makes over 50K a year based on certain demographic info.
- Use logistic regression to predict Iris flower species
- Predict outcomes of an NCAA basketball matchup.

### Datasets:

* [census income data](https://archive.ics.uci.edu/ml/datasets/census+income)
* [iris classification](https://archive.ics.uci.edu/ml/datasets/Iris)
* [Bureau of Labor income statistics](https://console.cloud.google.com/marketplace/details/bls-public-data/bureau-of-labor-statistics?filter=solution-type:dataset&id=e632a715-857e-4c41-8257-da123607ea89)
* [NCAA basketball statistics](https://console.cloud.google.com/marketplace/details/ncaa-bb-public/ncaa-basketball?filter=solution-type%3Adataset&id=f262fa22-2021-44c6-a628-15eab8237de5)

### Tools:
* [Workshop examples of linear regression](linear_regression)
* [Workshop examples of logistic regression](logistic_regression)
* [goml](https://github.com/cdipaolo/goml)
* [golearn](https://github.com/sjwhitworth/golearn)

