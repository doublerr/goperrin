#goperrin

[![Circle CI](https://circleci.com/gh/doublerr/goperrin.svg?style=svg)](https://circleci.com/gh/doublerr/goperrin)

This is a simple golang implementation of the perrin sequence, http://mathworld.wolfram.com/PerrinSequence.html, which can be a good backoff sequence for polling remote APIs.

There are two possibly useful functions to reset and cap the sequence included.  See examples for more.

##Documentation
https://godoc.org/github.com/doublerr/goperrin

## Examples

#### Continiously increment

```go
import "github.com/doublerr/goperrin"

func main {
        p := goperrin.Perrin()
        for i := 0; i < 100; i += 1) {
                fmt.Println(p())
        }
}

./goperrin

3
0
2
3
2
5
5
7
10
…
```

#### Cap the sequence at `max`:

```go
import "github.com/doublerr/goperrin"

func main {
        p := goperrin.PerrinMax(5)
        for i := 0; i < 100; i += 1 {
                fmt.Println(p())
        }
}

./goperrin

3
0
2
3
2
5
5
5
5
…
```

#### Reset the sequence at `max`:

```go
import "github.com/doublerr/goperrin"

func main {
        p := goperrin.PerrinReset(5)
        for i := 0; i < 100; i += 1) {
                fmt.Println(p())
        }
}

./goperrin

3
0
2
3
2
5
3
0
2
…
```

## Contributing

## Authors

* Ryan Richard <ryanrichard07@gmail.com>
