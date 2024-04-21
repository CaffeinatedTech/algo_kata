## Algorithm Kata
This tool will help you to practice your algorithms in different languages.
You will be given some data - usually an array of numbers, then you will be asked
to use a specific algorithm on it, and supply the result.  You will be timed.

You are not expected to supply the full code solution, just the output as either
a number, or a json encoded array `[1,2,3,4]`.

### Configuration
You are able to modify the algorithms, and the languages which will be asked
by editing the config.toml file.

When you run the app, you will be asked which of the supplied languages, and
algorithms that you wish to be included in your session.

### Supported Datastructures
Currently algo_kata supports one dimensional arrays of integers,
so we are limited to testing search and sort algorithms on those.

**Example algorithms:**  
* Binary search
* Linear search
* Quick sort
* Bubble sort

### Quickstart
You can either grab a binary from the releases page, or clone and build the repo.

**Binary Method:**  
* Download the binary appropriate to your machine from the [releases page](https://github.com/CaffeinatedTech/algo_kata/releases/latest).
* Run it

**Build Method:**  
Make sure you already have go v1.22 installed and working. `go version`  
* Clone the repo with `git clone https://github.com/CaffeinatedTech/algo_kata.git`
* Change into directory `cd algo_kata`
* Install dependencies `go install .`
* Build the binary with `go build .`
* Use, copy, move your shiny new binary `algo_kata` and keep the config.toml file with it.
