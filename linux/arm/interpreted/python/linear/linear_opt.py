# Optimizations:

# Cache length

# Use enumerate for slightly faster iteration (avoiding indexing)

# Early exit is already present


def linear_search(arr, key):
    length = len(arr)
    for index, value in enumerate(arr):
        if value == key:
            return index
    return -1

if __name__ == "__main__":
    arr = [2, 4, 6, 8, 10, 12, 14, 16, 18]
    key = 14

    result = linear_search(arr, key)
    print(f"Element {key} found at index {result}" if result != -1 else "Element not found")
