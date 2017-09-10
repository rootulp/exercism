import java.util.Arrays;
import java.util.Collections;
import java.util.List;

public class KindergartenGarden {

  private static String[] DEFAULT_STUDENTS = { "Alice", "Bob", "Charlie", "David", "Eve", "Fred", "Ginny", "Harriet", "Ileana", "Joseph", "Kincaid", "Larry"};

  private List<String> students;
  private char[][] plants;

  public KindergartenGarden(String diagram) {
    this(diagram, DEFAULT_STUDENTS);
  }

  public KindergartenGarden(String diagram, String[] students) {
    this.plants = generatePlantGrid(diagram);
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
    return Arrays.asList(getPlantAt(0,index * 2),
                         getPlantAt(0, index * 2 + 1),
                         getPlantAt(1, index * 2),
                         getPlantAt(1, index * 2 + 1));
  }

  private Plant getPlantAt(int row, int col) {
    return Plant.getPlant(plants[row][col]);
  }

  private char[][] generatePlantGrid(String diagram) {
    String[] rows = diagram.split("\n");
    char[][] plants = new char[2][rows[0].length()];

    for (int i = 0; i < rows.length; i++) {
      plants[i] = rows[i].toCharArray();
    }

    return plants;
  }

}
