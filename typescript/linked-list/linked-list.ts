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
            const lastNode = this.getLastNode(this.head);
            lastNode.next = newNode;
        }
    }
    public pop(): T {
        if (this.head == null) {
            throw new Error("Cannot call pop on an empty linked list")
        } else if (this.head.next == null) {
            const node = this.head;
            this.head = null;
            return node.data;
        } else {
            const secondToLastNode = this.getSecondToLastNode(this.head);
            const lastNode = secondToLastNode.next;

            if (lastNode == null) {
                throw new Error("Last node is null")
            } else {
                secondToLastNode.next = null;
                return lastNode.data;
            }
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

    private getLastNode(node: Node<T>): Node<T> {
        let current = node;
        while (current.next != null) {
            current = current.next;
        }
        return current;
    }

    private getSecondToLastNode(node: Node<T>): Node<T> {
        let current = node;
        while (current.next != null && current.next.next != null) {
            current = current.next;
        }
        return current;
    }
}
