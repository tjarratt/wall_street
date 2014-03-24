wall_street
===========

[![Build Status](https://travis-ci.org/tjarratt/wall_street.svg?branch=master)](http://travis-ci.org/tjarratt/wall_street)
An implementation of libreadline in Go

![wolf](http://cdn.four-pins.com/assets/2013/12/wolf-of-wall-street-leonardo-dicaprio2.jpg)

Usage
-----

```go
package whatever

import (
  "github.com/tjarratt/wall_street"
)

func main() {
  input := wall_street.ReadLine("Are you a god?")
  if input == "yes" {
    println("oh okay. carry on then.")
  } else {
    println("then you shall die")
  }

  return
}
```

What's the deal with the name?
------------------------------

While building a CLI for my day job, I came to the realization that Go really needed a kickass implementation of readline in pure Go (without using CG0). This of course meant that I needed to think of a really kickass name. Just calling the package "readline" or "go-readline" would have been admitting defeat. We don't need anymore packages with "Go" in the name, and just naming something after its inspiration is boring.

So I thought for a bit and waited a few days. In a fever dream, I got to playing with words and thought that "read line" sounds a bit like "greed line". It's a small conceptual leap from "greed line" to "wall street". Well, it's a small leap when your brain is getting fried, at least.
