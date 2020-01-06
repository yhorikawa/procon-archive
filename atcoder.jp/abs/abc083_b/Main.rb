def findsum(num)
  sum = 0
  while num > 0 do
    sum += num % 10
    num /= 10
  end
  return sum
end

N, A, B = gets.chomp.split.map(&:to_i)
total = 0
0.upto(N) do |n|
  sum = findsum(n)
  total += n if sum >= A && sum <= B
end
puts total
