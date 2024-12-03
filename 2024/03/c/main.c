#include <stdio.h>
#include <stdlib.h>
#include <regex.h>
#include <string.h>
#include <stdbool.h>

int partA(void) {
    // read the input file
    FILE *input = fopen("input", "r");
    if (input == NULL)
    {
        printf("Error: input file not found\n");
        return 1;
    }

    // create the regex for mul((d+),(d+))
    regex_t regexA;
    int retiA = regcomp(&regexA, "mul\\(([0-9]+),([0-9]+)\\)", REG_EXTENDED);
    if (retiA)
    {
        printf("Error: regex compilation failed\n");
        return 1;
    }

    // read the file line by line and apply the regex
    char *line = NULL;
    size_t len = 0;
    ssize_t read;
    int resultA = 0;
    for (int i = 1; (read = getline(&line, &len, input)) != -1; i++)
    {
        regmatch_t matches[3];
        char *cursor = line;
        while (regexec(&regexA, cursor, 3, matches, 0) == 0)
        {
            char *first = malloc(matches[1].rm_eo - matches[1].rm_so + 1);
            strncpy(first, cursor + matches[1].rm_so, matches[1].rm_eo - matches[1].rm_so);

            char *second = malloc(matches[2].rm_eo - matches[2].rm_so + 1);
            strncpy(second, cursor + matches[2].rm_so, matches[2].rm_eo - matches[2].rm_so);

            // printf("Line %d, match: %s * %s\n", i, first, second);
            resultA += atoi(first) * atoi(second);
            cursor+=matches[0].rm_eo;

            free(first);
            free(second);
        }
    }

    printf("Result A: %d\n", resultA);

    return 0;
}

int partB(void) {
   // read the input file
    FILE *input = fopen("input", "r");
    if (input == NULL)
    {
        printf("Error: input file not found\n");
        return 1;
    } 
    regex_t regexB;
    int retiB = regcomp(&regexB, "mul\\(([0-9]+),([0-9]+)\\)|do()|don't()", REG_EXTENDED);
    if (retiB)
    {
        printf("Error: regex compilation failed\n");
        return 1;
    }

    bool doIt = true;
    // read the file line by line and apply the regex
    char *line = NULL;
    size_t len = 0;
    ssize_t read;
    int resultB = 0;
    for (int i = 1; (read = getline(&line, &len, input)) != -1; i++)
    {
        regmatch_t matches[3];
        char *cursor = line;
        while (regexec(&regexB, cursor, 3, matches, 0) == 0)
        {
            char* match = malloc(matches[0].rm_eo - matches[0].rm_so + 1);
            strncpy(match, cursor + matches[0].rm_so, matches[0].rm_eo - matches[0].rm_so);
            if (strcmp(match, "do") == 0) {
                doIt = true;
            } else if (strcmp(match, "don't") == 0) {
                doIt = false;
            }

            char *first = malloc(matches[1].rm_eo - matches[1].rm_so + 1);
            strncpy(first, cursor + matches[1].rm_so, matches[1].rm_eo - matches[1].rm_so);

            char *second = malloc(matches[2].rm_eo - matches[2].rm_so + 1);
            strncpy(second, cursor + matches[2].rm_so, matches[2].rm_eo - matches[2].rm_so);

            if (doIt) {
                // printf("Line %d, match: %s * %s\n", i, first, second);
                resultB += atoi(first) * atoi(second);
            }
            
            cursor+=matches[0].rm_eo;

            free(match);
            free(first);
            free(second);
        }
    }
    
    printf("Result B: %d\n", resultB);

    // close the input file
    fclose(input);
    
    return 0;
}

int main()
{
    int resultA = partA();
    if (resultA != 0)
    {
        return resultA;
    }
    int resultB = partB();
    if (resultB != 0)
    {
        return resultB;
    }
}