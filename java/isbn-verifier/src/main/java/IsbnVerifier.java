class IsbnVerifier {

	static final char VALID_CHECK_DIGIT = 'X';

	static boolean isValidIsbn(String stringToVerify) {
		if(!isValidCheckDigit(stringToVerify)) {
			return false;
		}

		String parsed = removeHyphens(stringToVerify);
		Integer value = computeValueForIsbn(parsed);
		return value % 11 == 0;
	}

	static Integer computeValueForIsbn(String isbn) {
		String reversed = new StringBuilder(isbn).reverse().toString();
		int value = 0;
		for (int i = 0; i < isbn.length(); i++) {
			char c = reversed.charAt(i);
			value += getValueFrom(c, i);
		}
		return value;
	}

	static Integer getValueFrom(char c, int i) {
		if(i == 0 && c == 'X') {
			return 10;
		}
		return Integer.parseInt(String.valueOf(c)) * (i + 1);
	}

	static String removeHyphens(String s) {
		return s.replaceAll("-", "");
	}

	static boolean isValidCheckDigit(String s) {
		char checkDigit = s.charAt((s.length() - 1));
		return Character.isDigit(checkDigit) || checkDigit == VALID_CHECK_DIGIT;
	}

    boolean isValid(String stringToVerify) {
		return isValidIsbn(stringToVerify);
    }

}
