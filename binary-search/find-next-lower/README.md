This is a more efficient way to find the next lower element using binary search. For contrast, here is how we would do it using linear search: 

```go
func searchNextLower(timestamp int, arr []int) int {
    for i := len(arr) - 1; i >= 0; i-- {
        if arr[i] < timestamp {
            return arr[i]
        }
    }

    return -1
}

```