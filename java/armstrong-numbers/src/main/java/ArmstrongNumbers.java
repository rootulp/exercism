import java.util.ArrayList;
import java.util.Collections;
import java.util.List;
import java.util.stream.Collectors;

class ArmstrongNumbers {

    boolean isArmstrongNumber(int numberToCheck) {
		List<Integer> digits = getDigits(numberToCheck);
		double sumOfDigitsRaisedToPower = digits.stream().map(digit -> {
			return Math.pow(digit, numberOfDigits(numberToCheck));
		}).collect(Collectors.summingDouble(x -> x));
		return sumOfDigitsRaisedToPower == numberToCheck;
	}

	int numberOfDigits(int numberToCheck) {
		return Integer.toString(numberToCheck).length();
	}

	/**
	 * Split an integer into its individual digits
	 * NOTE: digits order is maintained - i.e. Least significant digit is at index[0]
	 * @param num positive integer
	 * @return array of digits
	 */
	public static List<Integer> getDigits(int num) {
		if (num < 0) { return new ArrayList<>(0); }
		List<Integer> digits = new ArrayList<Integer>();
		collectDigits(num, digits);
		Collections.reverse(digits);
		return digits;
	}

	private static void collectDigits(int num, List<Integer> digits) {
		if(num / 10 > 0) {
			collectDigits(num / 10, digits);
		}
		digits.add(num % 10);
	}
}
