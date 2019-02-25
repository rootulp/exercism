class TreeNode(object):
    def __init__(self, data, left=None, right=None):
        self.data = data
        self.left = left
        self.right = right

    def __str__(self):
        fmt = 'TreeNode(data={}, left={}, right={})'
        return fmt.format(self.data, self.left, self.right)


class BinarySearchTree(object):
    def __init__(self, tree_data):
        self.head = None
        for node_data in tree_data:
            self.insert(node_data)

    def data(self):
        return self.head

    def sorted_data(self):
        return [node.data for node in self.in_order_traversal(self.head)]

    def in_order_traversal(self, node):
        if node.left:
            self.in_order_traversal(node.left)

        yield node

        if node.right:
            self.in_order_traversal(node.right)


    def insert(self, node_data):
        node_to_insert = TreeNode(node_data)
        if self.head is None:
            self.head = node_to_insert
        else:
            self.insert_node(node_to_insert, self.head)

    def insert_node(self, node_to_insert, position):
        if node_to_insert.data <= position.data:
            if position.left is not None:
                self.insert_node(node_to_insert, position.left)
            else:
                position.left = node_to_insert
        elif node_to_insert.data > position.data:
            if position.right is not None:
                self.insert_node(node_to_insert, position.right)
            else:
                position.right = node_to_insert




