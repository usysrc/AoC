local imagedata = {}

---[[
for line in io.lines("input") do
    for i=1,#line do
        imagedata[#imagedata+1] = line:sub(i,i)
    end
end
--]]
--[[
local line = "0222112222120000"
for i=1,#line do
    imagedata[#imagedata+1] = line:sub(i,i)
end
--]]

local w,h = 25, 6
local layers = {}
local k = 0
for i=1,#imagedata do
    local layerNumber = math.floor(k / (w*h))
    local layer = layers[layerNumber] or {}
    layer[#layer+1] = imagedata[i]
    layers[layerNumber] = layer
    k = k + 1
end

--[[
-- part 1
local layerInfo = {}

function newLayerInfo(obj)
    local o = obj or {}
    o.numbers = {}
    for i=0, 9 do
        o.numbers[i] = 0
    end
    o.addNumber = function(num)
        local num = tonumber(num)
        o.numbers[num] =  o.numbers[num] + 1
    end
    return o
end


local min = math.huge
local minLayer = -1

for i, layer in pairs(layers) do
    local info = newLayerInfo(layerInfo[i])
    layerInfo[i] = info
    for k,pixel in ipairs(layer) do
        info.addNumber(pixel)
    end
    local zeros = info.numbers[0]
    if zeros < min then
        min = zeros
        minLayer = i
    end
end

--print(minLayer)
--local info = layerInfo[minLayer]
--print(info.numbers[1] * info.numbers[2])
--]]

local pixelmap = {
    [0] = " ",
    [1] = "O",
    [2] = " "

}

for j=1,h do
    local line = ""
    for i=1, w do
        local pixel = 2
        for layernumber=0, #imagedata/(w*h)-1 do
            local layer = layers[layernumber]
            if pixel == 2 then
                pixel = tonumber(layer[i+(j-1)*w])
            end
        end
        line = line .. pixelmap[tonumber(pixel)]
    end
    print(line)
end



