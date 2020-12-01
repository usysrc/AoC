
local lines = {}

-- read the lines in table 'lines'
for line in io.lines() do
	table.insert(lines, tonumber(line))
end

for i,k in ipairs(lines) do
	local cache = {}
	for j,kk in ipairs(lines) do
		if cache[2020 - kk - k] then
			print(k * kk * (2020 - kk - k))
			break
		end
		cache[kk] = true
	end
end

