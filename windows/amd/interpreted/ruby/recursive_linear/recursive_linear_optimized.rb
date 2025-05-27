def recursive_linear_search_optimized(arr, target, index = 0)
  # Early exit if remaining array is short
  return -1 if index >= arr.size
  return index if arr[index] == target
  recursive_linear_search_optimized(arr, target, index + 1)
end

# Test
arr = [2, 4, 6, 8, 10, 12, 14, 16, 18]
target = 14
index = recursive_linear_search_optimized(arr, target)
puts "Index: #{index}"
