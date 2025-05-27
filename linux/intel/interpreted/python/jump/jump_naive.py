import math

def jump_search(arr, key):
    n = len(arr)
    step = int(math.sqrt(n))
    prev = 0

    # Jump ahead in blocks until key is smaller than arr[step] or end reached
    while prev < n and arr[min(step, n) - 1] < key:
        prev = step
        step += int(math.sqrt(n))
        if prev >= n:
            return -1

    # Linear search within the identified block
    for i in range(prev, min(step, n)):
        if arr[i] == key:
            return i
    return -1

if __name__ == "__main__":
    arr = [2, 4, 6, 8, 10, 12, 14, 16, 18]
    key = 14

    result = jump_search(arr, key)
    print(f"Element {key} found at index {result}" if result != -1 else "Element not found")
