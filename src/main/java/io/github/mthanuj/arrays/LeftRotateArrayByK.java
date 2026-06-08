package io.github.mthanuj.arrays;

import io.github.mthanuj.Utils;

/*
* Problem:
* Given an integer array nums, rotate the array to the left by k.
*
*/
public class LeftRotateArrayByK {

    public int[] solve(int[] arr, int k) {
        Utils.reverse(arr, 0, k - 1);
        Utils.reverse(arr, k, arr.length - 1);
        Utils.reverse(arr, 0, arr.length - 1);

        return arr;
    }

}
