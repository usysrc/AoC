local crates = {}
local walls = {}

local player = {
    x = 5,
    y = 5,
}

local tw = 16
local th = 16

love.load = function()
    for k = 1, 20 do
        local i = math.random(1, 30)
        local j = math.random(1, 30)
        local crate = {
            x = i,
            y = j
        }
        crates[i .. "," .. j] = crate
        crates[(i + 1) .. "," .. j] = crate
    end
    for i = 0, 45 do
        walls[i .. ",0"] = true
        walls[i .. ",36"] = true
        walls["0," .. i] = true
        walls["45," .. i] = true
    end
    for i = 0, 50 do
        for j = 0, 50 do
            if math.random(1, 100) < 3 then
                walls[i .. "," .. j] = true
            end
        end
    end
end

love.draw = function()
    love.graphics.setColor(1, 1, 1)
    love.graphics.rectangle("fill", player.x * tw, player.y * th, tw, th)

    for _, crate in pairs(crates) do
        love.graphics.setColor(1, 0, 0)
        love.graphics.rectangle("fill", crate.x * tw, crate.y * th, 2 * tw, th)
        love.graphics.setColor(1, 1, 1)
        love.graphics.rectangle("line", crate.x * tw, crate.y * th, 2 * tw, th)
    end
    for i = 0, 50 do
        for j = 0, 50 do
            if walls[i .. "," .. j] then
                love.graphics.setColor(0, 1, 0)
                love.graphics.rectangle("fill", i * tw, j * th, tw, th)
            end
        end
    end
end

local function moveCrate(crate, dx, dy)
    crates[crate.x .. "," .. crate.y] = nil
    crates[(crate.x + 1) .. "," .. crate.y] = nil
    crate.x = crate.x + dx
    crate.y = crate.y + dy
    crates[crate.x .. "," .. crate.y] = crate
    crates[(crate.x + 1) .. "," .. crate.y] = crate
end


love.keypressed = function(key)
    local dx, dy = 0, 0
    if key == "left" then
        dx = -1
    end
    if key == "right" then
        dx = 1
    end
    if key == "up" then
        dy = -1
    end
    if key == "down" then
        dy = 1
    end
    if dx == 0 and dy == 0 then return end
    local crate = crates[player.x + dx .. "," .. player.y + dy]
    if crate then
        local moved = {}
        local moves = {}
        local resolve
        resolve = function(crate)
            if moved[crate] then
                return true
            end
            if walls[crate.x .. "," .. crate.y] then
                return false
            end
            local crateA = crates[crate.x + dx .. "," .. crate.y + dy]
            local crateB = crates[crate.x + (dx + 1) .. "," .. crate.y + dy]

            if walls[crate.x + dx .. "," .. crate.y + dy] then
                return false
            end
            if walls[crate.x + (dx + 1) .. "," .. crate.y + dy] then
                return false
            end

            if crateA and crateA ~= crate then
                if not resolve(crateA) then
                    return false
                end
            end
            if crateB and crateB ~= crate then
                if not resolve(crateB) then
                    return false
                end
            end
            if walls[crate.x + dx .. "," .. crate.y + dy] then
                return false
            end
            if walls[(crate.x + dx + 1) .. "," .. crate.y + dy] then
                return false
            end
            table.insert(moves, function()
                moveCrate(crate, dx, dy)
            end)
            moved[crate] = true
            return true
        end
        if resolve(crate) then
            for _, move in ipairs(moves) do
                move()
            end
        end
    end
    if crates[player.x + dx .. "," .. player.y + dy] == nil then
        player.x = player.x + dx
        player.y = player.y + dy
    end
end
