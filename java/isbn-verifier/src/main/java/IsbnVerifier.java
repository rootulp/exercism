class IsbnVerifier {

	static boolean isValidIsbn(String stringToVerify) {
		String parsed = removeHyphens(stringToVerify);
		Integer value = computeValueForIsbn(parsed);
		return value % 11 == 0;
	}

	static Integer computeValueForIsbn(String isbn) {
		String reversed = new StringBuilder(isbn).reverse().toString();
		int value = 0;
		for (int i=0; i<isbn.length(); i++) {
			value += Integer.parseInt(String.valueOf(reversed.charAt(i))) * (i + 1);
		}
		return value;
	}

	static String removeHyphens(String s) {
		return s.replaceAll("-", "");
	}

    boolean isValid(String stringToVerify) {
		return isValidIsbn(stringToVerify);
    }

}
