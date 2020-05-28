public class SpaceAge {

  private final double seconds;

  private enum Planet {
    EARTH(31557600),
    MERCURY(0.2408467),
    VENUS(0.61519726),
    MARS(1.8808158),
    JUPITER(11.862615),
    SATURN(29.447498),
    URANUS(84.016846),
    NEPTUNE(164.79132);

    // op -> Orbital Period relative to Earth years
    private final double op;

    Planet(double op) {
      this.op = op;
    }

    private final double op() {
      return this.op;
    }
  }

  public SpaceAge(double seconds) {
    this.seconds = seconds;
  }

  public double getSeconds() {
    return this.seconds;
  }

  public double onEarth() {
    return getSeconds() / Planet.EARTH.op();
  }

  public double onMercury() {
    return onEarth() / Planet.MERCURY.op();
  }

  public double onVenus() {
    return onEarth() / Planet.VENUS.op();
  }

  public double onMars() {
    return onEarth() / Planet.MARS.op();
  }

  public double onJupiter() {
    return onEarth() / Planet.JUPITER.op();
  }

  public double onSaturn() {
    return onEarth() / Planet.SATURN.op();
  }

  public double onUranus() {
    return onEarth() / Planet.URANUS.op();
  }

  public double onNeptune() {
    return onEarth() / Planet.NEPTUNE.op();
  }
}
