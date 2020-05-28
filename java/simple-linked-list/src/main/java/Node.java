public class Node {

  private Integer data;
  private Node next;

  public Node(Integer data) {
    this.data = data;
  }

  public Node(Integer data, Node next) {
    this.data = data;
    this.next = next;
  }

  public Integer getData() {
    return this.data;
  }

  public Node getNext() {
    return this.next;
  }

  public void setNext(Node next) {
    this.next = next;
  }
}
