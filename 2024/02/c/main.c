#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

// choose some sensible limits
#define MAX_REPORTS 10000
#define MAX_LEVELS 100

bool isReportSafe(int report[], int levelsCount)
{
    if (levelsCount < 2)
    {
        return false;
    }
    int last = report[0];
    bool ascending = false;
    bool descending = false;
    for (int i = 1; i < levelsCount; i++)
    {
        int v = report[i];
        if (v == last)
        {
            return false;
        }
        if (!ascending && !descending)
        {
            if (v > last)
            {
                ascending = true;
            }
            if (v < last)
            {
                descending = true;
            }
        }
        if (ascending && (v < last || v - last > 3))
        {
            return false;
        }
        if (descending && (v > last || last - v > 3))
        {
            return false;
        }
        last = v;
    }
    // print the report
    for (int i = 0; i < levelsCount; i++)
    {
        printf("%d ", report[i]);
    }
    printf("\n");
    return true;
}

int main()
{
    // read the input file
    FILE *input = fopen("input", "r");
    if (input == NULL)
    {
        printf("Error: input file not found\n");
        return 1;
    }

    int report[MAX_REPORTS][MAX_LEVELS];
    int levelsCount[MAX_REPORTS];
    
    int line[MAX_LEVELS];
    int reportsCount = 0;
    // parse the input file into report, we don't know how many levels are in a report
    while(!feof(input))
    {
        int i = 0;
        do
        {
            fscanf(input, "%d", &report[reportsCount][i]);
            i++;
            if (i > MAX_LEVELS)
            {
                printf("Error: more than MAX_LEVELS %d levels in report \n", MAX_LEVELS);
                return 1;
            }
        } while (fgetc(input) != '\n' && !feof(input));
        levelsCount[reportsCount] = i;
        reportsCount++;
        if (reportsCount > MAX_REPORTS)
        {
            printf("Error: more than MAX_REPORTS %d reports \n", MAX_REPORTS);
            return 1;
        }
    }
    int sumA = 0;

    for (int i = 0; i < reportsCount; i++)
    {
        // check if the report is safe
        if (isReportSafe(report[i], levelsCount[i]))
        {
            sumA++;
        }
    }

    printf("SumA: %d\n", sumA);

    // close the input file
    fclose(input);

    return 0;
}