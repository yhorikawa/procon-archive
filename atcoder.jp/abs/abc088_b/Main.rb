N = gets.to_i
a = gets.split.map(&:to_i)

a.sort!.reverse!

alice = 0
bob = 0

N.times do |n|
  if n % 2 == 0 then
    alice += a[n]
  else
    bob += a[n]
  end
end

puts alice - bob
