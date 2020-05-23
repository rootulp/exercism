class CustomUncheckedException extends RuntimeException {

    /**
	 *
	 */
	private static final long serialVersionUID = -4414219281389811351L;

	CustomUncheckedException() {
        throw new UnsupportedOperationException("Delete this statement and write your own implementation.");
    }

    CustomUncheckedException(String message) {
        throw new UnsupportedOperationException("Delete this statement and write your own implementation.");
    }

}
