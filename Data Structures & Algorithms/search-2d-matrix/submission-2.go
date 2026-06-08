func searchMatrix(matrix [][]int, target int) bool {

for i:=0;i<len(matrix);i++{

	l:=0
	r:=len(matrix[i])-1

	for l<=r{
		mid := (l+r)/2 

		if matrix[i][mid]==target{
			return true
		}
		if matrix[i][mid]<target{
			l=mid+1
		}
		if matrix[i][mid]>target{
			r=mid-1
		}
	}


}
return false

}
