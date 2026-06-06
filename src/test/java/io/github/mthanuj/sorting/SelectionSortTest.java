package io.github.mthanuj.sorting;

import static org.junit.Assert.assertArrayEquals;

import org.junit.Test;

public class SelectionSortTest {

    private SelectionSort selectionSort = new SelectionSort();

    @Test
    public void test1() {
        int[] arr = { 5, 3, 4, 1, 2 };

        int[] expected = { 1, 2, 3, 4, 5 };
        int[] actual = selectionSort.solve(arr);

        assertArrayEquals(expected, actual);
    }

    @Test
    public void test2() {
        int[] arr = { 5, 4, 3, 2, 1 };

        int[] expected = { 1, 2, 3, 4, 5 };
        int[] actual = selectionSort.solve(arr);

        assertArrayEquals(expected, actual);
    }

}
