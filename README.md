# gosci

`gosci` is a Go library providing functions for various scientific calculations, starting with geometry.

## Installation

```bash
go get github.com/KengoWada/gosci
```

## Usage

```go
package main

import (
    "fmt"

    "github.com/KengoWada/gosci/geometry"
)

func main() {
    square := geometry.Square{Side: 5}
    area, err := square.Area()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Square Area:", area)
}
```

## Contributing

Contributions are welcome! Please feel free to submit pull requests or open issues to suggest improvements or report bugs.
