# Gong

![image](assets/29.png)
![image](assets/gong.png)
![image](assets/27.png)

Gong helps you scheduling intervals and timeouts in an easy way, without having to worry about channels, for and select.

## Usage

To schedule a function that will be executed at a fixed interval you just need to provide the callback and the duration to the `gong.Interval` func

```go
// execute SayHello every 2 seconds
gong.Interval(ctx, SayHello, 2*time.Second)
```

In the same fashion you can schedule a function to be executed after a fixed time providing the callback and the duration to the `gong.Timeout` func

```go
// execute SayHello after a second
gong.Timeout(ctx, SayHello, time.Second)
```

These functions are non-blocking, so your program will need to wait to see the execution of the callback.

### Cancellation

To cancel a running interval or timeout you have to use an appropriate Context.
For example, to cancel an interval after a certain amount of time you can use a context.WithTimeout.

This example will run the callback for just 3 seconds, then the interval will be cancelled:

```go
func main() {
    ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
    defer cancel()

    gong.Interval(ctx, func() {
        log.Println("hello")
    }, time.Second)

    time.Sleep(time.Second * 8)
}
```

### Providing arguments and get returning values

If your callback takes any arguments you can use one of the `IntervalX`/`TimeoutX` func, or wrap it in a closure.
These functions give you type safety, keeping your code clean.

```go
func main() {
    gong.Interval1(context.Background(), SayHello, time.Second, "John")

    gong.Interval(context.Background(), func() {
        SayHello("John")
    }, time.Second)

    time.Sleep(time.Second * 8)
}

func SayHello(name string) {
    log.Println("hello", name)
}
```

If you need to get some values from the callback you can provide a channel where to push the results

```go
func main() {
    out := make(chan int)
    gong.Interval(context.Background(), RollDice(out), time.Second)

    go func() {
        for i := range out {
            log.Printf("Got %d\n", i)
        }
    }()

    time.Sleep(time.Second * 8)
}

func RollDice(out chan int) func() {
    return func() {
        out <- rand.Intn(6) + 1
    }
}
```
