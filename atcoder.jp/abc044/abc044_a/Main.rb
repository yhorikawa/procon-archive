n = gets.to_i
k = gets.to_i
x = gets.to_i
y = gets.to_i
sum = 0

1.upto(n) do |i|
  if i <= k
    sum += x
  else
    sum += y
  end
end

puts sum
