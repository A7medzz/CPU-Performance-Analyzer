def interpolation_search_naive(arr, target)
  low = 0
  high = arr.length - 1

  while low <= high && target >= arr[low] && target <= arr[high]
    # Avoid division by zero
    return -1 if arr[high] == arr[low]

    pos = low + ((target - arr[low]) * (high - low) / (arr[high] - arr[low]))

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
index = interpolation_search_naive(arr, target)
puts "Index: #{index}"
