class Solution:
    def rotate(self, matrix: List[List[int]]) -> None:
        # transpose: 

        for i in range(len(matrix)):
            for j in range(i, len(matrix[0])):
                matrix[i][j], matrix [j][i] = matrix [j][i], matrix [i][j]
        for i in range(len(matrix)):
             matrix[i] = matrix[i][::-1]

        
