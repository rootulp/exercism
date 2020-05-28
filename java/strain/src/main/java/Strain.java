import java.util.ArrayList;
import java.util.List;
import java.util.function.Predicate;

public class Strain {

  public static <T> List<T> keep(List<T> list, Predicate<T> p) {
    List<T> results = new ArrayList<>();

    for (T item : list) {
      if (p.test(item)) {
        results.add(item);
      }
    }

    return results;
  }

  public static <T> List<T> discard(List<T> list, Predicate<T> p) {
    List<T> results = new ArrayList<>();

    for (T item : list) {
      if (!p.test(item)) {
        results.add(item);
      }
    }

    return results;
  }
}
