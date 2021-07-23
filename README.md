# numeral

![golang][golang-badge]
![go-mod][go-mod-badge]
![size][size-badge]

A numeral utility package for Go

## Installation

To install the **numeral** library run:

```sh
go get github.com/suft/numeral
```

## Usage

### Ordinal

``numeral.Ordinal`` is a function that takes a unsigned integer and returns the English ordinal letters following the inputted integer.

```go
package main

import (
    "fmt"

    "github.com/suft/numeral"
)

func main() {
    visitor := numeral.Ordinal(3)
    fmt.Printf("You are the %s visitor", visitor)
}
```

```text
You are the 3rd visitor
```

## Issues

If you encounter any problems, please [file an issue][new-issue] along with a
detailed description.

## Contributing

Contributions are welcome, and they are greatly appreciated! Every little bit helps.

## License

![mit][mit-badge]

Distributed under the terms of the [MIT license][mit], numeral is free and
open source software.

[golang-badge]: https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white
[go-mod-badge]: https://img.shields.io/github/go-mod/go-version/suft/numeral?style=for-the-badge
[size-badge]: https://img.shields.io/github/repo-size/suft/numeral?style=for-the-badge
[mit]: ./LICENSE
[mit-badge]: https://img.shields.io/github/license/suft/numeral?style=for-the-badge
[new-issue]: https://github.com/suft/numeral/issues/new
