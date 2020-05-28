import java.util.Arrays;
import java.util.Comparator;
import java.util.HashMap;
import java.util.Map;

class ResistorColor {
  private static final Map<String, Integer> colorValues = initializeColorValues();

  private static Map<String, Integer> initializeColorValues() {
    final Map<String, Integer> colorValues = new HashMap<String, Integer>();
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

  int colorCode(final String color) {
    return colorValues.get(color);
  }

  String[] colors() {
    final String[] colorNames = colorValues.keySet().toArray(new String[colorValues.size()]);
    Arrays.sort(
        colorNames,
        new Comparator<String>() {

          @Override
          public int compare(final String color1, final String color2) {
            return colorValues.get(color1) - colorValues.get(color2);
          }
        });

    return colorNames;
  }
}
