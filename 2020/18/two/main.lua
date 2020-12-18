local precedence = function(a)
    if a == "(" then
        return 1000
    elseif a == "+" then
        return 2
    else
        return 1
    end
end

local toRPN = function(str)
    local output = {}
    local operators = {}
    for i=1,#str do
        local char = str:sub(i,i)
        if tonumber(char) and tonumber(char) <= 9 then
            output[#output+1] = char
        else
            if char == ")" then
                while operators[#operators] and operators[#operators] ~= '(' do
                    output[#output+1] = operators[#operators]
                    operators[#operators] = nil
                end
                operators[#operators] = nil
            elseif char == "(" then
                operators[#operators+1] = char               
            else
                while #operators > 0 and operators[#operators] ~= "(" and precedence(char) <= precedence(operators[#operators]) do
                    output[#output+1] = operators[#operators]
                    operators[#operators] = nil
                end
                operators[#operators+1] = char
            end
        end
    end
    for i=#operators, 0, -1 do
        output[#output+1] = operators[i]
    end
    -- print(table.concat(output, ","))
    return output
end

local ops = {
    ["+"] = function(a,b)
        return a+b
    end,
    ["*"] = function(a,b)
        return a*b
    end,
    ["-"] = function(a,b)
        return a-b
    end
}

-- 1 3 4 2 3 * 5 6
-- 1 3 4 [2 3 *] 5 6
-- 1 3 4 * 5 6

function slice(tbl, n)
    -- remove n, n-1 and n-2 from table
    local t = {}
    for i,v in ipairs(tbl) do
        if i == n - 1 or i == n - 2 then
        else
            t[#t+1] = v
        end
    end
    return t, n - 2
end

function solveRPN(input)
    local i = 0
    while i <= #input do
        local char = input[i]
        if tonumber(char) and tonumber(char) <= 9 then
            -- print(char)
        elseif ops[char] then
            local op = char
            local left = input[i-2]
            local right = input[i-1]
            input, i = slice(input, i)
            input[i] = ops[char](left, right)
            -- print(table.concat(input, ","))
        end
        i = i + 1
    end
    return input[#input]
end

local sum = 0
for line in io.lines() do
    local line = string.gsub(line, " ", "")
    sum = sum + solveRPN(toRPN(line))
end
print(sum)
