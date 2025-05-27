def linear_search_optimized(arr, target)
  i = 0
  n = arr.length
  while i < n
    return i if arr[i] == target
    i += 1
  end
  -1
end

# Test
arr = [2, 4, 6, 8, 10, 12, 14, 16, 18]
target = 14
index = linear_search_optimized(arr, target)
puts "Index: #{index}"
