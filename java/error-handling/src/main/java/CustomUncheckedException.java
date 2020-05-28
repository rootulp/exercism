class CustomUncheckedException extends RuntimeException {

  /** */
  private static final long serialVersionUID = -4414219281389811351L;

  CustomUncheckedException() {
    super();
  }

  CustomUncheckedException(String message) {
    super(message);
  }
}
