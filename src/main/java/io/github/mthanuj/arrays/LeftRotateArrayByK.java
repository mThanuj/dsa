package io.github.mthanuj.arrays;

import io.github.mthanuj.Utils;

public class LeftRotateArrayByK {

    public int[] solve(int[] arr, int k) {
        Utils.reverse(arr, 0, k - 1);
        Utils.reverse(arr, k, arr.length - 1);
        Utils.reverse(arr, 0, arr.length - 1);

        return arr;
    }

}
