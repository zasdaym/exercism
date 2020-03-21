class Matrix:
    def __init__(self, matrix_string):
        self.matrix = []
        rows = matrix_string.split("\n")
        for r in rows:
            self.matrix.append(list(map(int, r.split(" "))))

    def row(self, index):
        return self.matrix[index-1]

    def column(self, index):
        result = []
        for matrix_row in self.matrix:
            result.append(matrix_row[index-1])
        return result
