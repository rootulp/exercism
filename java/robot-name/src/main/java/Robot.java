import java.util.Random;

public class Robot {

  private String name;

  public Robot() {
    this.name = generateName();
  }

  public String getName() {
    return this.name;
  }

  public void reset() {
    this.name = generateName();
  }

  private String generateName() {
    StringBuilder sb = new StringBuilder();

    for (int i = 0; i < 2; i++) { sb.append(randChar()); }
    for (int i = 0; i < 3; i++) { sb.append(randDigit()); }

    return sb.toString();
  }

  private char randChar() {
    Random r = new Random();
    int n = r.nextInt(26);
    char randChar = (char) (n + 65);
    return randChar;
  }

  private char randDigit() {
    Random r = new Random();
    int n = r.nextInt(10);
    char randDigit = (char) (n + 48);
    return randDigit;
  }

}
