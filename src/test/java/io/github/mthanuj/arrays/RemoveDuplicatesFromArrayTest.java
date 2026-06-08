package io.github.mthanuj.arrays;

import static org.junit.jupiter.api.Assertions.assertArrayEquals;
import static org.junit.jupiter.api.Assertions.assertEquals;

import java.util.Arrays;

import org.junit.jupiter.api.Test;

import io.github.mthanuj.arrays.RemoveDuplicatesFromArray.Result;

class RemoveDuplicatesFromArrayTest {

    private RemoveDuplicatesFromArray removeDuplicatesFromArray = new RemoveDuplicatesFromArray();

    @Test
    void test1() {
        int[] arr = { 1, 1, 2, 2, 2, 3, 3 };

        Result expected = new Result(new int[] { 1, 2, 3 }, 3);

        Result actual = removeDuplicatesFromArray.solve(arr);

        assertEquals(expected.k(), actual.k());
        assertArrayEquals(expected.arr(), Arrays.copyOf(actual.arr(), actual.k()));
    }

}
