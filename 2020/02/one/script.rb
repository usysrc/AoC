found = 0
File.readlines('input').each do |line|
    policy, password = line.split(":")
    range, char = policy.split(" ")
    min, max = range.split("-")
    count = 0
    password.split('').each do |cipher|
        if char == cipher
            count += 1
        end
    end
    if count >= min.to_i && count <= max.to_i
        found += 1     
    end
end
puts found