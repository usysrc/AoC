local lines = {}
for line in io.lines("input") do
    lines[#lines+1] = line
end
for i,v in ipairs(lines) do
    for i in string.gmatch(example, "%S+") do
        print(i)
    end
end
