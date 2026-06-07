package io.github.mthanuj.sorting;

import static org.junit.jupiter.api.Assertions.assertArrayEquals;
import static org.junit.jupiter.api.Assertions.assertTimeoutPreemptively;

import java.time.Duration;
import java.util.Random;

import org.junit.jupiter.api.Test;

class MergeSortTest {

    private MergeSort mergeSort = new MergeSort();

    @Test
    void test1() {
        int[] arr = { 5, 3, 4, 1, 2 };

        int[] expected = { 1, 2, 3, 4, 5 };
        int[] actual = mergeSort.solve(arr);

        assertArrayEquals(expected, actual);
    }

    @Test
    void test2() {
        int[] arr = { 5, 4, 3, 2, 1 };

        int[] expected = { 1, 2, 3, 4, 5 };
        int[] actual = mergeSort.solve(arr);

        assertArrayEquals(expected, actual);
    }

    @Test
    void test3() {
        int n = 1_000_000;
        int[] arr = new int[n];

        Random random = new Random();

        for (int i = 0; i < n; i++) {
            arr[i] = random.nextInt();
        }

        assertTimeoutPreemptively(
                Duration.ofMillis(500),
                () -> mergeSort.solve(arr));
    }

}
