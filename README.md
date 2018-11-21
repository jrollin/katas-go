# TDD Code katas in the Go programming language

Includes following problems:

 * [FizzBuzz](http://codingdojo.org/kata/FizzBuzz/)

## Dependencies
All katas were developed with test driven development and thus make heavy
usage of the packages
  * [github.com/stretchr/testify](https://github.com/stretchr/testify)  
   install : go get github.com/stretchr/testify
  * testing (builtin)

watch lib :
  * [fresh](https://github.com/gravityblast/fresh)   
   Install: go get github.com/pilu/fresh 

## Running
Tests are run with
```bash
go test
```
in the project directory.

For example you may run the `fizzbuzz` kata tests with
```bash
go get github.com/jrollin/kataks-go/fizzbuzz
cd $GOPATH/src/github.com/jrollin/kataks-go/fizzbuzz
go test
```