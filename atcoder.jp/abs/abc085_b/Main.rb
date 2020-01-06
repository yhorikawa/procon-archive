N = gets.to_i
a = []

while n = gets do
  a.push(n.chomp)
end

puts a.uniq.length
