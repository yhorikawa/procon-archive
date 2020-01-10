X, Y = gets.chomp.split.map(&:to_i)

if X * 500 >= Y then
  puts "Yes"
else
  puts "No"
end
