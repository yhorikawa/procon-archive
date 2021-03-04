N, X = gets.split.map(&:to_i)
alcohol = Array.new(N){gets.split.map(&:to_i)}
count = 0
sum = 0

alcohol.each do |n|
  sum += n[0]*n[1]
  count += 1
  if sum > X*100
    p count
    exit
  end
end

puts -1