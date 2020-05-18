import java.math.BigInteger;
import java.util.stream.IntStream;

class Grains {

	private static final int MIN_BOARD_SQUARE = 1;
	private static final int MAX_BOARD_SQUARE = 64;

    BigInteger grainsOnSquare(final int square) {
        if(square < MIN_BOARD_SQUARE || square > MAX_BOARD_SQUARE ){
            throw new IllegalArgumentException(String.format("square must be between %d and %d", MIN_BOARD_SQUARE, MAX_BOARD_SQUARE));
        }
        return new BigInteger("2").pow(square - 1);
    }

    BigInteger grainsOnBoard() {
        return IntStream.rangeClosed(MIN_BOARD_SQUARE, MAX_BOARD_SQUARE).boxed().map(n -> grainsOnSquare(n)).reduce(BigInteger.ZERO, BigInteger::add);
    }

}
