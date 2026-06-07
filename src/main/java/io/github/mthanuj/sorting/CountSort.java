package io.github.mthanuj.sorting;

public class CountSort {

    public int[] solve(int[] arr) {
        int max = arr[0];
        for (int i : arr) {
            if (i > max) {
                max = i;
            }
        }

        int min = arr[0];
        for (int i : arr) {
            if (i < min) {
                min = i;
            }
        }

        int total = Math.abs(min) + Math.abs(max);

        int[] counts = new int[total + 1];
        for (int i : arr) {
            counts[i + min]++;
        }

        int k = 0;
        for (int i = 0; i < counts.length; i++) {
            for (int j = 0; j < counts[i]; j++) {
                arr[k++] = i - min;
            }
        }

        return arr;
    }

}
