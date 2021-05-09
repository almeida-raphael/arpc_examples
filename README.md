[![github](https://github.com/almeida-raphael/arpc_examples/workflows/Unit%20Tests/badge.svg)](https://github.com/almeida-raphael/arpc_examples)
[![codecov](https://codecov.io/gh/almeida-raphael/arpc_examples/branch/master/graph/badge.svg)](https://codecov.io/gh/almeida-raphael/arpc_examples)
# aRPC Examples
This is a repository for aRPC examples, [click here](https://github.com/almeida-raphael/arpc) if you want more details about aRCP.   

### Example List
* Greetings: This example makes an aRPC connection between one server and one client and make one aRPC call for the 
  greetings procedure. The greetings procedure receives a person data and return a greetings string contaning the 
  received person information.
* WordCount: This example makes an aRPC connection between one server and one client and make one aRPC call for the
  wordCount procedure. The wordCount procedure receives a text and returns the word frequency for the received text.

### Run
To run the examples you have to install [aRPC Code Generator](https://github.com/almeida-raphael/arpc_code_generator). 
Then clone this repository, install the dependencies with `go mod download` and then run 
`arpc_gen -input-path examples/arpc -packages-root-path examples`, the code generator will create the aRPC needed files. 
Now you can run any of the examples in the *cmd* folder.

### WIP
This project is currently a work in progress and should be not used in production environments.

### Authors
* [Raphael C. Almeida](https://github.com/almeida-raphael)
* [Vitor Vasconcellos](https://github.com/HeavenVolkoff)
* [Ericson Soares](https://github.com/fogodev)
