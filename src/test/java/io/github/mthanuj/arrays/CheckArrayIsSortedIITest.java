package io.github.mthanuj.arrays;

import static org.junit.jupiter.api.Assertions.assertFalse;
import static org.junit.jupiter.api.Assertions.assertTrue;

import org.junit.jupiter.api.Test;

class CheckArrayIsSortedIITest {

    private CheckArrayIsSortedII checkArrayIsSortedII = new CheckArrayIsSortedII();

    @Test
    void test1() {
        int[] arr = { 1, 2, 3, 4, 5 };

        assertTrue(() -> checkArrayIsSortedII.solve(arr));
    }

    @Test
    void test2() {
        int[] arr = { 4, 3, 2, 1, 5 };

        assertFalse(() -> checkArrayIsSortedII.solve(arr));
    }

}
