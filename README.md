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

## UI of my CLI TOOL

Made using tview

### menu
![Screenshot from 2023-02-13 20-11-14](https://user-images.githubusercontent.com/73513838/218524962-814b86eb-5319-490d-8551-28b8cacb2694.png)
### generate a random joke
![Screenshot from 2023-02-13 20-08-08](https://user-images.githubusercontent.com/73513838/218522778-9b18d20b-a54b-4251-839a-a521dc29e94a.png)

![Screenshot from 2023-02-13 20-08-34](https://user-images.githubusercontent.com/73513838/218523222-226b1d35-4c04-4192-a736-ad429fbb0ae8.png)
### get n no of jokes
![Screenshot from 2023-02-13 20-10-23](https://user-images.githubusercontent.com/73513838/218523664-235e196c-53fb-4f59-8296-3197e02cd031.png)

![Screenshot from 2023-02-13 20-09-09](https://user-images.githubusercontent.com/73513838/218524208-801af20b-1aec-4331-9c1e-6c4ce334b9d4.png)

### quit
` Closes the app `