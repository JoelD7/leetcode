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