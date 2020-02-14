# Data Structures and Algorithms

## Resources

- educative.io
- leetcode.com

## Pattern List

List of useful patterns

- Sliding Window
- Binary Search (3 versions)

  ```python
  def binarySearch(nums, target):
    """
    :type nums: List[int]
    :type target: int
    :rtype: int
    """
    if len(nums) == 0:
        return -1

    left, right = 0, len(nums) - 1
    while left <= right:
        mid = (left + right) // 2
        if nums[mid] == target:
            return mid
        elif nums[mid] < target:
            left = mid + 1
        else:
            right = mid - 1

    # End Condition: left > right
    return -1
  ```

- DFS Recursive
- DFS Iterative
- BFS Recursive
- BFS Iterative
- Permutation
- Subset
 
## Tricky List

List of problems where the solution might not be that obvious.

- `dynamic_programming/longest_common_substring/minimum_deletion_insertion_to_transform`
- `recursion/same_with_iteration/square_of_number`
- `recursion/with_numbers/gcd`
- `recursion/with_numbers/convert_to_binary`
- `recursion/with_arrays/reverse_array`
- `recursion/with_data_structures/reverse_linked_list`
- `recursion/with_data_structures/reverse_stack`

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
