# Optimize by:

# Avoiding repeated lookups (cache array length)

# Using local variables

# Removing redundant condition checks inside loop

def interpolation_search(arr, key):
    n = len(arr)
    low = 0
    high = n - 1
    arr_low = arr[0]
    arr_high = arr[high]

    while low <= high and key >= arr_low and key <= arr_high:
        if low == high:
            return low if arr[low] == key else -1

        # Calculate probe position using local variables
        denominator = arr_high - arr_low
        if denominator == 0:
            break
        pos = low + ((key - arr_low) * (high - low) // denominator)

        if arr[pos] == key:
            return pos
        elif arr[pos] < key:
            low = pos + 1
            arr_low = arr[low]
        else:
            high = pos - 1
            arr_high = arr[high]

    return -1


if __name__ == "__main__":
    arr = [2, 4, 6, 8, 10, 12, 14, 16, 18]
    key = 14

    result = interpolation_search(arr, key)
    print(f"Element {key} found at index {result}" if result != -1 else "Element not found")
