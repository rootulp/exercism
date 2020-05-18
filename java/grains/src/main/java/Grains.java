import java.math.BigDecimal;
import java.math.BigInteger;

class Grains {

    BigInteger grainsOnSquare(final int square) {
		return BigDecimal.valueOf(Math.pow(2, square - 1)).toBigInteger();
    }

    BigInteger grainsOnBoard() {
        throw new UnsupportedOperationException("Delete this statement and write your own implementation.");
    }

}
