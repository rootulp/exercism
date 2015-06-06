import java.util.List;
import java.util.ArrayList;

public class PrimeFactors {

  public static List<Integer> getForNumber(double input) {
    List<Integer> factors = new ArrayList<Integer>();
    double num = input;

    for (int i = 2; i <= num; i++) {
      if (num % i == 0) {
        factors.add(i);
        num = num / i;
        i = i - 1;
      }
    }

    return factors;
  }

}
