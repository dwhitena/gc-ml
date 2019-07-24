![Alt text](https://docs.google.com/drawings/d/e/2PACX-1vT37glyZXd8ViXedt0LCSpzsbWCUSSLhWuR3o5_74tL92fh7zeIo3hVtCzhnpw8IeWAM-KcI419cIkm/pub?w=745&h=310)

# Machine Learning and AI with Go

These materials are for the 2019 "Machine learning and AI with Go" workshop at GopherCon. The workshop provides an intensive and idiomatic view on training, utilizing, evaluating, and deploying machine learning and AI models using Go. This class is perfect for gophers wanting an approachable, code-first introduction to ML/AI.

- [Slides from the workshop](https://docs.google.com/presentation/d/1igJntH89r0qT3BhD-91AewOKz9CZ9FWfOmmicxino7k/edit?usp=sharing)
- Instructors 
  - Daniel Whitenack - [website/blog](http://www.datadan.io/), [twitter](https://twitter.com/dwhitena), [github](https://github.com/dwhitena)
  - Miriah Peterson - [twitter](https://twitter.com/captainnobody1), [github](https://github.com/Soypete)
- During the workshop, you will also need to work a bit at the command line. If you are new to the command line or need a refresher, look through [this guide](http://bit.ly/2SytJAR) or [this quick tutorial](https://lifehacker.com/5633909/who-needs-a-mouse-learn-to-use-the-command-line-for-almost-anything).

*Note: This material has been designed to be taught in a classroom environment at GopherCon. The code is well commented but missing some of the contextual concepts and ideas that will be covered in class.*

## Agenda

9:00-9:30 Introductions, Logistics 🎤   
9:30-10:45 [Introduction to ML/AI](#introduction-to-machine-learning-and-ai-) 🧠   
10:45-11:00 Break ☕   
11:00-12:00 [Linear and Logistic Regression](#linear-regression-) 📈   
12:00-1:15 Lunch 🍕   
1:15-2:45 [Neural Networks, Deep Learning](#neural-networks-) 🤖   
2:45-3:00 Break ☕   
3:00-4:30 [Hands on](#hands-on-%EF%B8%8F) ⌨️   
4:30-5:00 Next steps, Conclusions 💥   

## Introduction to Machine Learning and AI 🧠

This session introduces the basics of Machine Learning and AI along with a whole bunch of related jargon. Are ML and AI the same thing or are they different? What differentiates them from traditional software engineering? How does a machine learn anyway? We will answer these questions and many more, such that you can communicate well with data science teams and understand the steps involved in most machine learning workflows.

[See slides](https://docs.google.com/presentation/d/1igJntH89r0qT3BhD-91AewOKz9CZ9FWfOmmicxino7k/edit?usp=sharing)

## Linear Regression 📈

There are a couple fundamental AI building blocks that show up all over AI applications. These are _linear regression_ and _gradient descent_. In this session, we will explore these building blocks by implementing them from scratch in Go. We will then utilize the `github.com/sajari/regression` package to extend our linear regression model to a multiple regression model.

Linear regression code examples:

- [Example 1: Profiling advertising/sales data](linear_regression/example1/example1.go)
- [Example 2: Splitting the data into training/test](linear_regression/example2/example2.go)
- [Example 3: Training a linear regression model, SGD](linear_regression/example3/example3.go)
- [Example 4: Training a single regression model w/ github.com/sajari/regression](linear_regression/example4/example4.go)
- [Example 5: Training a multiple regression model w/ github.com/sajari/regression](linear_regression/example5/example5.go)
- [Example 6: Evaluating regression models](linear_regression/example6/example6.go)

## Logistic Regression 📈

Linear regression is great, but there's only so much you can do with it. It is, after all, limited to modeling rather simple relationships where one thing is proportional to changes in another thing. Not all relationships are like that. In particular, classification problems require us to model discrete states or categories, not continuous values. In this session, we will move from linear regression to _logistic regression_ such that we can model non-linear relationships.

Logistic regression code examples:

- [Example 1: Plotting a logistic function](logistic_regression/example1/example1.go)
- [Example 2: "Clean" loan data](logistic_regression/example2/example2.go)
- [Example 3: Profile the loan data](logistic_regression/example3/example3.go)
- [Example 4: Split the data into training/test](logistic_regression/example4/example4.go)
- [Example 5: Train a logistic regression model](logistic_regression/example5/example5.go)
- [Example 6: Evaluate the logistic regression model](logistic_regression/example6/example6.go)
- [Example 6: Train a logistic regression model w/ github.com/cdipaolo/goml](logistic_regression/example7/example7.go)

## Neural Networks 🤖

Ok, in this session we finally get to the thing everyone is interested in... _neural nets_! Despite the intimidation that normally surrounds these seemingly magical, brain-like things, they aren't that much more complicated than our previous regression models. In this session, we will implement a simple neural network from scratch with Go to understand the fundamentals.

Neural Network code examples:

- [Example 1: Building a simple neural network](neural_networks/example1/example1.go)
- [Example 2: Utilizing the simple neural network for classification](neural_networks/example2/example2.go)

## Deep Learning 🤖

In this final teaching session, the concepts get very deep. We will take our basic neural network and figure out how we could modify it to be a _deep learning_ neural network (because, after we know about deep learning, we can all ask for a raise at our company). These super powerful networks can do amazing things, and it's not that difficult to leverage them in your Go applications!

Deep Learning code examples:

- [Example 1: Using the TensorFlow Go bindings for inference](deep_learning/example1/example1.go), Setup required:
  - Download the Tensorflow model from [here](http://download.tensorflow.org/models/object_detection/ssd_mobilenet_v1_coco_2018_01_28.tar.gz)
  - Unzip the model files
- [Example 2: Using GoCV to analyze video](deep_learning/example2/example2.go), Setup required:
  - [Install GoCV](https://gocv.io/getting-started/)
  - Download the TensorFlow Inception model from [here](https://storage.googleapis.com/download.tensorflow.org/models/inception5h.zip)
  - Unzip the Inception model files.
- [Example 3: Sentiment analysis with MachineBox](deep_learning/example3/example3.go), Setup required
  - Run MachineBox's [Textbox](https://docs.veritone.com/#/developer/machine-box/boxes/textbox)

## Hands on ⌨️ 

Based on feedback from previous GopherCon workshops, this iteration will include a couple hands-on options for the final session:

1. Try your hand at training some of your own ML/AI models using publicly available data. [Miriah Peterson](https://github.com/Soypete) was nice enough to compile some data sets where you will be able flex your regression, classification, and neural net muscles. [See here for more information and instructions on getting started](projects.md).
2. Learn how infrastructure projects written in Go (Docker, Kubernetes, Pachyderm, and Minio) can be used to deploy production ML/AI model training and inference. If you choose this option, you will follow a tutorial that will walk you through deploying a production ML pipeline on top of Kubernetes. [See here for more information and instructions on getting started](https://github.com/dwhitena/pach-go-regression). 

The instructors and TAs will be available for questions during this time, but you can also follow up with the instructor during the rest of GopherCon or on Gophers Slack.



___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
