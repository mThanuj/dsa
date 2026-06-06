package io.github.mthanuj.sorting;

import static org.junit.jupiter.api.Assertions.assertArrayEquals;

import org.junit.jupiter.api.Test;

class InsertionSortTest {

    private InsertionSort insertionSort = new InsertionSort();

    @Test
    void test1() {
        int[] arr = { 5, 3, 4, 1, 2 };

        int[] expected = { 1, 2, 3, 4, 5 };
        int[] actual = insertionSort.solve(arr);

        assertArrayEquals(expected, actual);
    }

    @Test
    void test2() {
        int[] arr = { 5, 4, 3, 2, 1 };

        int[] expected = { 1, 2, 3, 4, 5 };
        int[] actual = insertionSort.solve(arr);

        assertArrayEquals(expected, actual);
    }

}
