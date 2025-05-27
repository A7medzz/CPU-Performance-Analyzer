def recursive_linear_search(arr, key, index=0):
    if index == len(arr):
        return -1
    if arr[index] == key:
        return index
    return recursive_linear_search(arr, key, index + 1)

if __name__ == "__main__":
    arr = [2, 4, 6, 8, 10, 12, 14, 16, 18]
    key = 14

    result = recursive_linear_search(arr, key)
    print(f"Element {key} found at index {result}" if result != -1 else "Element not found")
