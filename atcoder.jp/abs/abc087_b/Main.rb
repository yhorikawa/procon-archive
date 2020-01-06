A = gets.to_i
B = gets.to_i
C = gets.to_i
X = gets.to_i

count = 0

for a in 0..A do
  for b in 0..B do
    for c in 0..C do
      total = a * 500 + b * 100 + c * 50
      count += 1 if total == X
    end
  end
end
puts count
