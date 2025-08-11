# PUSH SWAP

Just a side project to grill my Go lang skill, get me out from comfort zone (since I am not a Go person)

## TODO

- [ ] Wrapper for basic operation (that print out steps)
- [ ] Dry run for diff operation
    - [ ] Dummy way, just search and push everything to b (sorted), and push back to a
    - [ ] Half Dummy way, push half of things to b, and set a mid pts, then sort a, and push everything back to a
    - [ ] chunk method
    - [ ] any other

## How to run

## Checker function

## How to unit test

This project includes unit tests to verify the correctness of the utility functions.

To run all tests, execute the following command from the root of the project:

```bash
go test ./...
```

This command will discover and run all test files (files ending in `_test.go`) in all sub-packages.

## Extra

### LIS algo explains

```
Recommended Algorithm: LIS + Smart Insertion

  This algorithm is efficient because it intelligently decides which elements to keep in Stack A and which to move, minimizing unnecessary pushes and rotations.

  Step 1: Find the Longest Increasing Subsequence (LIS)

   1. Identify the LIS: First, run an LIS algorithm on your input numbers in Stack A. The LIS is the longest subset of numbers that are already in sorted order relative to each
      other. For example, in [4, 6, 1, 7, 2, 8, 3], the LIS is [4, 6, 7, 8].
   2. Mark the LIS numbers: Keep track of which numbers are part of this LIS. These are your "anchor" elements. The core idea is to keep these elements in Stack A and move
      everything else.

  Step 2: Push Non-LIS Elements to Stack B

   1. Iterate through Stack A: Go through each element in Stack A from top to bottom.
   2. Push or Rotate:
       * If the current element is NOT in the LIS, push it to Stack B (pb).
       * If the current element IS in the LIS, rotate Stack A (ra) to move to the next element, leaving the LIS member in the stack.
   3. Result: After this step, Stack A will contain only the numbers from the LIS (already in relative sorted order), and Stack B will contain all the other numbers.

  Step 3: Insert Elements from B back to A (The "Smart" Part)

  This is the most critical part for efficiency. You need to insert each element from Stack B back into Stack A in the correct sorted position, using the minimum number of moves.

   1. Analyze Insertion Cost: For each number in Stack B (starting from the top), calculate the "cost" to move it to its correct place in Stack A. The cost is the total number
      of rotations required.
       * Find Target Position: Determine where the number from B should go in A. For example, if you're moving 5 from B to A, and A contains [2, 4, 8, 10], the target position
         for 5 is after 4.
       * Calculate Rotations for A: How many ra or rra moves does it take to bring that target position (4) to the top of Stack A?
       * Calculate Rotations for B: How many rb or rrb moves does it take to bring the number you want to move to the top of Stack B? (If you process one-by-one from the top,
         this cost is initially 0).
       * Combine Rotations: See if you can use rr (rotate A and B together) or rrr (reverse rotate both) to reduce the total move count. For example, if you need 5 ra and 3 rb,
         you can do 3 rr and 2 ra.
   2. Find the Cheapest Move: After calculating the cost for the top several elements in Stack B (or all of them), find the element that is "cheapest" to move right now.
   3. Execute and Push: Perform the cheapest set of rotations you identified, and then push the element from B to A (pa).
   4. Repeat: Continue this process until Stack B is empty.

  Step 4: Final Rotation of Stack A

   1. Align the Stack: At this point, Stack A is fully sorted, but the smallest number might not be at the top.
   2. Final Rotation: Perform a final series of ra or rra moves to rotate the stack until the absolute smallest number is at the top.

  Why is this better?

   * Fewer Pushes: By keeping the LIS in Stack A, you drastically reduce the number of elements you need to pb to Stack B and then pa back. Pushing and popping is expensive.
   * Optimized Rotations: The cost analysis for insertion ensures you don't waste moves. Instead of rotating A and B separately, you find the most efficient combination of moves
     (ra, rb, rr, rra, rrb, rrr) for each insertion.

  This approach is more complex to implement than simpler sorting algorithms, but it consistently yields a much lower number of steps, which is the primary goal of the Push Swap
  project.
```

