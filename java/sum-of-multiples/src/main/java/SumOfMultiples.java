import java.util.HashSet;
import java.util.Set;

class SumOfMultiples {

  private int number;
  private int[] set;

  SumOfMultiples(int number, int[] set) {
    this.number = number;
    this.set = set;
  }

  int getSum() {
    Set<Integer> multiplesOfSetOfNumbersUpToMaxNumber =
        multiplesOfSetOfNumbersUpToMaxNumber(number, set);
    return multiplesOfSetOfNumbersUpToMaxNumber.stream().mapToInt(Integer::intValue).sum();
  }

  private Set<Integer> multiplesOfSetOfNumbersUpToMaxNumber(int maxNumber, int[] set) {
    Set<Integer> multiples = new HashSet<>();

    for (int number : set) {
      multiples.addAll(multiplesOfNumberUpToMaxNumber(maxNumber, number));
    }

    return multiples;
  }

  private Set<Integer> multiplesOfNumberUpToMaxNumber(int maxNumber, int number) {
    Set<Integer> multiples = new HashSet<>();

    int multiplier = 1;
    while (number * multiplier < maxNumber) {
      multiples.add(number * multiplier);
      multiplier += 1;
    }

    return multiples;
  }
}
