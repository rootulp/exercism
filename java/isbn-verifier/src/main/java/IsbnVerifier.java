class IsbnVerifier {

	private static final char VALID_CHECK_DIGIT = 'X';

	/**
	 * Determine if the provided string is a valid ISBN according to the following formula
	 * (x1 * 10 + x2 * 9 + x3 * 8 + x4 * 7 + x5 * 6 + x6 * 5 + x7 * 4 + x8 * 3 + x9 * 2 + x10 * 1) mod 11 == 0
	 */
	public static boolean isValidIsbn(String stringToVerify) {
		String parsed = removeHyphens(stringToVerify);
		if(!isValidParsed(parsed)) {
			return false;
		}

		return computeValueForIsbn(parsed) % 11 == 0;
	}

	private static Integer computeValueForIsbn(String isbn) {
		String reversed = new StringBuilder(isbn).reverse().toString();
		int result = 0;
		for (int i = 0; i < isbn.length(); i++) {
			char c = reversed.charAt(i);
			result += getValueFrom(c, i);
		}
		return result;
	}

	private static Integer getValueFrom(char c, int i) {
		if(i == 0 && c == VALID_CHECK_DIGIT) {
			return 10;
		}
		return Integer.parseInt(String.valueOf(c)) * (i + 1);
	}

	private static String removeHyphens(String s) {
		return s.replaceAll("-", "");
	}

	private static boolean isValidParsed(String parsed) {
		return isValidLength(parsed) &&
			   isValidCheckDigit(parsed.charAt((parsed.length() - 1))) &&
			   doesContainOnlyDigits(parsed.substring(0, parsed.length() - 1));
	}

	private static boolean isValidCheckDigit(char checkDigit) {
		return Character.isDigit(checkDigit) || checkDigit == VALID_CHECK_DIGIT;
	}

	private static boolean isValidLength(String s) {
		return s.length() == 10;
	}

	private static boolean doesContainOnlyDigits(String s) {
		return s.codePoints().mapToObj(c -> (char) c).allMatch(c -> Character.isDigit(c));
	}

    public boolean isValid(String stringToVerify) {
		return isValidIsbn(stringToVerify);
    }
}
