package io.github.mthanuj.sorting;

import io.github.mthanuj.Utils;

public class BubbleSort {

    public int[] solve(int[] arr) {
        for (int i = 0; i < arr.length; i++) {
            for (int j = 0; j < arr.length - i - 1; j++) {
                if (arr[j] > arr[j + 1]) {
                    Utils.swap(arr, j, j + 1);
                }
            }
        }

        return arr;
    }

}
