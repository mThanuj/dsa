package io.github.mthanuj.arrays;

/*
 * Problem:
 * Given an array, we have to find the largest element in the array.
 *
 */
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
