import org.junit.Before;
import org.junit.Test;

import static org.junit.Assert.assertTrue;
import static org.junit.Assert.assertFalse;

public class IsbnVerifierTest {
    private IsbnVerifier isbnVerifier;

    @Before
    public void setUp() {
        isbnVerifier = new IsbnVerifier();
    }

    @Test
    public void validIsbnNumber() {
        assertTrue(isbnVerifier.isValid("3-598-21508-8"));
    }

    @Test
    public void invalidIsbnCheckDigit() {
        assertFalse(isbnVerifier.isValid("3-598-21508-9"));
    }

    @Test
    public void validIsbnNumberWithCheckDigitOfTen() {
        assertTrue(isbnVerifier.isValid("3-598-21507-X"));
    }

    @Test
    public void checkDigitIsACharacterOtherThanX() {
        assertFalse(isbnVerifier.isValid("3-598-21507-A"));
    }

    @Test
    public void invalidCharacterInIsbn() {
        assertFalse(isbnVerifier.isValid("3-598-P1581-X"));
    }

    @Test
    public void xIsOnlyValidAsACheckDigit() {
        assertFalse(isbnVerifier.isValid("3-598-2X507-9"));
    }

    @Test
    public void validIsbnWithoutSeparatingDashes() {
        assertTrue(isbnVerifier.isValid("3598215088"));
    }

    @Test
    public void isbnWithoutSeparatingDashesAndXAsCheckDigit() {
        assertTrue(isbnVerifier.isValid("359821507X"));
    }

    @Test
    public void isbnWithoutCheckDigitAndDashes() {
        assertFalse(isbnVerifier.isValid("359821507"));
    }

    @Test
    public void tooLongIsbnAndNoDashes() {
        assertFalse(isbnVerifier.isValid("3598215078X"));
    }

    @Test
    public void tooShortIsbn() {
        assertFalse(isbnVerifier.isValid("00"));
    }

    @Test
    public void isbnWithoutCheckDigit() {
        assertFalse(isbnVerifier.isValid("3-598-21507"));
    }

    @Test
    public void checkDigitOfXShouldNotBeUsedForZero() {
        assertFalse(isbnVerifier.isValid("3-598-21515-X"));
    }

    @Test
    public void emptyIsbn() {
        assertFalse(isbnVerifier.isValid(""));
    }

    @Test
    public void inputIsNineCharacters() {
        assertFalse(isbnVerifier.isValid("134456729"));
    }

    @Test
    public void invalidCharactersAreNotIgnored() {
        assertFalse(isbnVerifier.isValid("3132P34035"));
    }

    @Test
    public void inputIsTooLongButContainsAValidIsbn() {
        assertFalse(isbnVerifier.isValid("98245726788"));
    }
}
