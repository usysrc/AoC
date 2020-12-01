
functions = {
    [1] = function(program, a,b,c)
        program[c] = program[a] + program[b]
    end,
    [2] = function(program, a,b,c)
        program[c] = program[a] * program[b]
    end
}

function run(program)
    if (not program) then return end
    local i = 0
    while(program[i+1]) do
        print(table.concat(program, ","))
        local opcode = program[i+1]
        if (opcode == 99) then
            break
        end
        local arg1, arg2, arg3 = program[i+2], program[i+3], program[i+4]
        assert(functions[opcode])
        functions[opcode](program, arg1, arg2, arg3)
        i = i + 4
    end
    print(program[0])
end

run{1,9,10,3,
2,3,11,0,
99,
30,40,50}