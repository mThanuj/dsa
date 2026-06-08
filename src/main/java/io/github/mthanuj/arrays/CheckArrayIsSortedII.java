package io.github.mthanuj.arrays;

/*
* Problem:
* Given an array of size n, write a program to check if the given array is
* sorted in (ascending / Increasing / Non-decreasing) order or not.
* If the array is sorted then return True, Else return False.
*
*/
public class CheckArrayIsSortedII {

    public boolean solve(int[] arr) {
        for (int i = 0; i < arr.length - 1; i++) {
            if (arr[i] >= arr[i + 1]) {
                return false;
            }
        }

        return true;
    }

}
