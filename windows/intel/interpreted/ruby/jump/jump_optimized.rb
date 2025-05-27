def jump_search_optimized(arr, target)
  n = arr.length
  step = Math.sqrt(n).to_i
  prev = 0
  max_index = [step, n].min - 1

  # Jump ahead in blocks
  while max_index < n && arr[max_index] < target
    prev = step
    step += Math.sqrt(n).to_i
    max_index = [step, n].min - 1
  end

  # Linear scan within block
  for i in prev..max_index
    return i if arr[i] == target
  end

  -1
end

# Test
arr = [2, 4, 6, 8, 10, 12, 14, 16, 18]
target = 14
index = jump_search_optimized(arr, target)
puts "Index: #{index}"
