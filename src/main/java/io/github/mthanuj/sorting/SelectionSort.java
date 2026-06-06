package io.github.mthanuj.sorting;

import io.github.mthanuj.Utils;

public class SelectionSort {

    public int[] solve(int[] arr) {
        for (int i = 0; i < arr.length - 1; i++) {
            int smallest = i;
            for (int j = i + 1; j < arr.length; j++) {
                if (arr[j] <= arr[smallest]) {
                    smallest = j;
                }
            }

            Utils.swap(arr, i, smallest);
        }

        return arr;
    }

}
