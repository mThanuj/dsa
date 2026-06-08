package io.github.mthanuj.arrays;

import io.github.mthanuj.Utils;

/*
* Problem:
* You are given an array of integers, your task is to move all the zeros in the array to the
* end of the array and move non-negative integers to the front by maintaining their order.
*
*/
public class MoveAllZerosToEndOfArray {

    public int[] solve(int[] arr) {
        int left = 0;

        while (arr[left] != 0) {
            left++;
        }

        for (int right = left + 1; right < arr.length; right++) {
            if (arr[right] != 0) {
                Utils.swap(arr, left, right);
                left++;
            }
        }

        return arr;
    }

}
