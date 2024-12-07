#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

#define MAX_OPERANDS 1000

// check if there is a possible permutation of operations that will result in left = right
bool calcA(long current, long i, long right[], int right_size, long left)
{
    if (i >= right_size)
    {
        return current == left;
    }

    long r = right[i];

    if (calcA(current + r, i + 1, right, right_size, left))
    {
        return true;
    }

    if (calcA(current * r, i + 1, right, right_size, left))
    {
        return true;
    }

    return false;
}

long concat(long a, long b)
{
    long c = b;
    while (c > 0)
    {
        a *= 10;
        c /= 10;
    }
    return a + b;
}

// check if there is a possible permutation of operations that will result in left = right
bool calcB(long current, long i, long right[], int right_size, long left)
{
    if (current > left)
    {
        return false;
    }

    if (i >= right_size)
    {
        return current == left;
    }

    long r = right[i];

    if (calcB(current + r, i + 1, right, right_size, left))
    {
        return true;
    }

    if (calcB(current * r, i + 1, right, right_size, left))
    {
        return true;
    }

    if (calcB(concat(current, r), i + 1, right, right_size, left))
    {
        return true;
    }

    return false;
}

int main(int argc, char *argv[])
{
    // read the input file
    if (argc != 2)
    {
        printf("Usage: %s <input file>\n", argv[0]);
        return 1;
    }
    const char *filename = argv[1];
    FILE *input = fopen(filename, "r");
    if (input == NULL)
    {
        printf("Error: input file not found\n");
        return 1;
        ;
    }

    long sumA = 0;
    long sumB = 0;
    while (!feof(input))
    {
        do
        {
            long left = 0;
            long right[MAX_OPERANDS] = {0};
            // read in the line in the form of "left:right" where right is an array of integers
            fscanf(input, "%ld:", &left);

            int index = 0;
            while (fscanf(input, "%ld", &right[index]) == 1)
            {
                index++;
                if (index >= MAX_OPERANDS)
                {
                    printf("Error: too many operands\n");
                    return 1;
                }
                if (fgetc(input) == '\n' || feof(input))
                {
                    break;
                }
            }
            if (calcA(0, 0, right, index, left))
            {
                sumA += left;
            }

            if (calcB(0, 0, right, index, left))
            {
                sumB += left;
            }

        } while (!feof(input));
    }

    printf("Sum A: %ld\n", sumA);
    printf("Sum B: %ld\n", sumB);

    // close the input file
    fclose(input);

    return 0;
}