import java.util.ArrayList;
import java.util.List;
import java.util.function.Function;

public class Accumulate {
  public static <T> List<T> accumulate(List<T> list, Function<T, T> f) {
    List<T> results = new ArrayList<>();

    for (T item : list) {
      results.add(f.apply(item));
    }

    return results;
  }
}
