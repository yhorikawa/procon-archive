#!/usr/bin/env ruby
n = gets.chomp.to_i
a = gets.chomp.split.map(&:to_i)

a.sort!.reverse!
count = 0

(n - 1).times do |i|
  count += a[((i + 1) / 2.0).to_i]
end

puts count
