package io.github.mthanuj.sorting;

public class InsertionSort {

    public int[] solve(int[] arr) {
        for (int i = 1; i < arr.length; i++) {
            int j = i - 1;
            int value = arr[i];
            while (j > -1 && arr[j] > value) {
                arr[j + 1] = arr[j];
                j--;
            }

            arr[j + 1] = value;
        }

        return arr;
    }

}
