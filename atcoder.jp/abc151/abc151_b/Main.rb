N, K, M = gets.chomp.split.map(&:to_i)
a = gets.split.map(&:to_i)

if a.inject(:+) > N * M
  puts 0
else
  if  (N * M - a.inject(:+)) > K
    puts -1
  else
    puts N * M - a.inject(:+)
  end
end