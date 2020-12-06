local count = 0
local group = ""

local countAnswers = function(group)
    local answers = {}
    local persons = 0
    for person in group:gmatch("%w+") do 
        persons = persons + 1
        for i=1, #person do
            local c = person:sub(i,i)
            if not answers[c] then
                answers[c] = 1
            else
                answers[c] = answers[c] + 1
            end
        end
    end
    for k,v in pairs(answers) do
        if v == persons then
            count = count + 1
        end
    end
end

for line in io.lines() do
    group = group .. " " ..line
    if line == "" then
        -- print(group)
        countAnswers(group)
        group = ""
    end
end
countAnswers(group)

print(count)