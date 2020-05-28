import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class Acronym {

  static final Pattern regex = Pattern.compile("[A-Z]+[a-z]*|[a-z]+");

  public static String generate(String phrase) {
    StringBuilder sb = new StringBuilder();
    Matcher matcher = regex.matcher(phrase);
    while (matcher.find()) {
      sb.append(matcher.group().charAt(0));
    }
    return sb.toString().toUpperCase();
  }
}
