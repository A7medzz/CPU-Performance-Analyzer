# Improvements:

# Cache length and step size

# Use local variables

# Avoid multiple calls to min() inside loops

import math

def jump_search(arr, key):
    n = len(arr)
    step = int(math.sqrt(n))
    prev = 0
    last_index = n - 1
    block_end = step - 1 if step - 1 < last_index else last_index

    while prev <= last_index and arr[block_end] < key:
        prev = step
        step += int(math.sqrt(n))
        if prev > last_index:
            return -1
        block_end = step - 1 if step - 1 < last_index else last_index

    # Linear search within the block
    for i in range(prev, block_end + 1):
        if arr[i] == key:
            return i
    return -1

if __name__ == "__main__":
    arr = [2, 4, 6, 8, 10, 12, 14, 16, 18]
    key = 14

    result = jump_search(arr, key)
    print(f"Element {key} found at index {result}" if result != -1 else "Element not found")
