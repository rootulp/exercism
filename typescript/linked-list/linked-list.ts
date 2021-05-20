class Node<T> {
    public next: Node<T> | null;
    public data: T;

    constructor(data: T, next?: Node<T>) {
        this.data = data;
        this.next = next ?? null;
    }
}
export default class LinkedList<T> {
    private head: Node<T> | null;

    constructor() {
        this.head = null;
    }

    public push(data: T): void {
        const newNode = new Node(data)
        if (this.head == null) {
            this.head = newNode;
        } else {
            let current = this.head
            while (current.next != null) {
                current = current.next;
            }
            current.next = newNode;
        }
    }
    public pop(): T {
        if (this.head == null) {
            throw new Error("Cannot call pop on an empty Linked List")
        } else if (this.head.next == null) {
            const node = this.head;
            this.head = null;
            return node.data;
        } else {
            console.log(this.head);
            let current = this.head
            while (current.next != null && current.next.next != null) {
                current = current.next;
            }
            const node = current.next;
            if (node == null) {
                throw new Error("Null node")
            }
            current.next = null;
            return node.data;
        }
    }
    public shift(): T {
        return 0 as any;
    }
    public unshift(element: T): void {
        console.log(`Unshifting ${element}`);
    }
    public delete(element: T): void {
        console.log(`Deleting ${element}`);
    }
    public count(): number {
        return 0;
    }
}
