local creds = {}
local input = ""
for line in io.lines() do
    if line == "" or line == "\n" then
        -- print("====\n" .. input .. "\n====\n")
        creds[#creds + 1] = input
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

local count = 0
for _, passport in ipairs(creds) do 
    local valid = true
    for _, field in ipairs(fields) do
        if not string.match(passport, field) then 
            valid = false
        end
    end
    if valid then 
        count = count + 1
    end
end

print(count)