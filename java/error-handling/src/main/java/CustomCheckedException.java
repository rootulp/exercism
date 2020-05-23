class CustomCheckedException extends Exception {

    /**
	 *
	 */
	private static final long serialVersionUID = 3900554992987389814L;

	public CustomCheckedException() {
		super();
    }

    CustomCheckedException(String message) {
        throw new UnsupportedOperationException("Delete this statement and write your own implementation.");
    }

}
