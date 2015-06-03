public class Triangle {

  private final double sideA;
  private final double sideB;
  private final double sideC;

  public Triangle(double sideA, double sideB, double sideC) throws TriangleException {
    if (triangleInvalid(sideA, sideB, sideC)) { throw new TriangleException(); }

    this.sideA = sideA;
    this.sideB = sideB;
    this.sideC = sideC;
  }

  public TriangleKind getKind() {
    if (sideA == sideB && sideB == sideC && sideA == sideC) {
      return TriangleKind.EQUILATERAL;
    } else if (sideA == sideB || sideA == sideC || sideB == sideC) {
      return TriangleKind.ISOSCELES;
    } else {
      return TriangleKind.SCALENE;
    }
  }

  private boolean triangleInvalid(double sideA, double sideB, double sideC) {
    return triangleInequality(sideA, sideB, sideC) ||
           triangleInequality(sideC, sideB, sideA) ||
           triangleInequality(sideA, sideC, sideB);
  }

  private boolean triangleInequality(double side1, double side2, double side3) {
    return (side1 + side2 <= side3);
  }

}
