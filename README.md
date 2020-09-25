# Buycoins Research Assistant Application

Below are answers to the questions required for applying for a research assistant role at buycoins.

## Finance

- I explained the Bitcoin stock-to-flow model and why it is a bad model [here](finance/S2F.md)

- Workings for the Black-Scholes call price for Yara Inc is shown [here](finance/bsm.pdf)

## Computer Science

- [Here](cs/fibonacci.md), i explained why it is a bad idea to use recursion method to find the fibonacci of a number.

- [Here](cs/proth.go), contains the code for a function that takes in a Proth number and determines if it is prime using Proth's theorem.

- The test file for the code is [here](cs/proth_test.go)

    To run the program, you need to have the go programming language installed.

        - clone this repository,
        - navigate to the folder containing proth.go i.e cs/
        - run the program with a number as an argument  

Sample outputs

    ```
    $ go run proth.go 4
    4 is not a Proth number, 4 is not a Proth prime

    $ go run proth.go 9
    9 is a Proth number, 9 is not a Proth prime

    $ go run proth.go 41
    41 is a Proth number, 41 is a Proth prime
    ```

## Mathematics

- [Here](mathematics/math.pdf), is the solution to the question:

`Over all real numbers, find the minimum value of a positive real number, y such that
y = sqrt((x+6)^2 + 25) + sqrt((x-6)^2 + 121)`