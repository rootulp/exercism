import java.util.NoSuchElementException;

public class SimpleLinkedList {
  private Node head;

  public SimpleLinkedList() { }
  public SimpleLinkedList(Integer[] values) {
    for (Integer value: values) { push(value); }
  }

  public int size() {
    int count = 0;
    Node curr = this.head;
    while (curr != null) {
      count += 1;
      curr = curr.getNext();
    }
    return count;
  }

  public void push(int data) {
    Node newNode = new Node(data);
    if (head == null) {
      this.head = newNode;
    } else {
      getTail().setNext(newNode);
    }
  }

  public int pop() {
    if (size() == 0) { throw new NoSuchElementException(); }
    if (size() == 1) {
      Integer data = this.head.getData();
      this.head = null;
      return data;
    }
    Node curr = this.head;
    while (curr.getNext().getNext() != null) {
      curr = curr.getNext();
    }
    Integer data = curr.getNext().getData();
    curr.setNext(null);
    return data;
  }


  public void reverse() {
    reverseRecursive(this.head, null);
  }

  private Node reverseRecursive(Node curr, Node prev) {
    if (curr.getNext() == null) {
      this.head = curr;
      curr.setNext(prev);
      return null;
    }
    Node next = curr.getNext();
    curr.setNext(prev);
    reverseRecursive(next, curr);
    return this.head;
  }

  private Node getTail() {
    Node curr = this.head;
    while (curr.getNext() != null) {
      curr = curr.getNext();
    }
    return curr;
  }
}
