package io.github.mthanuj.sorting;

import io.github.mthanuj.Utils;

public class QuickSort {

    public int[] solve(int[] arr) {
        quickSort(arr, 0, arr.length - 1);
        return arr;
    }

    public void quickSort(int[] arr, int low, int high) {
        if (low < high) {
            int p = partition(arr, low, high);

            quickSort(arr, low, p - 1);
            quickSort(arr, p + 1, high);
        }
    }

    public int partition(int[] arr, int low, int high) {
        int pivot = arr[high];
        int i = low;

        for (int j = low; j < high; j++) {
            if (arr[j] <= pivot) {
                Utils.swap(arr, i, j);
                i++;
            }
        }

        Utils.swap(arr, i, high);
        return i;
    }

}
