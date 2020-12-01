
function calculateFuel(mass)
    return math.floor(mass/3) - 2
end

assert(calculateFuel(12) == 2)
assert(calculateFuel(14) == 2)
assert(calculateFuel(1969) == 654)
assert(calculateFuel(100756) == 33583)

local sum = 0
function addToSum(n)
    sum = sum + calculateFuel(n)
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



-- readFromStdIn()