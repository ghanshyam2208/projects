//{ Driver Code Starts
//Initial Template for javascript

"use strict";

function main() {
  let obj = new Solution();
  let res = obj.isSubset([1, 2, 3, 4, 5, 1, 1, 1], [1, 2, 3, 1], 8, 4);
  console.log(res);
}

// } Driver Code Ends

//User function Template for javascript

/**
 * @param {number[]} a1
 * @param {number[]} a2
 * @param {number} n
 * @param {number} m
 * @returns {string}
 */

class Solution {
  isSubset(a1, a2, n, m) {
    //code here
    const a1Map = {};
    const a2Map = {};

    for (let i = 0; i < a1.length; i++) {
      if (a1Map[a1[i]] === undefined) {
        a1Map[a1[i]] = 1;
      } else {
        a1Map[a1[i]] += 1;
      }
    }

    for (let i = 0; i < a2.length; i++) {
      if (a2Map[a2[i]] === undefined) {
        a2Map[a2[i]] = 1;
      } else {
        a2Map[a2[i]] += 1;
      }
    }

    console.log(a1Map);
    console.log(a2Map);
    let flag = true;
    for (let i = 0; i < a2.length; i++) {
      if (a1Map[a2[i]] === undefined) {
        return "No";
      }
      const a1o = a1Map[a2[i]];
      const a2o = a2Map[a2[i]];
      if (a1o < a2o) {
        return "No";
      }
    }
    return "Yes";
  }
}

main();
