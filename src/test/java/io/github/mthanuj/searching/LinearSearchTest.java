package io.github.mthanuj.searching;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

class LinearSearchTest {

    private LinearSearch linearSearch = new LinearSearch();

    @Test
    void test1() {
        int[] arr = { 1, 2, 3, 4, 5, 6 };
        int target = 4;

        int expected = 3;
        int actual = linearSearch.solve(arr, target);

        assertEquals(expected, actual);
    }

    @Test
    void test2() {
        int[] arr = { 5, 2, 4, 1, 3 };
        int target = 1;

        int expected = 3;
        int actual = linearSearch.solve(arr, target);

        assertEquals(expected, actual);
    }

    @Test
    void test3() {
        int[] arr = { 5, 2, 4, 1, 3 };
        int target = 10;

        int expected = -1;
        int actual = linearSearch.solve(arr, target);

        assertEquals(expected, actual);
    }

}
