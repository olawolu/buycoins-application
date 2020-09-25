# Recursion and Fibonacci

```go
func fib(n int) int {
    // recursion
    if n <= 1 {
        return n
    }
    return fib(n-1) + fib(n-2)
}
```

Let the function T(n) denote the number of elementary operations performed by calling fib(n).

- for n>1 the function will perform a fixed number of operations $k_2$, it will make recursive calls to fib(n-1) and fib(n-2). These recursive calls will perform T(n-1) and T(n-2) operations respectively. In total, we get T(n)=T(n-1) + T(n-2) + O(1).

Solving the above recursive equation, we get the upper bound as O(2^n).

Therefore, the time complexity: T(n) = T(n-1) + T(n-2) is exponetial.
