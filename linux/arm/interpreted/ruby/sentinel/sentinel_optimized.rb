def sentinel(arr, target)
  n = arr.length
  last = arr[n-1]
  arr[n-1] = target
  i = 0
  while arr[i] != target
    i += 1
  end
  arr[n-1] = last
  if i < n-1 || arr[n-1] == target
    return i
  end
  -1
end

arr = [2, 4, 6, 8, 10, 12, 14, 16, 18]
target = 14
puts "Index: #{sentinel(arr, target)}"
