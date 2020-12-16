-- read the notes
local notes = {}
for line in io.lines() do
    if line == "" then break end
    local left = string.match(line, "(%a+):")
    local range1 = string.match(line, ": (%S+)")
    local range2 = string.match(line, "or (%S+)")

    local range1start, range1end = string.match(range1, "(%d+)-(%d+)")
    local range2start, range2end = string.match(range2, "(%d+)-(%d+)")

    notes[#notes+1] = {
        field = left,
        ranges = {
            {start = tonumber(range1start), stop = tonumber(range1end)},
            {start = tonumber(range2start), stop = tonumber(range2end)}
        }
    }
end

-- read the ticket
local ticket = {}
for line in io.lines() do
    if line == "" then break end
    if line ~= "your ticket:" then 
        for num in string.gmatch(line, "(%d+)") do
            ticket[#ticket+1] = tonumber(num)
        end
    end
end

-- read nearby tickets
local tickets = {}
for line in io.lines() do
    if line == "" then break end
    if line ~= "nearby ticket:" then 
        local v = {}
        tickets[#tickets+1] = v
        for num in string.gmatch(line, "(%d+)") do
            v[#v+1] = tonumber(num)
        end
    end
end


function checkInRange(tickets)
    local notinrange = {}
    local inrange = {}

    local sum = 0
    for _, ticket in ipairs(tickets) do
        for _, num in ipairs(ticket) do
            if inrange[num] then
                -- print("found")
            else
                local found = false
                for ii, vv in ipairs(notes) do
                    for iii, range in ipairs(vv.ranges) do
                        -- print(num,range.start, range.stop)
                        if num >= range.start and num <= range.stop then
                            -- print('found')
                            found = true
                            inrange[num] = true
                            if notinrange[num] then
                                notinrange[num] = nil
                            end
                            break
                        end
                    end
                    if found then break end
                end
                if not found and not notinrange[num] then
                    -- print("not found")
                    notinrange[num] = true
                end
            end
            if not inrange[num] then
                sum = sum + num
                -- print("not found", num)
                notinrange[num] = true
            end
        end
    end
    return sum
end

print(checkInRange(tickets))