# name

[![Test](https://github.com/frederikhs/name/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/frederikhs/name/actions/workflows/test.yml)

Very, very simple module for getting reproducible random animal names

### about
Each function generates a reproducible random animal name prefixed with an adjective.

In short this means that given the same input, the method will return the same output, while the output being random.


### usage

```go
package main

import (
	"fmt"
	"github.com/frederikhs/name"
)

func main()  {
	pascal := name.GeneratePascalName("one")   // EasySpider
	camel := name.GenerateCamelCaseName("one") // easySpider
	fmt.Println(pascal, camel)
}
```
