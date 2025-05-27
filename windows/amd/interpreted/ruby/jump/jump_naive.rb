def jump_search_naive(arr, target)
  n = arr.length
  step = Math.sqrt(n).to_i
  prev = 0

  while arr[[step, n].min - 1] < target
    prev = step
    step += Math.sqrt(n).to_i
    return -1 if prev >= n
  end

  while arr[prev] < target
    prev += 1
    return -1 if prev == [step, n].min
  end

  return prev if arr[prev] == target

  -1
end

# Test
arr = [2, 4, 6, 8, 10, 12, 14, 16, 18]
target = 14
index = jump_search_naive(arr, target)
puts "Index: #{index}"
