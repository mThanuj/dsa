package io.github.mthanuj.arrays;

import java.util.ArrayList;
import java.util.List;

/*
* Problem:
* Given two sorted arrays, arr1, and arr2 of size n and m. Find the union of two sorted arrays.
* The union of two arrays can be defined as the common and distinct elements in the two arrays.
*
* NOTE: Elements in the union should be in ascending order.
*
*/
public class UnionOfTwoSortedArrays {

    public int[] solve(int[] arr1, int[] arr2) {
        List<Integer> result = new ArrayList<>();

        int i = 0;
        int j = 0;

        while (i < arr1.length && j < arr2.length) {
            if (arr1[i] < arr2[j]) {
                addIfNew(result, arr1[i]);
                i++;
            } else {
                addIfNew(result, arr2[j]);
                j++;
            }
        }

        while (i < arr1.length) {
            addIfNew(result, arr1[i]);
            i++;
        }

        while (j < arr2.length) {
            addIfNew(result, arr2[j]);
            j++;
        }

        return result.stream().mapToInt(Integer::intValue).toArray();
    }

    private void addIfNew(List<Integer> list, int value) {
        if (list.isEmpty() || list.getLast() != value) {
            list.add(value);
        }
    }

}
