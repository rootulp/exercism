public class PhoneNumber {

  private static final String invalid = "0000000000";
  private final String number;

  public PhoneNumber(String dirtyNum) {
    this.number = clean(dirtyNum);
  }

  private String clean(String dirtyNum) {
    String num = dirtyNum.replaceAll("[^0-9]", "");
    if (num.length() == 10) {
      return num;
    }
    if (num.length() == 11 && num.charAt(0) == '1') {
      return num.substring(1);
    }
    return invalid;
  }

  public String pretty() {
    return "(" + getAreaCode() + ") " + getPrefixNum() + "-" + getLineNum();
  }

  public String getNumber() {
    return this.number;
  }

  public String getAreaCode() {
    return this.number.substring(0, 3);
  }

  private String getPrefixNum() {
    return this.number.substring(3, 6);
  }

  private String getLineNum() {
    return this.number.substring(6);
  }
}
