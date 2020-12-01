-- from right to left
function forEveryDigitIn(number, fn)
    local num = number
    while num > 0 do
        local digit = num%10
        fn(digit)
        num = math.floor(num/10)
    end
end


function testForIncrease(number)
    local last = 99
    local result = true
    forEveryDigitIn(number, function(digit)
        if digit > last then
            result = false
        else
            last = digit
        end
    end)
    return result
end

function testForDoubles(number)
    local last = math.huge
    local doublecounter = 0
    local result = false
    forEveryDigitIn(number, function(digit)
        if digit == last then
            doublecounter = doublecounter + 1
            print(digit, doublecounter)
        else
            if doublecounter == 1 then
                result = true
            end
            last = digit
            doublecounter = 0
        end
    end)
    if doublecounter == 1 then
        result = true
    end
    return result
end

function call(fnA, fnB, val)
    return fnA(val) and fnB(val)
end

--[[
print("#####test increase#####")
print(testForIncrease(123456))
print(testForIncrease(7654321))
print(testForIncrease(123056))
print(testForIncrease(7657321))
--]]
print("#####test doubles#####")
print(testForDoubles(112233))
print(testForDoubles(123444))
print(testForDoubles(111122))
print(testForDoubles(221111))
print(testForDoubles(123345))

---[[

local start = 372304
local stop  = 847060

local count = 0
for i=start, stop do
    if call(testForDoubles, testForIncrease, i) then
        count = count + 1
    end
end

print(count)
--]]