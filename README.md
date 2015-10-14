#goperrin

This is a simple golang implementation of the perrin sequence (http://mathworld.wolfram.com/PerrinSequence.html) which can be a good backoff sequence for polling remote APIs.

There are two possibly useful functions to reset and cap the sequence included.  See examples for more.

## Examples

In all the examples, I set the initial value of the for loop to `3`. I do this since the initial value of the sequence is `3` and I haven't figured out a better way to get the first return to be `3`. Let me know if you have one!

#### Continiously increment 
 * According to the perrin sequence (The for loop will stop once `i` has passed `10000`):

```go
import "github.com/doublerr/goperrin"

func main {
        p := goperrin.Perrin()
        for i := 3; i < 10000; i = p() {
                fmt.Println(i)
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
        for i := 3; i < 10000; i = p() {
                fmt.Println(i)
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
        for i := 3; i < 10000; i = p() {
                fmt.Println(i)
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
