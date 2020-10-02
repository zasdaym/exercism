class SgfTree:
    def __init__(self, properties=None, children=None):
        self.properties = properties or {}
        self.children = children or []

    def __eq__(self, other):
        if not isinstance(other, SgfTree):
            return False
        for k, v in self.properties.items():
            if k not in other.properties:
                return False
            if other.properties[k] != v:
                return False
        for k in other.properties.keys():
            if k not in self.properties:
                return False
        if len(self.children) != len(other.children):
            return False
        for a, b in zip(self.children, other.children):
            if a != b:
                return False
        return True

    def __ne__(self, other):
        return not self == other


def parse(input_string):
    if input_string == "":
        raise ValueError("Input can't be empty")
    if input_string == "()":
        raise ValueError("Tree can't be empty")
    if input_string == ";":
        raise ValueError("Node must be in a tree")
    if input_string == "(;)":
        return SgfTree()

    input_string = input_string[2:-1]
    prop_node, *child_nodes = input_string.replace(")", "").replace(";", "(;").replace("((;", "(;").split("(;")

    return SgfTree(
        properties=create_node(prop_node),
        children=[SgfTree(create_node(child)) for child in child_nodes]
    )

def create_node(input_string):
    input_entries = input_string \
        .replace("\t", " ") \
        .replace("\\]", "¤") \
        .replace("]", "#") \
        .replace("#[", "[") \
        .split("#")

    new_node = {}
    for node_entry in filter(None, input_entries):
        node_key, *node_values = node_entry.split("[")

        if not node_key.isupper():
            raise ValueError("Node key must be uppercase")
        if len(node_values) == 0:
            raise ValueError("Key must have values associated to it")

        new_node[node_key] = [value.replace("]", "").replace("¤", "]") for value in node_values]
    return new_node
