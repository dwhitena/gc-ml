![Alt text](https://docs.google.com/drawings/d/e/2PACX-1vT37glyZXd8ViXedt0LCSpzsbWCUSSLhWuR3o5_74tL92fh7zeIo3hVtCzhnpw8IeWAM-KcI419cIkm/pub?w=745&h=310)

# Machine Learning and AI with Go

This is material for the 2019 "Machine learning and AI with Go" workshop at GopherCon. The workshop provides an intensive, comprehensive and idiomatic view on training, utilizing, evaluating, and deploying machine learning  and AI models using Go. We believe this class is perfect for anyone wishing to build data-driven applications that produce valuable insights, have reproducible behavior, and can be deployed within modern architectures.

- [Slides from the workshop](https://docs.google.com/presentation/d/1igJntH89r0qT3BhD-91AewOKz9CZ9FWfOmmicxino7k/edit?usp=sharing)
- Instructors 
  - Daniel Whitenack - [website/blog](http://www.datadan.io/), [twitter](https://twitter.com/dwhitena), [github](https://github.com/dwhitena)
  - Mariah Peterson - [twitter](https://twitter.com/captainnobody1), [github](https://github.com/Soypete)
- During the workshop, you will also need to work a bit at the command line. If you are new to the command line or need a refresher, look through [this quick tutorial](https://lifehacker.com/5633909/who-needs-a-mouse-learn-to-use-the-command-line-for-almost-anything).

*Note: This material has been designed to be taught in a classroom environment at GopherCon. The code is well commented but missing some of the contextual concepts and ideas that will be covered in class.*

## Agenda

9:00-9:30 Introductions, Logistics 🎤   
9:30-10:30 [Introduction to ML/AI](#introduction-to-machine-learning-and-ai-) 🧠   
10:30-11:00 Break ☕   
11:00-12:00 [Linear and Logistic Regression](#linear-regression-) 📈   
12:00-1:00 Lunch 🍕   
1:00-2:00 [Neural Networks, Deep Learning](#neural-networks-) 🤖   
2:30-3:00 Break ☕   
3:00-4:30 [Hands on](#hands-on-%EF%B8%8F) ⌨️   
4:30-5:00 Next steps, Conclusions 💥   

## Introduction to Machine Learning and AI 🧠

[See slides](https://docs.google.com/presentation/d/1igJntH89r0qT3BhD-91AewOKz9CZ9FWfOmmicxino7k/edit?usp=sharing)

## Linear Regression 📈

Linear regression code examples:

- [Example 1: Profiling advertising/sales data](linear_regression/example1/example1.go)
- [Example 2: Splitting the data into training/test](linear_regression/example2/example2.go)
- [Example 3: Training a linear regression model, SGD](linear_regression/example3/example3.go)
- [Example 4: Training a single regression model w/ github.com/sajari/regression](linear_regression/example4/example4.go)
- [Example 5: Training a multiple regression model w/ github.com/sajari/regression](linear_regression/example5/example5.go)
- [Example 6: Evaluating regression models](linear_regression/example6/example6.go)

## Logistic Regression 📈

Logistic regression code examples:

- [Example 1: Plotting a logistic function](logistic_regression/example1/example1.go)
- [Example 2: "Clean" loan data](logistic_regression/example2/example2.go)
- [Example 3: Profile the loan data](logistic_regression/example3/example3.go)
- [Example 4: Split the data into training/test](logistic_regression/example4/example4.go)
- [Example 5: Train a logistic regression model](logistic_regression/example5/example5.go)
- [Example 6: Evaluate the logistic regression model](logistic_regression/example6/example6.go)
- [Example 6: Train a logistic regression model w/ github.com/cdipaolo/goml](logistic_regression/example7/example7.go)

## Neural Networks 🤖

Neural Network code examples:

- [Example 1: Building a simple neural network](neural_networks/example1/example1.go)
- [Example 2: Utilizing the simple neural network for classification](neural_networks/example2/example2.go)

## Deep Learning 🤖

Deep Learning code examples:

- [Example 1: Using the TensorFlow Go bindings for inference](deep_learning/example1/example1.go)
- [Example 2: Using GoCV to analyze video](deep_learning/example2/example2.go)
- [Example 3: Sentiment analysis with MachineBox](deep_learning/example3/example3.go)

## Hands on ⌨️ 

[See here for hands-on/ project ideas](projects.md)

___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
