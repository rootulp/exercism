import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;
import java.util.stream.Collectors;

class ResistorColorDuo {
  private static Map<String, Number> COLOR_VALUES = initializeColorValues();

  private static Map<String, Number> initializeColorValues() {
    final Map<String, Number> colorValues = new HashMap<String, Number>();
    colorValues.put("black", 0);
    colorValues.put("brown", 1);
    colorValues.put("red", 2);
    colorValues.put("orange", 3);
    colorValues.put("yellow", 4);
    colorValues.put("green", 5);
    colorValues.put("blue", 6);
    colorValues.put("violet", 7);
    colorValues.put("grey", 8);
    colorValues.put("white", 9);
    return colorValues;
  }

  int value(final String[] colors) {
    return Integer.parseInt(
        String.join(
            "",
            Arrays.asList(colors).stream()
                .map(s -> COLOR_VALUES.get(s))
                .map(n -> n.toString())
                .collect(Collectors.toList())
                .subList(0, 2)));
  }
}
