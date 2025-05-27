def interpolation_search_optimized(arr, target)
  low = 0
  high = arr.length - 1

  while low <= high && target >= arr[low] && target <= arr[high]
    range_diff = arr[high] - arr[low]
    return -1 if range_diff == 0

    # Precompute position with more clarity and minimal repeated operations
    numerator = (target - arr[low]) * (high - low)
    pos = low + (numerator / range_diff)

    return pos if arr[pos] == target

    if arr[pos] < target
      low = pos + 1
    else
      high = pos - 1
    end
  end

  -1
end

# Test
arr = [2, 4, 6, 8, 10, 12, 14, 16, 18]
target = 14
index = interpolation_search_optimized(arr, target)
puts "Index: #{index}"
