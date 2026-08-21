class Solution:
	def checkInclusion(self, s1: str, s2: str) -> bool:
		s1_count = [0 for i in range(26)]
		for c in s1:
			s1_count[ord(c)-ord('a')]+=1
		s2_count = [0 for i in range(26)]

		i = 0
		for j in range(len(s2)):
			s2_count[ord(s2[j])-ord('a')]+=1
			if j-i+1 == len(s1):
				if s2_count == s1_count:
					return True
				else:
					s2_count[ord(s2[i])-ord('a')]-=1
					i+=1
		return False




	
        