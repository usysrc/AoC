local type CacheKey = string
local function getKey(char: string, i: integer): CacheKey
    return char.."_"..i 
end
local cache: {CacheKey:integer} = {}

-- Remove leading zeros from a string
local function removeLeadingZeros(char: string): string
    while char:sub(1, 1) == "0" and #char > 1 do
        char = char:sub(2)
    end
    return char
end

-- Recursive function to simulate and count the stones
local function countStones(char: string, i: integer, max: integer): integer
    local key: CacheKey = getKey(char, i)
    if cache[key] then
        return cache[key]
    end
    local count:integer = 0
    if i == max then
        return 1
    end
    if char == "0" then
        count = count + countStones("1", i+1, max)
    elseif #char%2 == 0 then
        local left: string = removeLeadingZeros(char:sub(1, math.floor(#char/2)))
        local right: string = removeLeadingZeros(char:sub(math.floor(#char/2)+1, #char))
        count = count + countStones(left, i+1, max)
        count = count + countStones(right, i+1, max)
    else
        count = count + countStones(tostring(tonumber(char)*2024), i+1, max)
    end
    cache[key] = count
    return count
end

-- Read input and simulate the stone counting
for line in io.lines("input") do
    local sum = 0
    for word in string.gmatch(line, "%S+") do
        sum = sum + countStones(word, 0, 75)
    end
    print("sum: "..sum)
end