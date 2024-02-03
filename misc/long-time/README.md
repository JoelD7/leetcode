# Problem
This is a problem that I saw online. 

Write a function that receives an integer `seconds` and returns a string representation of those seconds in a long-time format. 

## Examples
- **Input**: 180
- **Output**: 3 minutes
---
- **Input**: 4896
- **Output**: 1 hour, 21 minutes, and 36 seconds
---
- **Input**: 16162572
- **Output**: 187 days, 1 hour, 38 minutes and 12 seconds
---
- **Input**: 32445788
- **Output**: 1 year, 234 days, 18 hours, 43 minutes and 8 seconds
---
- **Input**: 5
- **Output**: 5 seconds
---
- **Input**: 90061
- **Output**: 1 day, 1 hour, 1 minute and 1 second
- ---
- **Input**: 0
- **Output**: 0 seconds

## Caveats
- Note that the time unit should be plural only if it is greater than 1.
- Time units are separated by comma and space, except by the last one, which is separated by "and".
- The biggest time unit `seconds` will get to is 1 year.  
- `1 <= seconds <= 63071999`
  - 63071999 seconds = 1 year, 364 days, 23 hours, 59 minutes and 59 seconds

# Solution
The solution parts of the following principle: the division between `timeU` and `modOp`, where `timeU` is a numeric representation of a time unit `x`, and `modOp` is the amount of `timeU`'s required to make 1 of the next bigger time unit `y`, produces a `quotient` and a `remainder`. `quotient` is the amount of time units `y`, and `remainder` is the amount of time units `x`. When combined, they both represent the same amount of time that `timeU`, but expressed in combined time units. For example, let's say we have the following values:
    
    timeU = 182
    modOp = 60
    x = seconds
    y = minutes
then, 
    
    timeU/modOp = quotient and remainder
    remainder = 182 mod 60
    remainder = 2
    quotient = 180 / 60
    quotient = 3

    time = 3 minutes and 2 seconds = 182 seconds
Note that: 
- the `quotient` is the result of the division **without** the remainder
- `modOp = 60` because there are 60x in 1y, or, 60 seconds in 1 minute 

----
In the code, I applied this principle by first using several modulus that will change depending on the time unit we are converting. Since the input of the code is seconds, we start with seconds and increase from there: seconds, minutes, hours, days and years. 

We first need to use a loop to build the time units contained withing `seconds`. On each iteration, the value of `timeU` is updated to obtain the next bigger time unit, excluding the remainder, that is kept in an array. When we reach the final element of `modulus` or `modOp` > `timeU`(which means `timeU` has reached it's largest time unit), then the loop is done.

After that with iterate from back to front(the built array is on ascending order) and append the text part to each time unit with the help of an auxiliary array `longUnits`