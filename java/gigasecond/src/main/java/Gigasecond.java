import java.time.LocalDate;
import java.time.LocalDateTime;

public class Gigasecond {

  private LocalDateTime now;
  private static long GIGASECONDS = 1000000000L;

  public Gigasecond(LocalDate localDate) {
    this.now = localDate.atStartOfDay();
  }

  public Gigasecond(LocalDateTime localDateTime) {
    this.now = localDateTime;
  }

  public LocalDateTime getDate() {
    return now.plusSeconds(GIGASECONDS);
  }

}
