import java.util.ArrayList;
import java.util.List;

public class Raindrops {

  public static String convert(int input) {
    List<Integer> factors = getForNumber(input);

    StringBuilder sb = new StringBuilder();
    if (factors.contains(3)) {
      sb.append("Pling");
    }
    if (factors.contains(5)) {
      sb.append("Plang");
    }
    if (factors.contains(7)) {
      sb.append("Plong");
    }
    if (sb.toString().isEmpty()) {
      sb.append(Integer.toString(input));
    }

    return sb.toString();
  }

  public static List<Integer> getForNumber(int input) {
    List<Integer> factors = new ArrayList<Integer>();
    int num = input;

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
