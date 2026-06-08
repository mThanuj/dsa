package io.github.mthanuj.arrays;

import io.github.mthanuj.Utils;

/*
* Problem:
* Given an integer array nums, rotate the array to the left by one.
*
*/
public class LeftRotateArrayByOne {

    public int[] solve(int[] arr) {
        Utils.reverse(arr, 0, 0);
        Utils.reverse(arr, 1, arr.length - 1);
        Utils.reverse(arr, 0, arr.length - 1);

        return arr;
    }

}
