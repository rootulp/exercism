public class Bob {
  public String hey(String s) {
    if (s.trim().isEmpty()) {
      return "Fine. Be that way!";
    } else if (s.equals(s.toUpperCase()) && !s.equals(s.toLowerCase())) {
      return "Whoa, chill out!";
    } else if (s.charAt(s.length() - 1) == '?') {
      return "Sure.";
    } else {
      return "Whatever.";
    }
  }
}
