class MicroBlog {
    public String truncate(String input) {
		return input.substring(0, Math.min(input.length(), 5));
    }
}
