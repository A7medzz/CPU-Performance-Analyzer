def interpolation_search(arr, key):
    low = 0
    high = len(arr) - 1

    while low <= high and key >= arr[low] and key <= arr[high]:
        if low == high:
            if arr[low] == key:
                return low
            return -1

        # Calculate probe position
        pos = low + ((key - arr[low]) * (high - low) // (arr[high] - arr[low]))

        if arr[pos] == key:
            return pos
        elif arr[pos] < key:
            low = pos + 1
        else:
            high = pos - 1

    return -1


if __name__ == "__main__":
    arr = [2, 4, 6, 8, 10, 12, 14, 16, 18]
    key = 14

    result = interpolation_search(arr, key)
    print(f"Element {key} found at index {result}" if result != -1 else "Element not found")
