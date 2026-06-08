package io.github.mthanuj.arrays;

/*
* Problem:
* Given an array, second largest element in the array.
* Print ‘-1’ in the event that either of them doesn’t exist.
*
*/
public class SecondLargestElement {

    public int solve(int[] arr) {
        int max1 = arr[0];
        int max2 = arr[1];

        for (int i : arr) {
            if (i > max1) {
                max2 = max1;
                max1 = i;
            } else if (i > max2) {
                max2 = i;
            }
        }

        return max2;
    }

}
