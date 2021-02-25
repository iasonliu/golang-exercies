https://www.youtube.com/watch?v=f6kdp27TYZs
https://talks.golang.org/2012/concurrency.slide

# Channels
A channel in Go provides a connection between two goroutines, allowing them to communicate.

    // Declaring and initializing.
    var c chan int
    c = make(chan int)
    // or
    c := make(chan int)
    // Sending on a channel.
    c <- 1
    // Receiving from a channel.
    // The "arrow" indicates the direction of data flow.
    value = <-c