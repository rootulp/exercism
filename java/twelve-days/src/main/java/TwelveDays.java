import java.util.stream.Collectors;
import java.util.stream.IntStream;

class TwelveDays {

  String sing() {
    int FIRST_VERSE = 1;
    int LAST_VERSE = 12;
    return verses(FIRST_VERSE, LAST_VERSE);
  }

  String verses(int startVerse, int endVerse) {
    return IntStream.rangeClosed(startVerse, endVerse)
            .mapToObj(this::verse)
            .collect(Collectors.joining("\n"));
  }

  String verse(int verseIndex) {
    return Verse.toString(verseIndex);
  }

}
