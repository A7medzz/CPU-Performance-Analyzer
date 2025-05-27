#include <stdio.h>

int linear_search(int arr[], int n, int x) {
    for (int i = 0; i < n; i++) {
        if (arr[i] == x)
            return i;
    }
    return -1;
}

int main() {
    int arr[] = {2, 4, 6, 8, 10, 12, 14, 16, 18};
    int n = sizeof(arr) / sizeof(arr[0]);
    int x = 14;

    int result = linear_search(arr, n, x);
    if (result != -1)
        printf("Element %d found at index %d\n", x, result);
    else
        printf("Element %d not found\n", x);

    return 0;
}
