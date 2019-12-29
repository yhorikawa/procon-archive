require 'prime'

x = gets.to_i

while x.prime? == false do
    x = x + 1
end

puts x