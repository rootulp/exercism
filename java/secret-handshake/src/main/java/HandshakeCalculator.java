import java.util.*;

final class HandshakeCalculator {

  private static Map<Integer, Signal> BINARY_TO_ACTION = new HashMap<>();

  static {
    BINARY_TO_ACTION.put(1, Signal.WINK);
    BINARY_TO_ACTION.put(2, Signal.DOUBLE_BLINK);
    BINARY_TO_ACTION.put(4, Signal.CLOSE_YOUR_EYES);
    BINARY_TO_ACTION.put(8, Signal.JUMP);
  }

  public List<Signal> calculateHandshake(int decimal) {
    List<Signal> handshake = new ArrayList<>();

    for (Map.Entry<Integer, Signal> entry : BINARY_TO_ACTION.entrySet()) {
      if ((decimal & entry.getKey()) != 0) {
        handshake.add(entry.getValue());
      }
    }

    if (shouldReverseHandshake(decimal)) {
      handshake = reverseHandshake(handshake);
    }

    return handshake;
  }

  private boolean shouldReverseHandshake(int decimal) {
    return (decimal & 16) != 0;
  }

  private List<Signal> reverseHandshake(List<Signal> handshake) {
    Collections.reverse(handshake);
    return handshake;
  }
}
