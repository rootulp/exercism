import java.util.HashMap;
import java.util.Map;

class Verse {

  private static final Map<Integer, String> NUMBER_TO_ORDINAL = numberToOrdinal();
  private static final Map<Integer, String> NUMBER_TO_GIFT = numberToGift();

  static String toString(int verseIndex) {
    return prefixForVerse(verseIndex) + giftsForVerse(verseIndex);
  }

  private static String prefixForVerse(int verseIndex) {
    return String.format("On the %s day of Christmas my true love gave to me, ", ordinal(verseIndex));
  }

  private static String giftsForVerse(int verseIndex) {
    StringBuilder giftsForVerse = new StringBuilder();
    for (int currentVerse = verseIndex; currentVerse > 0; currentVerse -= 1) {
      if (verseIndex > 1 && currentVerse == 1) {
        giftsForVerse.append("and ");
      }
      giftsForVerse.append(giftsForNumber(currentVerse));
    }
    giftsForVerse.append("\n");
    return giftsForVerse.toString();
  }


  private static String ordinal(int number) {
    return NUMBER_TO_ORDINAL.get(number);
  }

  private static String giftsForNumber(int number) {
    if (number > 1) {
      return NUMBER_TO_GIFT.get(number) + ", ";
    }
    return NUMBER_TO_GIFT.get(number);
  }

  private static Map<Integer, String> numberToGift() {
    Map<Integer, String> numberToGift = new HashMap<>();
    numberToGift.put(1, "a Partridge in a Pear Tree.");
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

}
