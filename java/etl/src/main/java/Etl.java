import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class Etl {

  public Map<String, Integer> transform(Map<Integer, List<String>> oldMap) {
    Map<String, Integer> newMap = new HashMap<String, Integer>();

    oldMap.forEach(
        (val, list) -> {
          list.forEach(
              (character) -> {
                if (!newMap.containsKey(character)) {
                  newMap.put(character.toLowerCase(), val);
                }
              });
        });

    return newMap;
  }
}
