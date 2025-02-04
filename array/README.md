# Array

Arrays are a basic and essential data structure in computer science. They consist of a fixed-size contiguous block of memory, and offer O(1) read and write time complexity. As a fundamental element of programming languages, arrays come built-in as part of their core.

To provide a real-world analogy, consider an array of athletes preparing for a sprinting match. Each athlete occupies a specific position within the array, which is typically denoted as 1, 2,..., n. While it is technically possible for each athlete to be in a different position, the positions generally carry some form of significance, such as alphabetical order or seniority within the sport.

## Implementation

In the Go programming language, arrays are considered values rather than pointers, and represent the entirety of the array. Whenever an array is passed to a function, a copy of the array is created, resulting in additional memory usage. However, to avoid this issue, it is possible to pass a pointer to the array instead.

In order to define an array in Go, it is necessary to specify the size of the array using a constant. By using constants in this manner, it is no longer necessary to utilize the make function in order to create the array.

```Go
package main

import "fmt"

func main() {
	var nums1 [2]int
	nums2 := [3]int{1, 2, 3}
	fmt.Println(nums1, nums2) // Prints [0 0] [1 2 3]
}
```

Although arrays are fundamental data structures in Go, their constant size can make them inflexible and difficult to use in situations where a variable size is required. To address this issue, Go provides [slices](https://blog.golang.org/slices-intro) which are an abstraction of arrays that offer more convenient access to sequential data typically stored in arrays.

Slices enable the addition of values using the `append` function, which allows for dynamic resizing of the slice. Additionally, selectors of the format [low:high] can be used to select or manipulate data in the slice. By utilizing slices instead of arrays, Go programmers gain a more flexible and powerful tool to manage their data structures.

```Go
package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	nums = append([]int{0}, nums...) // Add new element to the start
	nums = append(nums, 4)           // Add new element to the end
	nums = nums[:len(nums)-1]        // Removes last element
	fmt.Println(nums)                // Prints [0 1 2 3]
}
```

The [make](https://golang.org/pkg/builtin/#make) function can be used to create a zeroed slice of a given length and capacity.

```Go
package main

import "fmt"

func main() {
	nums := make([]int, 2)
	nums[0], nums[1] = 0, 1
	fmt.Println(nums, len(nums), cap(nums)) // Prints [0 1] 2 2
}
```

Slice expressions in the form of input[low:high] can be used to manipulate slices or access their elements

```Go
package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	nums = nums[:len(nums)-1] // drop last element
	nums = nums[1:]           // drop first element
	nums = nums[1:]           // all elements from index 1 to the end
	nums = nums[:2]           // all elements from index 0 to 2
	nums = nums[1:2]          // the element from index 1 to 2
	fmt.Println(nums)         // Prints [4]
}
```

## Complexity

In computer science, the act of accessing an element within an array using an index i has an O(1) time complexity. This means that regardless of the size of the array, the read and write operations for a given element can be performed in constant time.

While arrays are useful for certain tasks, searching an unsorted array can be a time-consuming O(n) operation. Since the target item could be located anywhere in the array, every element must be checked until the item is found. Due to this limitation, alternative data structures such as trees and hash tables are often more suitable for search operations.

Both addition and deletion operations on arrays can be O(n) operations in Arrays. The process of removing an element can create an empty slot that must be eliminated by shifting the remaining items. Similarly, adding items to an array may require shifting existing items in order to create space for the new item. These inefficiencies can make alternative data structures such as [trees](../tree) or [hash tables](../hashtable) more suitable for managing operations involving additions and deletions.

## Application

Arrays are used wherever sequential data or more than one piece of data is needed. The fast read and write access to a given element makes arrays suitable for implementing other data structures such as [strings](../strings), [stacks](../stack), [queues](../queue) and [hash tables](../hashtable).

## Rehearsal

### Reverse Array In-place

Given an array of integers, a start index, and an end index, reverse the integers in the array in-place without using any extra memory. [Solution](reverse_inplace.go), [Tests](reverse_inplace_test.go)

### Add Numbers

Given two numbers as an array like [2,9] and [9,9,9] return the sum of the numbers they represent like [1,0,2,8], because 29+999=1028. [Solution](add_two_numbers.go), [Tests](add_two_numbers_test.go)

### Find Duplicate in Array

Given an array of integers (1,2,...,n), find a duplicate number in O(n) time. [Solution](find_duplicate_in_array.go), [Tests](find_duplicate_in_array_test.go)

### Zero Sum Triplets

Given an array of numbers like `{1, 2, -4, 6, 3}` returns unique triplets from the numbers with sum that equals zero like `{-4, 1, 3}`. [Solution](zero_sum_triplets.go), [Tests](zero_sum_triplets_test.go)

### Product of All Other Elements

Given an array of integers A, construct a new array B such that B[i] = product of all items in A except A[i] without using division in O(n) time. [Solution](product_of_all_other_elements.go), [Tests](product_of_all_other_elements_test.go)

### Equal Sum Sub-arrays

Given an array of integers A, return two sub-arrays with equal sums without changing the order of elements. . [Solution](equal_sum_subarrays.go), [Tests](equal_sum_subarrays_test.go)

### Rotate K Steps

Given an array of integers and a number k, rotate the array k times. [Solution](rotate_k_steps.go), [Tests](rotate_k_steps_test.go)
