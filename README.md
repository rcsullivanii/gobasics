# Go Basics

## Race Conditions
Do you find your code showing race conditions? Try it with larger p values to get many concurrent processes. What can you do about that? Why or why not are you seeing race conditions?

While researching, I learned that Goroutines run in the same address space, so access to shared memory must be synchronized.

The first program avoids the race conditions as each "recursive" call passes its own copy of the result. That is, the first program contains no shared data.

The second program's usage of the global variable res means that res is updated concurrently by multiple gorountines (unsychronized access) which does mean there are race conditions. 

## Assignment Description
Write 2 Go programs, both of which compute N ^ P, N raised to the power P. Both will use goroutines (concurrent processes). Have the main function define the variable n and p as integers, and for simplicity, these parameters to the problem will be set as constants in the code rather than worrying for now about getting input from the kayboard. So set them to some initial value in the code for main (e.g., n = 127 and p = 12). Have function main call power(n,p). Then, the programs differ as follows:

Prog One
Use the processes example from class as a model, and have the result of power(n,P) be calculated by a "helper" three-parameter function ( call it power_h(n, p, res) ) that will be tail recursive. Do the exponentiation by repeated multiplication, with each invocation of power_h "recursing" on p-l and passing the increased accumulator . It's not really resursion, as we will make the "recursive" call a go routine. The final process spawned will see that it is p of 1 and then not "recurse" another go routine, but simply print the accumulated result.

This is, of course, not the best use for concurrency, as there is none. But it will help you become familiar with Go syntax.

In Prog Two we will do actual concurrency.
Prog Two
Do the same computation but this time, our function "power(n,p)" will spawn all n processes directly and have them executing concurrently. No need for the helper "power_h". Similar to Prog One, we can call the function "power(n,p)" from main, but we need only two parameters, as the results accumulator will be done with global data (global to the functions in the package).

So declare our results accumulator as a global variable (let's call it res) any initialize it appropriately.

Each goroutine will read the value of global "res" and multiply it by n.

As before in main we will declare n and p and set them to whatever values we wish to compute n^p for. We then call power(n,p) to make processes and compute the result. Make main sleep a few seconds after the call to power(n,p), to give the processes time to all complete. Once main awakens, have it print the result.

## Sources
https://go.dev/tour/concurrency/1