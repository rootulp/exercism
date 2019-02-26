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
        return [node.data for node in self.inorder_traversal(self.head)]

    def insert(self, node_data):
        new_node = TreeNode(node_data)
        if self.head is None:
            self.head = new_node
        else:
            self.insert_node_at(new_node, self.head)

    def insert_node_at(self, new_node, position):
        if new_node.data <= position.data:
            if position.left is None:
                position.left = new_node
            else:
                self.insert_node_at(new_node, position.left)
        elif new_node.data > position.data:
            if position.right is None:
                position.right = new_node
            else:
                self.insert_node_at(new_node, position.right)

    def inorder_traversal(self, node):
        if node.left:
            yield from self.inorder_traversal(node.left)
        yield node
        if node.right:
            yield from self.inorder_traversal(node.right)
