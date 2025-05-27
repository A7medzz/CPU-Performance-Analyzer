# Using sentinel technique: append the key at the end temporarily to avoid bounds checking inside the loop

def sentinel_search(arr, key):
    n = len(arr)
    last = arr[-1]
    arr[-1] = key

    i = 0
    while arr[i] != key:
        i += 1

    arr[-1] = last

    if i < n - 1 or arr[-1] == key:
        return i
    return -1

if __name__ == "__main__":
    arr = [2, 4, 6, 8, 10, 12, 14, 16, 18]
    key = 14

    # Convert to a list so we can modify the last element
    arr = list(arr)

    result = sentinel_search(arr, key)
    print(f"Element {key} found at index {result}" if result != -1 else "Element not found")
