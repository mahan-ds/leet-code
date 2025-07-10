class Solution(object):
    def searchRange(self, nums, target):
        l, r = 0, len(nums) - 1
        first = -1
        while l <= r:
            mid = (l + r) // 2
            if nums[mid] < target:
                l = mid + 1
            else:
                if nums[mid] == target:
                    first = mid
                r = mid - 1
                
                
        l, r = 0, len(nums) - 1
        last = -1
        while l <= r:
            mid = (l + r) // 2
            if nums[mid] > target:
                r = mid - 1
            else:
                if nums[mid] == target:
                    last = mid
                l = mid + 1
                
                
        return [first, last]
                
                
                    
            


a = Solution().searchRange([5, 7, 7, 8, 8, 10], 6)
print(a)
