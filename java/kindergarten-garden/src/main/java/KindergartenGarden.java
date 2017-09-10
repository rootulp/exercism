import java.util.Arrays;
import java.util.Collections;
import java.util.List;
import java.util.Scanner;

public class KindergartenGarden {

  private static final String[] DEFAULT_STUDENTS = { "Alice", "Bob", "Charlie", "David", "Eve", "Fred", "Ginny", "Harriet", "Ileana", "Joseph", "Kincaid", "Larry"};

  private final Plant[][] plants;
  private final List<String> students;

  public KindergartenGarden(String diagram) {
    this(diagram, DEFAULT_STUDENTS);
  }

  public KindergartenGarden(String diagram, String[] students) {
    this.plants = parseDiagram(diagram);
    this.students = Arrays.asList(students);
    Collections.sort(this.students);
  }

  public List<Plant> getPlantsOfStudent(String student) {
    return getPlantsForIndex(getAlphabeticIndex(student));
  }

  private int getAlphabeticIndex(String student) {
    return students.indexOf(student);
  }

  private List<Plant> getPlantsForIndex(int index) {
    return Arrays.asList(plants[0][index * 2],
                         plants[0][index * 2 + 1],
                         plants[1][index * 2],
                         plants[1][index * 2 + 1]);
  }

  private Plant[][] parseDiagram(String diagram) {
    final int rows = 2;
    final int cols = diagram.length() / 2;
    Plant[][] plants = new Plant[rows][cols];
    Scanner sc = new Scanner(diagram.replaceAll("\n","")).useDelimiter("");
    for (int r = 0; r < rows; r++) {
      for (int c = 0; c < cols; c++) {
        plants[r][c] = Plant.getPlant(sc.next().charAt(0));
      }
    }
    return plants;
  }

}
