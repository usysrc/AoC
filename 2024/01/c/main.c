#include <stdio.h>
#include <stdlib.h>

// Comparison function
int compare(const void* a, const void* b) {
   return (*(int*)a - *(int*)b);
}

int main() {
    // read the input file
    FILE *input = fopen("input", "r");
    if (input == NULL) {
        printf("Error: input file not found\n");
        return 1;
    }

    // parse the input file into two arrays, format: <left number> three space <right number>
    int left[1000], right[1000];
    int rows = 0;
    while (fscanf(input, "%d   %d\n", &left[rows], &right[rows]) != EOF) {
        rows++;
    }

    // sort the two arrays with the stdlib
    qsort(left, rows, sizeof(int), compare);
    qsort(right, rows, sizeof(int), compare);

    // sum the two arrays
    long sumA = 0;
    for (int i = 0; i < rows; i++) {
        // calculate the abs distance of left[i] and right[i]
        int distance = abs(left[i]-right[i]);
        sumA += distance;
    }

    printf("Sum A: %ld\n", sumA);


    // count the number of occurences in the right array
    int count[10000000] = {0};
    for (int i = 0; i < rows; i++) {
        count[right[i]]++;
    }

    // calculate the similarity and sum it
    long sumB = 0;
    for (int i = 0; i < rows; i++) {
        if (count[right[i]] == 0) {
            continue;
        }
        int similarity = count[left[i]];
        sumB += left[i] * similarity;
    }

    printf("Sum B: %ld\n", sumB);

    // close the input file
    fclose(input);

    return 0;
}