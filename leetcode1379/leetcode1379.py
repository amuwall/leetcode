import typing


class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    def getTargetCopy(self, original: TreeNode, cloned: TreeNode, target: TreeNode) -> TreeNode:
        return self.search_node(cloned, target)

    def search_node(self, node: TreeNode, target: TreeNode) -> typing.Optional[TreeNode]:
        if node is None:
            return None
        if node.val == target.val:
            return node

        left = self.search_node(node.left, target)
        if left is not None:
            return left
        right = self.search_node(node.right, target)
        if right is not None:
            return right

        return None
