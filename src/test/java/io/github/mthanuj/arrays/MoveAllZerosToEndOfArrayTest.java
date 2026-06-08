package io.github.mthanuj.arrays;

import static org.junit.jupiter.api.Assertions.assertArrayEquals;

import org.junit.jupiter.api.Test;

class MoveAllZerosToEndOfArrayTest {

    private MoveAllZerosToEndOfArray moveAllZerosToEndOfArray = new MoveAllZerosToEndOfArray();

    @Test
    void test1() {
        int[] arr = { 1, 0, 2, 3, 0, 4, 0, 1 };

        int[] expected = { 1, 2, 3, 4, 1, 0, 0, 0 };
        int[] actual = moveAllZerosToEndOfArray.solve(arr);

        assertArrayEquals(expected, actual);
    }

    @Test
    void test2() {
        int[] arr = { 1, 2, 0, 1, 0, 4, 0 };

        int[] expected = { 1, 2, 1, 4, 0, 0, 0 };
        int[] actual = moveAllZerosToEndOfArray.solve(arr);

        assertArrayEquals(expected, actual);
    }

}
