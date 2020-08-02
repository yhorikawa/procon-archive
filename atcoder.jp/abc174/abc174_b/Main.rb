#!/usr/bin/env ruby
n, d = gets.chomp.split.map(&:to_i)

count = 0

n.times do
  x, y = gets.chomp.split.map(&:to_i)
  distance = Math.sqrt(x ** 2 + y ** 2)
  count += 1 if distance <= d
end

puts count
