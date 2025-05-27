#include <stdio.h>

int sentinel_search(int arr[], int n, int x) {
    int last = arr[n - 1];
    arr[n - 1] = x;
    int i = 0;

    while (arr[i] != x)
        i++;

    arr[n - 1] = last;

    if ((i < n - 1) || (arr[n - 1] == x))
        return i;
    return -1;
}

int main() {
    int arr[] = {2, 4, 6, 8, 10, 12, 14, 16, 18};
    int n = sizeof(arr) / sizeof(arr[0]);
    int x = 14;

    int result = sentinel_search(arr, n, x);
    if (result != -1)
        printf("Element %d found at index %d\n", x, result);
    else
        printf("Element %d not found\n", x);

    return 0;
}
