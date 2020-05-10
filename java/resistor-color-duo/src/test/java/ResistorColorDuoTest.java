import org.junit.Before;
import org.junit.Test;

import static org.junit.Assert.assertEquals;

public class ResistorColorDuoTest {
    private ResistorColorDuo resistorColorDuo;

    @Before
    public void setup() {
        resistorColorDuo = new ResistorColorDuo();
    }

    @Test
    public void testBrownAndBlack() {
        String[] input = { "brown", "black" };
        int expected = 10;
        int actual = resistorColorDuo.value(input);

        assertEquals(expected, actual);
    }

    @Test
    public void testBlueAndGrey() {
        String[] input = { "blue", "grey" };
        int expected = 68;
        int actual = resistorColorDuo.value(input);

        assertEquals(expected, actual);
    }

    @Test
    public void testYellowAndViolet() {
        String[] input = { "yellow", "violet" };
        int expected = 47;
        int actual = resistorColorDuo.value(input);

        assertEquals(expected, actual);
    }

    @Test
    public void testOrangeAndOrange() {
        String[] input = { "orange", "orange" };
        int expected = 33;
        int actual = resistorColorDuo.value(input);

        assertEquals(expected, actual);
    }

    @Test
    public void testIgnoreAdditionalColors() {
        String[] input = { "green", "brown", "orange" };
        int expected = 51;
        int actual = resistorColorDuo.value(input);

        assertEquals(expected, actual);
    }
}
