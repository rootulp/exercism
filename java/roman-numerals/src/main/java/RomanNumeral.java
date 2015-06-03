import java.util.Map;
import java.util.LinkedHashMap;

public class RomanNumeral {

  private final int num;
  private static final Map<Integer, String> ROMAN = populateRoman();
  private static final Map<Integer, String> populateRoman() {
    Map<Integer, String> roman = new LinkedHashMap<Integer, String>();
    roman.put(1000, "M");
    roman.put(900, "CM");
    roman.put(500, "D");
    roman.put(400, "CD");
    roman.put(100, "C");
    roman.put(90, "XC");
    roman.put(50, "L");
    roman.put(40, "XL");
    roman.put(10, "X");
    roman.put(9, "IX");
    roman.put(5, "V");
    roman.put(4, "IV");
    roman.put(1, "I");
    return roman;
  }

  public RomanNumeral(int num) {
    this.num = num;
  }

  public String getRomanNumeral() {
    int currVal = this.num;
    String currRoman = "";

    for (Map.Entry<Integer, String > entry : ROMAN.entrySet()) {
      Integer key = entry.getKey();
      String val = entry.getValue();

      while (currVal >= key) {
        currVal -= key;
        currRoman += val;
      }
    }

    return currRoman;
  }

}
