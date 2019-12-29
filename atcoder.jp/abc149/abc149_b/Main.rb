a,b,k = gets.split.map(&:to_i) 
while k > 0 do
    if k < a then
        a = a - k
        k = 0
        break
    else
        k = k - a
        a = 0
    end
    if k < b then
        b = b -k
        k = 0
        break
    else
        k = k - b
        b = 0
    end
    if a = 0 && b = 0 then
        break
    end
end
puts "#{a} #{b}"