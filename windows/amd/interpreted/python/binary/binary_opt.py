# Iterative version to avoid recursion overhead and local caching of variables:

def binary_search(arr, key):
    low = 0
    high = len(arr) - 1

    while low <= high:
        mid = (low + high) >> 1  # faster than (low + high) // 2
        mid_val = arr[mid]

        if mid_val == key:
            return mid
        elif mid_val < key:
            low = mid + 1
        else:
            high = mid - 1

    return -1

if __name__ == "__main__":
    arr = [2, 4, 6, 8, 10, 12, 14, 16, 18]
    key = 14

    result = binary_search(arr, key)
    print(f"Element {key} found at index {result}" if result != -1 else "Element not found")
