package io.github.mthanuj.arrays;

public class LargestElement {

    public int solve(int[] arr) {
        int max = arr[0];

        for (int i : arr) {
            if (i > max) {
                max = i;
            }
        }

        return max;
    }

}
