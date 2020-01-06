N = gets.to_i
a = gets.split(" ").map(&:to_i)
ans = 0
flag = false
while true
  a.each do |n|
    flag = true if n.odd?
  end
  break if flag == true

  for i in 0..(N-1) do
    a[i] = a[i] / 2
  end
  ans += 1
end
puts ans
