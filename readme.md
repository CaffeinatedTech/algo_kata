## Algorithm Kata
This tool will help you to practice your algorithms in different languages.
You will be given some data - usually an array of numbers, then you will be asked
to use a specific algorithm on it, and supply the result.  You will be timed.

### Configuration
You are able to modify the algorithms, and the languages which will be asked
by editing the config.toml file.
You may also supply a list of languages from the command line with the
-l -languages parameter which takes a comma separated list

`algo_kata -l Javascript,Go`

This will limit the languages to Javascript, and Go.

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
* Download the binary appropriate to your machine from the releases page.
* Run it

**Build Method:**  
Make sure you already have go v1.22 installed and working. `go version`  
* Clone the repo with `git clone https://github.com/CaffeinatedTech/algo_kata.git`
* Build the binary with `go build .`
* Use, copy, move your shiny new binary `algo_kata`
