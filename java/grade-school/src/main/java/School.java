import java.util.*;

public class School {

  private Map<Integer, List<String>> db = new HashMap<Integer, List<String>>();

  public Map<Integer, List<String>> db() {
    return db;
  }

  public void add(String name, Integer grade) {
    List<String> temp = grade(grade);
    temp.add(name);
    db.put(grade, temp);
  }

  public List<String> grade(Integer grade) {
    if (db.containsKey(grade)) {
      return db.get(grade);
    }
    return new ArrayList<String>();
  }

  public Map<Integer, List<String>> sort() {
    for (Integer grade : db.keySet()) {
      List<String> curr = db.get(grade);
      Collections.sort(curr);
    }
    return db;
  }
}
