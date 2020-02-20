# Data Structures and Algorithms

## Resources

- educative.io
- leetcode.com
- Study Guide https://www.reddit.com/r/cscareerquestions/comments/6luszf/a_leetcode_grinding_guide/?st=J92XK6ZK&sh=c2576763

## ASCII to Remember

- '0' is 48
- 'A' is 65
- 'a' is 97

## Pattern List

List of useful patterns

- Sliding Window
- Two Pointer
- Binary Search (3 versions)
    - Template I

        For accessing a single index in the array.

        ```go
        func binarySearch(nums []int, target int) int {
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

- DFS Recursive
- DFS Iterative
- BFS Recursive
- BFS Iterative
- Permutation
- Subset
 
## Tricky List

List of problems where the solution might not be that obvious.

- `data_structures/array_string/rotate_image`
- `data_structures/array_string/group_anagrams`
- `data_structures/array_string/reverse_words_in_a_string`
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
