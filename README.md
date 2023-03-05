Details : https://youtu.be/xjr__JqZ-hQ


<p align="left"> <img src="https://komarev.com/ghpvc/?username=golangast&label=Profile%20views&color=0e75b6&style=flat" alt="golangast" /> </p>


![GitHub repo file count](https://img.shields.io/github/directory-file-count/golangast/contribute) 
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/golangast/contribute)
![GitHub repo size](https://img.shields.io/github/repo-size/golangast/contribute)
![GitHub commit activity](https://img.shields.io/github/commit-activity/w/golangast/contribute)
![Go 100%](https://img.shields.io/badge/Go-100%25-blue)


## Contribute
* [General info](#general-info)
* [Why build this?](#why-build-this)
* [Technologies](#technologies)
* [Setup](#setup)
* [Repository overview](#repository-overview)
* [Guidelines](#guidelines)
* [Special thanks](#special-thanks)

## License is GNU Version 3 
To ensure people can share but need to keep original code for the public use. This way your are still recognized for your
contributions.

## General info
This is just a beginner project to contribute to and play with.  
We all need to contribute more.  That is the point of the internet and open source.


## Why build this?
* It would be nice if people got involved.
* A community creates inspiration.
* People need to contribute to things.
* I need to contribute more.


## Technologies
Project is created with:
* [Cobra](https://github.com/spf13/cobra) - To create separate commands


## Setup
To run this project for development, download it and run the following
```
go mod init "yourproject"

go mod tidy

go install github.com/spf13/cobra-cli@latest

export PATH="~/go/bin:$PATH" //if you dont add it to your bashrc file then its just temparory anyways

cobra-cli init
```

to create command

```
cobra add your command name 

```

example: cobra add second


to run it 

```
go run main.go your command name

```

example: go run main.go first

## Repository overview

```
├── cmd
└── main.go
```

- Resources that I use [Resources](https://docs.google.com/document/d/1Zb9GCWPKeEJ4Dyn2TkT-O3wJ8AFc-IMxZzTugNCjr-8/edit?usp=sharing)


## Special thanks
Your name will be here

## Guidelines
- Please try to make the code readable
- Please try to make it verbose so everyone can learn
- Please no bash that downloads other programs
- Please no make files
- Have fun and dont worry about details or bad code or anything.  Just be creative if you want.
