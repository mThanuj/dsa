package io.github.mthanuj.arrays;

import java.util.Arrays;

/*
 * Problem:
 * Given an integer array sorted in non-decreasing order,
 * remove the duplicates in place such that each unique element appears only once.
 * The relative order of the elements should be kept the same.
 * If there are k elements after removing the duplicates,
 * then the first k elements of the array should hold the final result.
 * It does not matter what you leave beyond the first k elements.
 *
 */
public class RemoveDuplicatesFromArray {

    public record Result(
            int[] arr,
            int k) {

        @Override
        public String toString() {
            return "Result {arr=" + Arrays.toString(arr) + ", k=" + k + "}";
        }

        @Override
        public int hashCode() {
            final int prime = 31;
            int result = 1;
            result = prime * result + Arrays.hashCode(arr);
            result = prime * result + k;
            return result;
        }

        @Override
        public boolean equals(Object obj) {
            if (this == obj) {
                return true;
            }

            if (obj == null) {
                return false;
            }

            if (getClass() != obj.getClass()) {
                return false;
            }

            Result other = (Result) obj;

            return !(!Arrays.equals(arr, other.arr) || k != other.k);
        }
    }

    public Result solve(int[] arr) {
        int left = 0;
        int right = 1;

        while (right < arr.length) {
            if (arr[right] != arr[left]) {
                arr[++left] = arr[right];
            }

            right++;
        }

        return new Result(arr, left + 1);
    }

}
