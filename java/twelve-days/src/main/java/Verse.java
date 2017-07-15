import java.util.HashMap;
import java.util.Map;
import java.util.stream.Collectors;
import java.util.stream.IntStream;

class Verse {

  private static final Map<Integer, String> NUMBER_TO_ORDINAL = numberToOrdinal();
  private static final Map<Integer, String> NUMBER_TO_GIFT = numberToGift();

  static String toString(int verseIndex) {
    return prefix(verseIndex) + body(verseIndex) + suffix(verseIndex);
  }

  private static String prefix(int verseIndex) {
    return String.format("On the %s day of Christmas my true love gave to me, ", ordinal(verseIndex));
  }

  private static String body(int verseIndex) {
    return revRangeClosed(1, verseIndex)
            .mapToObj(Verse::gifts)
            .collect(Collectors.joining(", "));
  }

  private static String suffix(int verseIndex) {
    if (verseIndex > 1) {
      return ", and a Partridge in a Pear Tree.\n";
    }
    return "a Partridge in a Pear Tree.\n";
  }

  private static String ordinal(int number) {
    return NUMBER_TO_ORDINAL.get(number);
  }

  private static String gifts(int giftIndex) {
    return NUMBER_TO_GIFT.get(giftIndex);
  }

  private static Map<Integer, String> numberToGift() {
    Map<Integer, String> numberToGift = new HashMap<>();
    numberToGift.put(2, "two Turtle Doves");
    numberToGift.put(3, "three French Hens");
    numberToGift.put(4, "four Calling Birds");
    numberToGift.put(5, "five Gold Rings");
    numberToGift.put(6, "six Geese-a-Laying");
    numberToGift.put(7, "seven Swans-a-Swimming");
    numberToGift.put(8, "eight Maids-a-Milking");
    numberToGift.put(9, "nine Ladies Dancing");
    numberToGift.put(10, "ten Lords-a-Leaping");
    numberToGift.put(11, "eleven Pipers Piping");
    numberToGift.put(12, "twelve Drummers Drumming");
    return numberToGift;
  }

  private static Map<Integer, String> numberToOrdinal() {
    Map<Integer, String> numberToOrdinal = new HashMap<>();
    numberToOrdinal.put(1, "first");
    numberToOrdinal.put(2, "second");
    numberToOrdinal.put(3, "third");
    numberToOrdinal.put(4, "fourth");
    numberToOrdinal.put(5, "fifth");
    numberToOrdinal.put(6, "sixth");
    numberToOrdinal.put(7, "seventh");
    numberToOrdinal.put(8, "eighth");
    numberToOrdinal.put(9, "ninth");
    numberToOrdinal.put(10, "tenth");
    numberToOrdinal.put(11, "eleventh");
    numberToOrdinal.put(12, "twelfth");
    return numberToOrdinal;
  }

  private static IntStream revRangeClosed(int from, int to) {
    return IntStream.range(from, to)
            .map(i -> to - i + from);
  }

}
