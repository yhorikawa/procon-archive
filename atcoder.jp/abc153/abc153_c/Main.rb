H, K = gets.chomp.split.map(&:to_i)
h = gets.chomp.split.map(&:to_i)

h.sort!

K.times do
  h.pop
end

if sum = h.inject(:+)
  puts sum
else
  puts 0
end
