# Problem
https://leetcode.com/problems/car-fleet/

There are n cars going to the same destination along a one-lane road. The destination is target miles away.

You are given two integer array position and speed, both of length n, where position[i] is the position of the ith car and speed[i] is the speed of the ith car (in miles per hour).

A car can never pass another car ahead of it, but it can catch up to it and drive bumper to bumper at the same speed. The faster car will slow down to match the slower car's speed. The distance between these two cars is ignored (i.e., they are assumed to have the same position).

A car fleet is some non-empty set of cars driving at the same position and same speed. Note that a single car is also a car fleet.

If a car catches up to a car fleet right at the destination point, it will still be considered as one car fleet.

Return the number of car fleets that will arrive at the destination.



### Example 1:

    Input: target = 12, position = [10,8,0,5,3], speed = [2,4,1,1,3]
    Output: 3
    Explanation:
    The cars starting at 10 (speed 2) and 8 (speed 4) become a fleet, meeting each other at 12.
    The car starting at 0 does not catch up to any other car, so it is a fleet by itself.
    The cars starting at 5 (speed 1) and 3 (speed 3) become a fleet, meeting each other at 6. The fleet moves at speed 1 until it reaches target.
    Note that no other cars meet these fleets before the destination, so the answer is 3.

### Example 2:
    
    Input: target = 10, position = [3], speed = [3]
    Output: 1
    Explanation: There is only one car, hence there is only one fleet.

### Example 3:
    
    Input: target = 100, position = [0,2,4], speed = [4,2,1]
    Output: 1
    Explanation:
    The cars starting at 0 (speed 4) and 2 (speed 2) become a fleet, meeting each other at 4. The fleet moves at speed 2.
    Then, the fleet (speed 2) and the car starting at 4 (speed 1) become one fleet, meeting each other at 6. The fleet moves at speed 1 until it reaches target.



### Constraints:

    n == position.length == speed.length
    1 <= n <= 105
    0 < target <= 106
    0 <= position[i] < target
    All the values of position are unique.
    0 < speed[i] <= 106

# Solution
### Intuition

By calculating the time that each car fleet(on its own) will take to reach target, we can know whether a car fleet will meet with another.

One key into understanding why this approach works is inside the problem’s constraints which states: “*A car cannot pass another car, but it can catch up and then travel next to it at the speed of the slower car.*” In others words, a car X that starts behind a car Y, cannot pass car Y even if it is faster than Y. What will happen though is that car X ***may*** catch up car Y and travel at the same speed of Y.

**How can we know if the faster car will catch up to a slower car?** Basic physics, i.e., calculating the time it takes for each car to reach `target`. If a car X takes less time to reach `target` than a car Y starting directly ahead of it, X will catch up to Y and form a car fleet. Conversely, if a car takes more time to reach the `target` than the car ahead, it won't catch up before the destination and will therefore form a new car fleet separate from the one ahead.

### Algorithm

1. **Sort cars/fleets by position**. The described approach works if we sort cars by position so we can easily get the car that is directly ahead of another.
    1. The cars should be sorted in descending order because we’ll be making decisions using the car that is “ahead” of another
    2. Here we use a matrix to store both the position and speeds of each car.
2. **Use a stack to keep track of the fleets**. The values of the stack will be the times taken to reach `target`
    1. A stack is ideal for this problem because it allows O(1) to the last added element which is precisely what we need. We don’t need to access elements other than the last one.
3. **Iterate over the sorted cars matrix**.
    1. Calculate the time taken to reach `target` for the car in the current iteration.
        1. `timeTaken` should be a float to maintain accurate time tracking. If we have that variable as an int, a time of 1.455666 would be bundled together with a time of 1 because of how rounding works with intergers and this is not accurate.
    2. If the time taken of the current car to reach target is greater than the car in the previous iteration(i.e., the car that is directly ahead because we sorted in desc order), then there is no way for the current car to reach the previous one, so we push that time into the stack to signal **a separate fleet** will be formed
    3. If the time is equal, then we don’t push into the stack because this means that this and the previous car will form a fleet that will reach `target` at the same time
    4. If the time is less, this means that the current car is faster than the previous one(the car that started directly ahead) and it will catch up to it to form a fleet, so we don’t push into the stack. Remember that faster cars cannot pass slower ones.
4. **Return the stack size**. Since we pushed to the stack only when a new car fleet was formed, the stack size represents the total fleets.