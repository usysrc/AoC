-- this is broken...


-- read the notes
local notes = {}
for line in io.lines() do
    if line == "" then break end
    local left = string.match(line, "(%a+%s*%a+):")
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
            print(#ticket, num)
        end
    end
end

print()
-- read nearby tickets
local tickets = {}
for line in io.lines() do
    if line == "" then break end
    if line ~= "nearby tickets:" then 
        local v = {}
        for num in string.gmatch(line, "(%d+)") do
            v[#v+1] = tonumber(num)
        end
        tickets[#tickets+1] = v
    end
end

function getInRange(tickets)
    local notinrange = {}
    local inrange = {}
    local ticketsInRange = {}

    local sum = 0
    for _, ticket in ipairs(tickets) do
        for _, num in ipairs(ticket) do
            if inrange[num] then
                -- print("found")
            else
                local found = false
                for ii, vv in ipairs(notes) do
                    for iii, range in ipairs(vv.ranges) do
                        if num >= range.start and num <= range.stop then
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
                    notinrange[num] = true
                    inrange[num] = nil
                end
            end
            if not inrange[num] then
                sum = sum + num
                notinrange[num] = true
                inrange[num] = nil
                tout = true
            end
        end
        local tout = false
        for _, num in ipairs(ticket) do
            if notinrange[num] then
                tout = true
            end
        end
        if not tout then
            ticketsInRange[#ticketsInRange+1] = ticket
        end
    end
    return ticketsInRange
end

local ticketsInRange = getInRange(tickets)
local fields = {}

for i,v in ipairs(ticketsInRange[1]) do
    fields[i] = {}
end

for _,v in ipairs(ticketsInRange) do
    for i,vv in ipairs(v) do
        table.insert(fields[i], vv)
    end
end

local cat = {}
local bat = {}

for i,v in ipairs(fields) do
    for _, note in ipairs(notes) do
        local n = 0
        for _, num in ipairs(v) do
            for _, range in ipairs(note.ranges) do
                if num >= range.start and num <= range.stop then
                    n = n + 1
                    break
                end
            end
        end
        if n >= #v-1 and not cat[note.field] then
            cat[note.field] = i
            bat[i] = note.field
            print(i, note.field)
            break
        end
    end
end

-- -- validate ticketsinrange
for _,v in ipairs(ticketsInRange) do
    for i, num in ipairs(v) do
        local find = false
        local idx = cat[bat[i]]
        for _, range in ipairs(notes[idx].ranges) do
            if num >= range.start and num <= range.stop then
                find = true
                break
            end
        end
        if not find then
            print("validation failed for field ", idx, i)
            -- print(i, bat[i],idx)
        end
    end
end

print()
local n = 1
for k, idx in pairs(cat) do
    if string.match(k, "departure") then
        -- print(k,idx)
        n = n * ticket[idx]
        -- print(k, idx, ticket[idx], n)
    end
end
print(n)