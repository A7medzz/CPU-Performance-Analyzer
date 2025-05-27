def linear_search_naive(arr, target)
  for i in 0...arr.length
    return i if arr[i] == target
  end
  -1
end

# Test
arr = [2, 4, 6, 8, 10, 12, 14, 16, 18]
target = 14
index = linear_search_naive(arr, target)
puts "Index: #{index}"
