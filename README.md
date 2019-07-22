![Alt text](https://docs.google.com/drawings/d/e/2PACX-1vT37glyZXd8ViXedt0LCSpzsbWCUSSLhWuR3o5_74tL92fh7zeIo3hVtCzhnpw8IeWAM-KcI419cIkm/pub?w=745&h=310)

# Machine Learning and AI with Go

This is material for the 2019 "Machine learning and AI with Go" workshop at GopherCon. The workshop provides an intensive, comprehensive and idiomatic view on training, utilizing, evaluating, and deploying machine learning  and AI models using Go. We believe this class is perfect for anyone wishing to build data-driven applications that produce valuable insights, have reproducible behavior, and can be deployed within modern architectures.

- [Slides from the workshop](https://docs.google.com/presentation/d/1igJntH89r0qT3BhD-91AewOKz9CZ9FWfOmmicxino7k/edit?usp=sharing)
- Instructors 
  - Daniel Whitenack - [website/blog](http://www.datadan.io/), [twitter](https://twitter.com/dwhitena), [github](https://github.com/dwhitena)
  - Mariah Peterson - [twitter](https://twitter.com/captainnobody1), [github](https://github.com/Soypete)
- During the workshop, you will also need to work a bit at the command line. If you are new to the command line or need a refresher, look through [this quick tutorial](https://lifehacker.com/5633909/who-needs-a-mouse-learn-to-use-the-command-line-for-almost-anything).

*Note: This material has been designed to be taught in a classroom environment at GopherCon. The code is well commented but missing some of the contextual concepts and ideas that will be covered in class.*

## Introduction to Machine Learning and AI 

[See slides](https://docs.google.com/presentation/d/1igJntH89r0qT3BhD-91AewOKz9CZ9FWfOmmicxino7k/edit?usp=sharing)

## Linear and Logistic Regression

Linear regression code examples:

- [Example 1: Profiling advertising/sales data](linear_regression/example1/example1.go)
- [Example 2: Splitting the data into training/test](linear_regression/example2/example2.go)
- [Example 3: Training a linear regression model, SGD](linear_regression/example3/example3.go)
- [Example 4: Training a single regression model w/ github.com/sajari/regression](linear_regression/example4/example4.go)
- [Example 5: Training a multiple regression model w/ github.com/sajari/regression](linear_regression/example5/example5.go)
- [Example 6: Evaluating regression models](linear_regression/example6/example6.go)

Logistic regression code examples:

- [Example 1: Plotting a logistic function](logistic_regression/example1/example1.go)
- [Example 2: "Clean" loan data](logistic_regression/example2/example2.go)
- [Example 3: Profile the loan data](logistic_regression/example3/example3.go)
- [Example 4: Split the data into training/test](logistic_regression/example4/example4.go)
- [Example 5: Train a logistic regression model](logistic_regression/example5/example5.go)
- [Example 6: Evaluate the logistic regression model](logistic_regression/example6/example6.go)
- [Example 6: Train a logistic regression model w/ github.com/cdipaolo/goml](logistic_regression/example7/example7.go)

## Building a complete Go-based ML workflow

This material walks you through a lab in which you will implement a full ML workflow with Go, from data ingress to training to evaluation to inference. Once you are done with this material, you know how to implement and deploy the stages of the ML workflow in Go (for at least one type of ML model and data), and you will be able to transfer this workflow scaffolding to other problems.

[Building a complete Go-based ML workflow](ml_workflow)

___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
