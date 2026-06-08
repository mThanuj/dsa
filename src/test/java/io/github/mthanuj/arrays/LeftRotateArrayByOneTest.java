package io.github.mthanuj.arrays;

import static org.junit.jupiter.api.Assertions.assertArrayEquals;

import org.junit.jupiter.api.Test;

class LeftRotateArrayByOneTest {

    private LeftRotateArrayByOne leftRotateArrayByOne = new LeftRotateArrayByOne();

    @Test
    void test1() {
        int[] arr = { 1, 2, 3, 4, 5 };

        int[] required = { 2, 3, 4, 5, 1 };
        int[] actual = leftRotateArrayByOne.solve(arr);

        assertArrayEquals(required, actual);
    }

    @Test
    void test2() {
        int[] arr = { 2, 3, 4, 5, 1 };

        int[] required = { 3, 4, 5, 1, 2 };
        int[] actual = leftRotateArrayByOne.solve(arr);

        assertArrayEquals(required, actual);
    }

}
