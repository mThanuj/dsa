package io.github.mthanuj.arrays;

import io.github.mthanuj.Utils;

public class LeftRotateArrayByOne {

    public int[] solve(int[] arr) {
        Utils.reverse(arr, 0, 0);
        Utils.reverse(arr, 1, arr.length - 1);
        Utils.reverse(arr, 0, arr.length - 1);

        return arr;
    }

}
