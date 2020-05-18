import java.math.BigInteger;

class Grains {

    BigInteger grainsOnSquare(final int square) {
		BigInteger b1 = new BigInteger("2");
		return b1.pow(square -1);
    }

    BigInteger grainsOnBoard() {
        throw new UnsupportedOperationException("Delete this statement and write your own implementation.");
    }

}
