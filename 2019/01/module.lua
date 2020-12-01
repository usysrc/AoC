
function calculateFuel(mass)
    return math.floor(mass/3) - 2
end

function calculateFuelOfFuel(mass)
    local fuel = calculateFuel(mass)
    local fuelOfFuel = calculateFuel(fuel)
    while (fuelOfFuel > 0) do
        if (fuelOfFuel > 0 ) then
            fuel = fuel + fuelOfFuel
        end
        fuelOfFuel = calculateFuel(fuelOfFuel)
    end
    return fuel
end

assert(calculateFuelOfFuel(12) == 2)
assert(calculateFuelOfFuel(1969) == 966)
assert(calculateFuelOfFuel(100756) == 50346)

local sum = 0
function addToSum(n)
    sum = sum + calculateFuelOfFuel(n)
    print(sum)
end

function process(input)
    addToSum(input)
end

function readFromStdIn()
    input = io.read()
    while(input) do
        process(input)
        input = io.read()
    end
end

readFromStdIn()