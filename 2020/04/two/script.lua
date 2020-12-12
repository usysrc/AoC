-- this is broken and not working

local creds = {}
local input = {}
for line in io.lines() do
    if line == "" or line == "\n" then
        -- print("====\n" .. input .. "\n====\n")
        table.sort(input)
        creds[#creds + 1] = " "..table.concat(input, " ").." "
        input = {}
    else
        for id in string.gmatch(line, "%S+") do
            input[#input+1] = id
        end
    end
end
creds[#creds + 1] = " "..table.concat(input, " ").." "

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
    print(passport)

    local valid = true
    for _, field in ipairs(fields) do
        if not string.match(passport, field) then 
            print('missing field', field)
            valid = false
        end
    end

    local byr = tonumber(string.match(passport, "byr:(%d%d%d%d)%s"))
    if byr and (byr < 1920 or byr > 2002) then
        print("problem with byr")
        valid = false 
    elseif not byr then
        print("problem with byr")
        valid = false
    end

    local iyr = tonumber(string.match(passport, "iyr:(%d%d%d%d)%s"))
    if iyr and (iyr < 2010 or iyr > 2020) then
        print("problem with iyr")
        valid = false 
    elseif not iyr then
        print("problem with iyr")
        valid = false
    end

    local eyr = tonumber(string.match(passport, "eyr:(%d%d%d%d)%s"))
    if eyr and (eyr < 2020 or eyr > 2030) then
        print("problem with eyr")
        valid = false
    elseif not eyr then
        print("problem with eyr")
        valid = false
    end

    local hgt = tonumber(string.match(passport, "hgt:(%d%d)in%s"))
    if hgt then if (hgt < 59 or hgt > 76) then
        print("problem with hgt, inch", hgt)
        valid = false
    end end

    local hgtt = tonumber(string.match(passport, "hgt:(%d%d%d)cm%s"))
    if hgtt then if (hgtt < 150 or hgtt > 193) then
        print("problem with hgt, cm")
        valid = false 
    end end
    if not hgt and not hgtt then
        print("problem with hgt")
        valid = false
    end

    local hcl = string.match(passport, "hcl:#(%x%x%x%x%x%x)%s+")
    if not hcl then
        print("problem with hcl")
        valid = false
    end

    local ecl = string.match(passport, "ecl:(%a%a%a)%s+")
    if ecl then
        local found = false
        for i,v in ipairs(eyecolors) do
            if v == ecl then
                found = true
            end
        end
        if not found then 
            print("problem with ecl")
            valid = false
        end
    else
        valid = false
    end

    local pid = tonumber(string.match(passport, "pid:(%d%d%d%d%d%d%d%d%d)%s"))
    if not pid then
        print("problem with pid", pid)
        valid = false
    end

    print("Valid: ", valid)
    print()
    if valid then 
        -- print(passport)
        count = count + 1
    end
end

print(count)