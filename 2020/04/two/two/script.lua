-- this is broken and not working

local creds = {}
local input = ""
for line in io.lines() do
    if line == "" or line == "\n" then
        -- print("====\n" .. input .. "\n====\n")
        creds[#creds + 1] = input.." "
        input = ""
    else
        input = input .. " " .. line
    end
end
creds[#creds + 1] = input



local fields = {
    "byr",
    "iyr",
    "eyr",
    "hgt",
    "hcl",
    "ecl",
    "pid",
    --"cid"
}

local eyecolors = {
    "amb",
    "blu",
    "brn",
    "gry",
    "grn",
    "hzl",
    "oth"
}

local count = 0
for _, passport in ipairs(creds) do 
    local valid = true
    for _, field in ipairs(fields) do
        if not string.match(passport, field) then 
            valid = false
        end
    end
    local byr = tonumber(string.match(passport, "byr:(%S+)%s"))
    if byr and (byr < 1920 or byr > 2002) then valid = false end
    
    local iyr = tonumber(string.match(passport, "iyr:(%S+)%s"))
    if iyr and (iyr < 2010 or iyr > 2020) then valid = false end

    local eyr = tonumber(string.match(passport, "eyr:(%S+)%s"))
    if eyr and (eyr < 2020 or eyr > 2030) then valid = false end

    local hgt = tonumber(string.match(passport, "hgt:(%S+)in%s"))
    if hgt then if (hgt < 59 or hgt > 73) then valid = false end end

    local hgt = tonumber(string.match(passport, "hgt:(%S+)cm%s"))
    if hgt then if (hgt < 150 or hgt > 193) then valid = false end end

    local hcl = string.match(passport, "hcl:#(%w%w%w%w%w%w)%s")
    if not hcl then valid = false end

    local ecl = tonumber(string.match(passport, "ecl:(%S+)%s"))
    if ecl then
        local found = false
        for i,v in ipairs(eyecolors) do
            if v == ecl then
                found = true
            end
        end
        if not found then valid = false end
    end

    local pid = tonumber(string.match(passport, "pid:(%d%d%d%d%d%d%d%d%d)%s"))
    if not pid then
        valid = false
    end
    
    if valid then 
        print(passport)
        count = count + 1
    end
end

print(count)