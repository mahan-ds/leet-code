



def lengthOfLongestSubstring(word:str) -> int:
    hash_map = {}
    pointer = 0
    max_length = 0
    for idx, letter in enumerate(word):
        if letter not in hash_map:
            hash_map[letter] = idx
        else:
            value = hash_map[letter]
            hash_map[letter] = idx
            if value < pointer:
                pass
            else:
                pointer = value + 1
            
            
        max_length = max(max_length, idx - pointer + 1)
        
    return max_length
    



    """
{a:0} -> {b:1}    
    
    
    
    """
print(lengthOfLongestSubstring("abcabcbb"))