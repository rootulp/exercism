import java.util.HashMap;
import java.util.Map;
import java.util.Set;

public class Hexadecimal {

  private static final Map<Character, Integer> HEX_MAP = hexMap();
  private static final Set<Character> VALID_CHARACTERS = HEX_MAP.keySet();

  private static Map<Character, Integer> hexMap() {
    Map<Character, Integer> hexMap = new HashMap<>();
    hexMap.put('0', 0);
    hexMap.put('1', 1);
    hexMap.put('2', 2);
    hexMap.put('3', 3);
    hexMap.put('4', 4);
    hexMap.put('5', 5);
    hexMap.put('6', 6);
    hexMap.put('7', 7);
    hexMap.put('8', 8);
    hexMap.put('9', 9);
    hexMap.put('a', 10);
    hexMap.put('b', 11);
    hexMap.put('c', 12);
    hexMap.put('d', 13);
    hexMap.put('e', 14);
    hexMap.put('f', 15);
    return hexMap;
  }

  public static int toDecimal(String str) {
    if (!valid(str)) {
      return 0;
    }
    return calculateDecimal(str);
  }

  private static int calculateDecimal(String str) {
    int decimal = 0;
    for (int i = 0; i < str.length(); i++) {
      int multiplier = multiplier(i, str.length());
      int digit = digitToDecimal(str.charAt(i));
      decimal += digit * multiplier;
    }
    return decimal;
  }

  private static int multiplier(int index, int strLength) {
    return (int) Math.pow(16, (strLength - 1 - index));
  }

  private static int digitToDecimal(Character digit) {
    return HEX_MAP.get(digit);
  }

  private static boolean valid(String str) {
    for (Character c : str.toCharArray()) {
      if (!VALID_CHARACTERS.contains(c)) {
        return false;
      }
    }
    return true;
  }
}
