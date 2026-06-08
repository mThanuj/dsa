package io.github.mthanuj.arrays;

import static org.junit.jupiter.api.Assertions.assertArrayEquals;

import org.junit.jupiter.api.Test;

class UnionOfTwoSortedArraysTest {

    private UnionOfTwoSortedArrays unionOfTwoSortedArrays = new UnionOfTwoSortedArrays();

    @Test
    void test1() {
        int[] arr1 = { 1, 2, 3, 4, 5 };
        int[] arr2 = { 2, 3, 4, 4, 5 };

        int[] expected = { 1, 2, 3, 4, 5 };
        int[] actual = unionOfTwoSortedArrays.solve(arr1, arr2);

        assertArrayEquals(expected, actual);
    }

    @Test
    void test2() {
        int[] arr1 = { 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 };
        int[] arr2 = { 2, 3, 4, 4, 5, 11, 12 };

        int[] expected = { 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12 };
        int[] actual = unionOfTwoSortedArrays.solve(arr1, arr2);

        assertArrayEquals(expected, actual);
    }

}
