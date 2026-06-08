package io.github.mthanuj.arrays;

import static org.junit.jupiter.api.Assertions.assertArrayEquals;

import org.junit.jupiter.api.Test;

class LeftRotateArrayByKTest {

    private LeftRotateArrayByK leftRotateArrayByK = new LeftRotateArrayByK();

    @Test
    void test1() {
        int[] arr = { 1, 2, 3, 4, 5 };

        int[] required = { 2, 3, 4, 5, 1 };
        int[] actual = leftRotateArrayByK.solve(arr, 1);

        assertArrayEquals(required, actual);
    }

    @Test
    void test2() {
        int[] arr = { 1, 2, 3, 4, 5 };

        int[] required = { 3, 4, 5, 1, 2 };
        int[] actual = leftRotateArrayByK.solve(arr, 2);

        assertArrayEquals(required, actual);
    }

    @Test
    void test3() {
        int[] arr = { 1, 2, 3, 4, 5 };

        int[] required = { 1, 2, 3, 4, 5 };
        int[] actual = leftRotateArrayByK.solve(arr, 5);

        assertArrayEquals(required, actual);
    }

}
