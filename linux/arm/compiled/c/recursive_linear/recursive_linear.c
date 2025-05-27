#include <stdio.h>

int recursive_linear_search(int arr[], int n, int x) {
    if (n == 0)
        return -1;
    if (arr[n - 1] == x)
        return n - 1;
    return recursive_linear_search(arr, n - 1, x);
}

int main() {
    int arr[] = {2, 4, 6, 8, 10, 12, 14, 16, 18};
    int n = sizeof(arr) / sizeof(arr[0]);
    int x = 14;

    int result = recursive_linear_search(arr, n, x);
    if (result != -1)
        printf("Element %d found at index %d\n", x, result);
    else
        printf("Element %d not found\n", x);

    return 0;
}
