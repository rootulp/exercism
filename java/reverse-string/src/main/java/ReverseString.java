class ReverseString {

  String reverse(String inputString) {
    StringBuilder builder = new StringBuilder(inputString);
    return builder.reverse().toString();
  }
}
