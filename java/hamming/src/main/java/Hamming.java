public class Hamming {

  public static Integer compute(String a, String b) {
    if (a.length() != b.length()) {
      throw new IllegalArgumentException("strands are unequal in length");
    }

    int total = 0;
    for (int i = 0, n = a.length(); i < n; i++) {
      if (a.charAt(i) != b.charAt(i)) {
        total += 1;
      }
    }
    return total;
  }
}
