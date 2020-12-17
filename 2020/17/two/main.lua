local pos = function(x,y,z)
    return x..","..y..","..z
end

local posFromCell = function(cell)
    return pos(cell.x, cell.y, cell.z)
end

local coords = {}

local get = function(x,y,z)
    return coords[pos(x,y,z)]
end

local newCell = function(x, y, z, char)
    return {
        x = x,
        y = y,
        z = z,
        char = char
    }
end

local copyCell = function(cell)
    return newCell(cell.x, cell.y, cell.z, cell.char)
end

local y = 0
local z = 0
for line in io.lines() do
    y = y + 1
    for x=1,#line do
        coords[pos(x,y,z)] = newCell(x,y,z,line:sub(x,x))
    end
end

local neighbours = function(fn, ...)
    for ix=-1,1 do
        for iy=-1,1 do
            for iz=-1,1 do
                if ix == 0 and iy == 0 and iz == 0 then
                else
                    fn(ix, iy, iz, ...)
                end
            end
        end
    end
end


local isActive = function(char)
    return char == "#"
end

local ifCell = function(ix, iy, iz, cell, thenDo, ...)
    local ncell = get(cell.x + ix, cell.y + iy, cell.z + iz)
    if ncell then
        thenDo(ncell, ...)
    end
    return res
end

local thenDo = function(ncell, count)
    if isActive(ncell.char) then 
        count()
    end
end

local counter = function()
    local c = 0
    return function()
        return c
    end, function()
        c = c + 1
    end
end

local addNeighbours = function()
    local newCoords = {}
    for p, cell in pairs(coords) do
        newCoords[p] = copyCell(cell)
        neighbours(function(ix, iy, iz)
            local p = pos(cell.x + ix, cell.y + iy, cell.z + iz)
            if not coords[p] then
                newCoords[p] = newCell(cell.x + ix, cell.y + iy, cell.z + iz, ".")
            end
        end)
    end
    coords = newCoords
end

local simulate = function()
    local newCoords = {}
    addNeighbours()
    for p, cell in pairs(coords) do
        local getCount, incCount = counter()
        neighbours(ifCell, cell, thenDo, incCount)
        if isActive(cell.char) then
            if getCount() == 2 or getCount() == 3  then 
                newCoords[p] = copyCell(cell)
            else
                newCoords[p] = newCell(cell.x, cell.y, cell.z, ".")
            end
        else
            if getCount() == 3 then 
                newCoords[p] = newCell(cell.x, cell.y, cell.z, "#")
            else
                newCoords[p] = newCell(cell.x, cell.y, cell.z, ".")
            end
        end
    end
    coords = newCoords
end

for i=1, 6 do
    simulate()
end

local getCount, incCount = counter()
for k, cell in pairs(coords) do
    if isActive(cell.char) then incCount() end
end
print(getCount())