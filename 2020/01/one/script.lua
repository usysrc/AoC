
local lines = {}

-- read the lines in table 'lines'
for line in io.lines() do
	table.insert(lines, tonumber(line))
end

local cache = {}
for i,k in ipairs(lines) do
--	print(k)
	if cache[2020 - k] then
		print(k * (2020-k))
		break
	end
	cache[k] = true
end

