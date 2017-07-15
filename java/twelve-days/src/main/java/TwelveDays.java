import java.util.HashMap;
import java.util.Map;
import java.util.stream.Collectors;
import java.util.stream.IntStream;

class TwelveDays {

  private int FIRST_VERSE = 1;
  private int LAST_VERSE = 12;
  private Map<Integer, String> NUMBER_TO_ORDINAL = numberToOrdinal();
  private Map<Integer, String> NUMBER_TO_GIFT = numberToGift();

  String sing() {
    return verses(FIRST_VERSE, LAST_VERSE);
  }

  String verses(int startVerse, int endVerse) {
    return IntStream.rangeClosed(startVerse, endVerse)
            .mapToObj(currentVerse -> verse(currentVerse))
            .collect(Collectors.joining("\n"));
  }

  String verse(int verseNumber) {
    return prefixForVerse(verseNumber) + giftsForVerse(verseNumber);
  }

  String prefixForVerse(int verseNumber) {
    return String.format("On the %s day of Christmas my true love gave to me, ", ordinal(verseNumber));
  }

  String giftsForVerse(int verseNumber) {
    StringBuilder giftsForVerse = new StringBuilder();
    for (int currentVerse = verseNumber; currentVerse > 0; currentVerse -= 1) {
      if (verseNumber > 1 && currentVerse == 1) {
        giftsForVerse.append("and ");
      }
      giftsForVerse.append(giftsForNumber(currentVerse));
    }
    giftsForVerse.append("\n");
    return giftsForVerse.toString();
  }


  String ordinal(int number) {
    return NUMBER_TO_ORDINAL.get(number);
  }

  String giftsForNumber(int number) {
    if (number > 1) {
      return NUMBER_TO_GIFT.get(number) + ", ";
    }
    return NUMBER_TO_GIFT.get(number);
  }

  private Map<Integer, String> numberToGift() {
    Map<Integer, String> numberToGift = new HashMap<Integer, String>();
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

  private Map<Integer, String> numberToOrdinal() {
    Map<Integer, String> numberToOrdinal = new HashMap<Integer, String>();
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
