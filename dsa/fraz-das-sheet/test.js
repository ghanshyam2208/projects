/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number}
 */
var threeSumClosest = function (nums, target) {
  nums.sort((a, b) => a - b);
  // if (target < 0) target *= -1;
  const len = nums.length;
  let gap = Infinity;
  let retSum = undefined;
  for (let i = 0; i < len; i++) {
    let start = i + 1;
    let end = len - 1;

    while (start < end) {
      const loSum = nums[i] + nums[start] + nums[end];

      if (loSum > target) {
        if (loSum - target < gap) {
          gap = loSum - target;
          retSum = loSum;
        }

        end -= 1;
      } else if (target > loSum) {
        if (target - loSum < gap) {
          gap = target - loSum;
          retSum = loSum;
        }
        start += 1;
      } else {
        return target - loSum;
      }
    }
  }
  return retSum;
};

console.log(threeSumClosest([1, 1, 1, 0], -100));
