# Data Structures and Algorithms

Repository for my practice of Data Structures and Algorithms

## Resources

- educative.io
- leetcode.com
- [Study Guide](https://www.reddit.com/r/cscareerquestions/comments/6luszf/a_leetcode_grinding_guide/?st=J92XK6ZK&sh=c2576763)

## Patterns to Remember
### ASCII

- '0' is 48
- 'A' is 65
- 'a' is 97

### Sliding Window

### Two Pointer

### Binary Search
#### Pattern I

```go
func binarySearch(nums []int, target int) int {
  if len(nums) == 0 {
    return -1
  }

  left := 0
  right := len(nums) - 1

  for left <= right {
    mid := left + (right - left) / 2
    if nums[mid] == target {
      return mid
    } else if nums[mid] < target {
      left = mid + 1
    } else {
      right = mid - 1
    }
  }

  return -1
}
```

Description:
- Most basic and elementary form of Binary Search
- For accessing a single index in the array.
- Search condition can be determined without comparing to the element's neighbors (or use specific elements around it)
- No post-processing required because at each step, you are checking to see if the element has been found. If you reach the end, then you know the element is not found.

Distinguishing syntax:
- Initial condition: `left = 0, right = length - 1`
- Termination `left > right`
- Searching left `right = mid - 1`
- Searching righ `left = mid + 1`

#### Pattern II

```go
func binarySearch(nums []int, target int) int {
  if len(nums) == 0 {
    return -1
  }

  left := 0
  right := len(nums) - 1

  for left < right {
    mid := left + (right - left) / 2
    if nums[mid] == target {
      return mid
    } else if nums[mid] < target {
      left = mid + 1
    } else {
      right = mid
    }
  }

  if left != nums.length && nums[left] == target {
    return left
  }

  return -1
}
```

Description:
- An advanced way to implement Binary Search.
- Search condition needs to access element's immediate right neighbor
- Use element's right neighbor to determine if condition is met and decide whether to go left or right
- Guarantees search space is at least 2 in size at each step
- Post-processing required. Loop/recursion ends when you have 1 element left. Need to assess if the remaining element meets the condition.

Distinguishing syntax:
- Initial condition: `left = 0, right = length`
- Termination `left == right`
- Searching left `right = mid`
- Searching righ `left = mid + 1`

#### Pattern II

```go
func binarySearch(nums []int, target int) int {
  if len(nums) == 0 {
    return -1
  }

  left := 0
  right := len(nums) - 1

  for left + 1 < right {
    mid := left + (right - left) / 2
    if nums[mid] == target {
      return mid
    } else if nums[mid] < target {
      left = mid
    } else {
      right = mid
    }
  }

  if nums[left] == target {
    return left
  }

  if nums[right] == target {
    return right
  }

  return -1
}
```

Description:
- An alternative way to implement Binary Search
- Search Condition needs to access element's immediate left and right neighbors
- Use element's neighbors to determine if condition is met and decide whether to go left or right
- Guarantees Search Space is at least 3 in size at each step
- Post-processing required. Loop/Recursion ends when you have 2 elements left. Need to assess if the remaining elements meet the condition.

Distinguishing syntax:
- Initial Condition: left = 0, right = length-1
- Termination: left + 1 == right
- Searching Left: right = mid
- Searching Right: left = mid

### Tree Traversal
#### In Order
##### Iterative

```go
func iterativeInorder(root *TreeNode) []int {
  result := []int{}

  if root == nil {
    return result
  }

  current := root
  stack := []*TreeNode{}

  for current != nil || len(stack) > 0 {
    for current != nil {
      stack = append(stack, current)
      current = current.Left
    }

    current = stack[len(stack) - 1]
    stack = stack[:len(stack) - 1]

    result = append(result, current.Val)
    current = current.Right
  }

  return result
}
```

##### Recursive

```go
func recursiveInOrder(root *TreeNode) []int {
  result := []int{}

  if root == nil {
    return result
  }

  recurse(root, &result)

  return result
}

func recurse(root *TreeNode, result *[]int) {
  if root == nil {
    return
  }

  recurse(root.Left, result)
  *result = append(*result, root.Val)
  recurse(root.Right, result)
}
```

#### Pre Order
##### Iterative

```go
func iterativePreOrder(root *TreeNode) []int {
  result := []int{}

  if root == nil {
    return result
  }

  stack := []*TreeNode{root}

  for len(stack) > 0 {
    node := stack[len(stack) - 1]
    stack = stack[:len(stack) - 1]

    result = append(result, node.Val)

    if node.Right != nil {
      stack = append(stack, node.Right)
    }

    if node.Left != nil {
      stack = append(stack, node.Left)
    }
  }

  return result
}
```

##### Recursive

```go
func recursivePreOrder(root *TreeNode) []int {
  result := []int{}

  if root == nil {
    return result
  }

  recurse(root, &result)

  return result
}

func recurse(root *TreeNode, result *[]int) {
  if root == nil {
    return
  }

  *result = append(*result, root.Val)
  recurse(root.Left, result)
  recurse(root.Right, result)
}
```

#### Post Order
##### Iterative

```go
func iterativePostOrder(root *TreeNode) []int {
  result := []int{}

  if root == nil {
    return result
  }

  stack := []*TreeNode{root}

  for len(stack) > 0 {
    node := stack[len(stack) - 1]
    stack = stack[:len(stack) - 1]

    result = append(result, node.Val)

    if node.Left != nil {
      stack = append(stack, node.Left)
    }

    if node.Right != nil {
      stack = append(stack, node.Right)
    }
  }

  reverse(result)

  return result
}
```

##### Recursive

```go
func recursivePostOrder(root *TreeNode) []int {
  result := []int{}

  if root == nil {
    return result
  }

  recurse(root, &result)

  return result
}

func recurse(root *TreeNode, result *[]int) {
  if root == nil {
    return
  }

  recurse(root.Left, result)
  recurse(root.Right, result)
  *result = append(*result, root.Val)
}
```

#### Level Order
##### Iterative

```go
func iterativeLevelOrder(root *TreeNode) [][]int {
  result := [][]int{}

  if root == nil {
    return result
  }

  queue := []*TreeNode{root}

  for len(queue) > 0 {
    level := []int{}
    qLen := len(queue)

    for i := 0; i < qLen; i++ {
      node := queue[0]
      queue = queue[1:]

      level = append(level, node.Val)

      if node.Left != nil {
        queue = append(queue, node.Left)
      }

      if node.Right != nil {
        queue = append(queue, node.Right)
      }
    }

    result = append(result, level)
  }

  return result
}
```

##### Recursive

Warning! This is not really BFS because by nature recursion is DFS. See GeeksforGeeks for more.

```go
func recursiveLevelOrder(root *TreeNode) [][]int {
  result := [][]int{}

  if root == nil {
    return result
  }

  level := 0
  recurse(root, level, &result)

  return result
}

func recurse(root *TreeNode, level int, result *[][]int) {
  if root == nil {
    return
  }

  if len(*result) - 1 < level {
    newLevel := []int{}
    *result = append(*result, newLevel)
  }

  (*result)[level] = append((*result)[level], root.Val)
  recurse(root.Left, level + 1, result)
  recurse(root.Right, level + 1, result)
}
```

### Permutation

```go
func generatePermutation(choices []int) [][]int {
  result := [][]int{}

  if len(nums) == 0 {
    return result
  }

  permutation := []int{}

  recurse(nums, &permutation, &result)

  return result
}

func recurse(choices []int, permutation *[]int, result *[][]int) {
  if len(choice) == 0 {
    *result = append(*result, createCopy(permutation))
    return
  }

  for i := 0; i < len(choices); i++ {
    newChoices := createNewChoicesExceptCurrent(choices, i)

    *permutation = append(permutation, choices[i])
    recurse(newChoices, permutation, result)
    *permutation = (*permutation)[:len(*permutation) - 1]
  }
}
```

### Subset

```go
func generateSubsets(set []int) [][]int {
  result := [][]int{}

  if len(nums) == 0 {
    return result
  }

  subset := []int{}
  startIndex := 0

  recurse(set, startIndex, &subset, &result)

  return result
}

func recurse(set []int, index int, subset *[]int, result *[][]int) {
  if index == len(nums) {
    *result = append(*result, createCopy(*subset))
    return
  }
  
  recurse(nums, index + 1, subset, result)

  *subset = append(*subset, nums[index])
  recurse(nums, index + 1, subset, result)
  *subset = (*subset)[:len(*subset) - 1]
}
```
 
## Tricky List

List of problems where it seems tricky:
- `data_structures/array_string/atoi`
- `data_structures/array_string/product_of_array_except_self`
- `data_structures/array_string/rotate_image`
- `data_structures/array_string/group_anagrams`
- `data_structures/array_string/reverse_words_in_a_string`
- `data_structures/array_string/next_permutation`
- `data_structures/array_string/integer_to_english_words`
- `data_structures/linked_list/reorder_list`
- `data_structures/linked_list/copy_list_with_random_pointers`
- `data_structures/binary_tree/diameter_of_binary_tree`
- `data_structures/binary_tree/binary_tree_maximum_path_sum`
- `data_structures/binary_tree/lowest_common_ancestor`
- `data_structures/binary_tree/binary_tree_vertical_order_traversal`
- `data_structures/binary_tree/serialize_deserialize_binary_tree`
- `data_structures/graph/find_the_celebrity`
- `data_structures/graph/clone_graph`
- `dynamic_programming/longest_common_substring/minimum_deletion_insertion_to_transform`
- `recursion/same_with_iteration/square_of_number`
- `recursion/with_numbers/gcd`
- `recursion/with_numbers/convert_to_binary`
- `recursion/with_arrays/reverse_array`
- `recursion/with_strings/remove_invalid_parentheses`
- `recursion/with_data_structures/reverse_linked_list`
- `recursion/with_data_structures/reverse_stack`
- `recursion/with_data_structures/unique_binary_search_trees_2`
- `two_pointers/remove_nth_node_from_end_of_list`
- `two_pointers/remove_duplicates_from_sorted_array`
- `sliding_window/longest_substring_without_repeating_chars`
- `sort_search/two_sum_2/`
- `sort_search/search_in_rotated_sorted_array/`
- `sort_search/intersection_of_two_arrays_2/`
- `sort_search/find_peak_element/`
- `sort_search/meeting_rooms_2/`
- `sort_search/3sum`

## TODO List
- Educative.io
  - Dynamic Programming
      - Palindromic Subsequence
        - Count of Palindromic Substrings
      - Longest Common Substring
        - Longest Common Substring
        - Maximum Sum Increasing Subsequence
        - Shortest Common Super-sequence
        - Minimum Deletions to Make a Sequence Sorted
        - Longest Repeating Subsequence
        - Subsequence Pattern Matching
        - Longest Bitonic Subsequence
        - Longest Alternative Subsequence
        - Edit Distance
        - Strings Interleaving
      - Try to do everything from bottom up style
  - Patterns for Coding Questions
      - Sliding Window
      - Two Pointers
      - Fast & Slow Pointers
      - Merge Intervals
      - Cyclic Sort
      - Linked List Reversal In Place
      - BFS
      - DFS
      - Two Heaps
      - Subsets
      - Modified Binary Search
      - Bitwise XOR
      - Top K Elements
      - K Way Merge
      - 0/1 Knapsack
      - Topological Sort
      - Misc
