local count = 0
local group = ""

local countAnswers = function()
    local answers = {}
    for i=1,#group do
        local c = group:sub(i,i)
        if not answers[c] then
            count = count + 1
            answers[c] = true
        end
    end
end


for line in io.lines() do
    group = group .. line
    if line == "" then
        countAnswers(group)
        group = ""
    end
end
countAnswers(group)

print(count)