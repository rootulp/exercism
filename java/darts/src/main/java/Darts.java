class Darts {

	private static final int RADIUS_OF_OUTER_CIRCLE = 10;
	private static final int RADIUS_OF_MIDDLE_CIRCLE = 5;
	private final double x;
	private final double y;

    Darts(double x, double y) {
		this.x = x;
		this.y = y;
	}

    int score() {
		if(isInsideMiddleCircle()) {
			return 5;
		}
		if(isInsideOutserCircle()) {
			return 1;
		}
		return 0;
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
