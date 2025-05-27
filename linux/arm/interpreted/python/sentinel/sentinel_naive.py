def sentinel_search(arr, key):
    for i in range(len(arr)):
        if arr[i] == key:
            return i
    return -1

if __name__ == "__main__":
    arr = [2, 4, 6, 8, 10, 12, 14, 16, 18]
    key = 14

    result = sentinel_search(arr, key)
    print(f"Element {key} found at index {result}" if result != -1 else "Element not found")
