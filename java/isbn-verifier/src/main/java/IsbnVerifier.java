class IsbnVerifier {

	static boolean isValidIsbn(String stringToVerify) {
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

    boolean isValid(String stringToVerify) {
		return isValidIsbn(stringToVerify);
    }

}
