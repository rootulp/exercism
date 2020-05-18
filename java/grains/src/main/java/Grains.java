import java.math.BigInteger;

class Grains {

    BigInteger grainsOnSquare(final int square) {
        if(square < 1 || square > 64 ){
			throw new IllegalArgumentException("square must be between 1 and 64");
		}
		BigInteger b1 = new BigInteger("2");
		return b1.pow(square -1);
    }

    BigInteger grainsOnBoard() {
        throw new UnsupportedOperationException("Delete this statement and write your own implementation.");
    }

}
