import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.List;

public class KindergartenGarden {

  private static String[] DEFAULT_STUDENTS = { "Alice", "Bob", "Charlie", "David", "Eve", "Fred", "Ginny", "Harriet", "Ileana", "Joseph", "Kincaid", "Larry"};

  private List<Plant> plants;

  private String diagram;
  private List<String> students;
  private char[][] plantGrid;

  public KindergartenGarden(String diagram) {
    this(diagram, DEFAULT_STUDENTS);
  }

  public KindergartenGarden(String diagram, String[] students) {
    this.diagram = diagram;
    this.students = Arrays.asList(students);
    Collections.sort(this.students);
    generatePlantGrid();
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
    return Plant.getPlant(plantGrid[row][col]);
  }

  private void generatePlantGrid() {
    String[] rows = diagram.split("\n");
    plantGrid = new char[2][rows[0].length()];

    for (int i = 0; i < rows.length; i++) {
      plantGrid[i] = rows[i].toCharArray();
    }
  }

}
