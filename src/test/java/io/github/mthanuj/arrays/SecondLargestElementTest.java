package io.github.mthanuj.arrays;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

class SecondLargestElementTest {

    private SecondLargestElement secondLargestElement = new SecondLargestElement();

    @Test
    void test1() {
        int[] arr = { 4, 7, 2, 3, 1, 5, 6, 8, 9, 0 };

        int expected = 8;
        int actual = secondLargestElement.solve(arr);

        assertEquals(expected, actual);
    }

    @Test
    void test2() {
        int[] arr = { 1, 2, 3, 4, 5 };

        int expected = 4;
        int actual = secondLargestElement.solve(arr);

        assertEquals(expected, actual);
    }

}
