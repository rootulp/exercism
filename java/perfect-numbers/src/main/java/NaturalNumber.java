import java.util.HashSet;
import java.util.Set;

class NaturalNumber {

  private int naturalNumber;

  public NaturalNumber(int naturalNumber) {
    if (invalidNaturalNumber(naturalNumber)) {
      throw new IllegalArgumentException(
          "You must supply a natural number (positive integer)",
          new Throwable(Integer.toString(naturalNumber)));
    }
    this.naturalNumber = naturalNumber;
  }

  public Classification getClassification() {
    int aliquotSum = aliquotSum(naturalNumber);

    if (aliquotSum > naturalNumber) {
      return Classification.ABUNDANT;
    } else if (aliquotSum < naturalNumber) {
      return Classification.DEFICIENT;
    } else {
      return Classification.PERFECT;
    }
  }

  private int aliquotSum(int number) {
    return factorsExcludingNumber(number).stream().mapToInt(Integer::intValue).sum();
  }

  private Set<Integer> factorsExcludingNumber(int number) {
    Set<Integer> factorsExcludingNumber = factorsOf(number);
    factorsExcludingNumber.remove(number);
    return factorsExcludingNumber;
  }

  private Set<Integer> factorsOf(int number) {
    Set<Integer> factors = new HashSet<Integer>();
    for (int potentialFactor = 1;
        potentialFactor < Math.ceil(Math.sqrt(number));
        potentialFactor += 1) {
      if (number % potentialFactor == 0) {
        factors.add(potentialFactor);
        factors.add(number / potentialFactor);
      }
    }
    return factors;
  }

  private boolean invalidNaturalNumber(int number) {
    return number <= 0;
  }
}
