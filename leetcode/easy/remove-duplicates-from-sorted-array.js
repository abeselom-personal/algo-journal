/**
 * @param {number[]} nums
 * @return {number}
 */
var removeDuplicates = function(nums) {
    const num = new Set(nums)
    return [...num]
};

console.log(removeDuplicates([0,0,1,1,1,2,2,3,3,4]))
