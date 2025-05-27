def binary_search(arr, key, low=0, high=None):
    if high is None:
        high = len(arr) - 1

    if low > high:
        return -1

    mid = (low + high) // 2
    if arr[mid] == key:
        return mid
    elif arr[mid] < key:
        return binary_search(arr, key, mid + 1, high)
    else:
        return binary_search(arr, key, low, mid - 1)

if __name__ == "__main__":
    arr = [2, 4, 6, 8, 10, 12, 14, 16, 18]
    key = 14

    result = binary_search(arr, key)
    print(f"Element {key} found at index {result}" if result != -1 else "Element not found")
