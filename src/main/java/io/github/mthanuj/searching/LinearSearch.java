package io.github.mthanuj.searching;

/*
 * Problem:
 * Given an array, and an element num the task is to find if num is present in the given array or not.
 * If present return the index of the element or print -1.
 *
 */
public class LinearSearch {

    public int solve(int[] arr, int target) {
        for (int i = 0; i < arr.length; i++) {
            if (arr[i] == target) {
                return i;
            }
        }

        return -1;
    }

}
