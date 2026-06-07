package io.github.mthanuj.arrays;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

class LargestElementTest {

    private LargestElement largestElement = new LargestElement();

    @Test
    void test1() {
        int[] arr = { 4, 7, 2, 3, 1, 5, 6, 8, 9, 0 };

        int expected = 9;
        int actual = largestElement.solve(arr);

        assertEquals(expected, actual);
    }

    @Test
    void test2() {
        int[] arr = { 1, 2, 3, 4, 5 };

        int expected = 5;
        int actual = largestElement.solve(arr);

        assertEquals(expected, actual);
    }

}
