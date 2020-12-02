import fileinput

validCount = 0

for line in fileinput.input():
    policy, password = line.rstrip().split(":")
    policyRange, character = policy.split(" ")
    first, second = policyRange.split("-")
    if (password[int(first)] == character != password[int(second)]):
        validCount = validCount + 1
    elif (password[int(second)] == character != password[int(first)]):
        validCount = validCount + 1
    print(password[int(first)], password[int(second)], "== " + character, validCount, line)


print(validCount)
