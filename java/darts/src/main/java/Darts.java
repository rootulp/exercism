class Darts {

  private static final int RADIUS_OF_INNER_CIRCLE = 1;
  private static final int RADIUS_OF_MIDDLE_CIRCLE = 5;
  private static final int RADIUS_OF_OUTER_CIRCLE = 10;

  private final double x;
  private final double y;

  Darts(double x, double y) {
    this.x = x;
    this.y = y;
  }

  int score() {
    if (isInsideInnerCircle()) {
      return 10;
    }
    if (isInsideMiddleCircle()) {
      return 5;
    }
    if (isInsideOutserCircle()) {
      return 1;
    }
    return 0;
  }

  boolean isInsideInnerCircle() {
    return distanceToOrigin() <= RADIUS_OF_INNER_CIRCLE;
  }

  boolean isInsideMiddleCircle() {
    return distanceToOrigin() <= RADIUS_OF_MIDDLE_CIRCLE;
  }

  boolean isInsideOutserCircle() {
    return distanceToOrigin() <= RADIUS_OF_OUTER_CIRCLE;
  }

  double distanceToOrigin() {
    return Math.sqrt(Math.pow(this.x, 2) + Math.pow(y, 2));
  }
}
