# JOKESFORALL

`A CLI tool that is created to make you giggle and laugh`

## DESCRIPTION

A CLI tool written in Golang and Cobra to generate random jokes. Grab a coffee and have fun!!!

## PREREQUISITE

The only prerequisite is to have `go1.19.5` and `cobra` installed in the local system!!


## GET STARTED 

### INSTALLATION
```
go install github.com/Dee-6777/jokesforall
```

### HOW TO USE THIS TOOL? 
* Get a random joke: `This command returns a random joke into your terminal.`
```
jokesforall 
```
* Get a random joke with any number of times: `This command returns a random joke any number of time you want to have.`
```
jokesforall <number>
```

### MAKEFILE

A makefile has been created to simplify the experience of the user. 

#### FUNCTIONALITY
* Build the binary file
```
make build
```
* Install
```
make install
```
* Run main.go
```
make run
```
* Test 
```
make test
```
* Get a random joke
```
make joke
```
* Get a random joke any number of times
```
make choose
```
* Build mod file
```
make mod
```
* Remove the binary file
```
make clean
```