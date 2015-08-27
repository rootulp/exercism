import java.util.List;
import java.util.ArrayList;

public final class Sieve {

  private Integer limit;
  private List<Integer> primes;

  public Sieve(Integer limit) {
    this.limit = limit;
    primes = findPrimes();
  }

  public List<Integer> getPrimes() {
    return primes;
  }

  private List<Integer> findPrimes() {
    List<Integer> primes = new ArrayList<Integer>();
    List<Integer> possible = findPossible();

    while (!possible.isEmpty()) {
      Integer curr = possible.remove(0);
      possible = removeMultiples(curr, possible);
      primes.add(curr);
    }

    return primes;
  }

  private List<Integer> findPossible() {
    List<Integer> possible = new ArrayList<Integer>();

    for (int i = 2; i <= limit; i++) {
      possible.add(i);
    }

    return possible;
  }

  private List<Integer> removeMultiples(Integer curr, List<Integer> possible) {
    List<Integer> multiples = new ArrayList<Integer>();
    Integer multiplier = 2;

    while (!possible.isEmpty() &&
            curr * multiplier <= possible.get(possible.size() - 1)) {
      multiples.add(curr * multiplier);
      multiplier += 1;
    }

    possible.removeAll(multiples);
    return possible;
  }

}
