public class Twofer {

  public String twofer(String name) {
    if (noNameGiven(name)) {
      return "One for you, one for me.";
    }
    return String.format("One for %s, one for me.", name);
  }

  public boolean noNameGiven(String name) {
    return name == null;
  }
}