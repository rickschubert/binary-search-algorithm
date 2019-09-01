Attempting the binary search problem
===================================

This package represents my own take at the famous [binary search problem](http://codekata.com/kata/kata02-karate-chop/). The problem is: Given that we have a *sorted* list of integers, what is the fastest way to find the index of an particular element in that list? I.e. we have a list from 1 to 10 `[1, 2, 3, 4, 5, 6, 7, 8, 9, 10]`, how do we find the index of the number 4 (result should be `3`)?

## The approach
Let's say we are looking for `4` in a list from `1` to `10`. Since the list is sorted, we can break the list into two halfs. We then compare our value in question (`4`) against the last element in the first half (`5`): If the the value is smaller or equal to the last element of the first half, we know that the desired element must be in the first half. Otherwise it must be in the second half.

By halfing the list every time, we can check even big lists efficiently. If we would look for the element 1000 in a list from 1 to 1000 and would test every single element, we would take 1000 steps/checks - using my method, we only need 10!

## Difficulties
Apart from coming up with the halfing solution, I found it difficult to keep track of the index (since we don't just want to find the element in question but also its index in the list). Furthermore, my initial solution mutated the original list, which I found out and remedied during unit testing.

I created this solution completely on my own and found it **very** challenging, since I am far from a math genius. Even prouder am I of the result!

## Steps to run unit tests
- `go get`
- `go test`
