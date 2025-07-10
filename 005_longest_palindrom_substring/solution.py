class Solution(object):
    def longestPalindrome(self, s):
        word = ""
        for i in range(len(s)):
            current_word = ""
            j =  k = i
            while (j >= 0 and k < len(s) and s[j] == s[k]): 
                current_word = s[j:k+1]
                j -= 1
                k += 1
            if len(current_word) > len(word):
                word = current_word
                
            
            current_word = ""
            j = i
            k = i + 1
            while (j >= 0 and k < len(s) and s[j] == s[k]): 
                current_word = s[j:k+1]
                j -= 1
                k += 1
            if len(current_word) > len(word):
                word = current_word
                

            
        return word
    
    
    
    
a = Solution()

print(a.longestPalindrome("cbbbbd"))


