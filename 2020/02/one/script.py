import fileinput

validCount = 0

for line in fileinput.input():
    policy, password = line.rstrip().split(":")
    policyRange, character = policy.split(" ")
    min, max = policyRange.split("-")
    count = 0
    for c in password:
        if c == character:
            count = count + 1
    if count >= int(min) and count <= int(max):
        validCount = validCount + 1

print(validCount)
