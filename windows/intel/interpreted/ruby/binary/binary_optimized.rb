def binary_search_optimized(arr, target, left = 0, right = nil)
  right = arr.length - 1 if right.nil?

  return -1 if left > right

  mid = (left + right) / 2
  if arr[mid] == target
    return mid
  elsif arr[mid] < target
    binary_search_optimized(arr, target, mid + 1, right)
  else
    binary_search_optimized(arr, target, left, mid - 1)
  end
end

# Test
arr = [2, 4, 6, 8, 10, 12, 14, 16, 18]
target = 14
index = binary_search_optimized(arr, target)
puts "Index: #{index}"
