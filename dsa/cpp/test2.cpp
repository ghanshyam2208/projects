//{ Driver Code Starts
//Initial template for C++

#include<bits/stdc++.h>
using namespace std;

// Function declaration
size_t getLen(string arr[]);

class Solution{
  public:
    
    string longestCommonPrefix (string arr[], int N)
    {
        int firstSize = arr[0].size();
        string res = "";
        for (int i = 0; i < firstSize; i++) {
            cout << arr[0][i] << "\t";
            
        }

        cout << "\n";
    }
};

int main() {
    Solution sol;
    string arr[] = {"apple", "apricot", "appetizer"};
    int N = getLen(arr);

    // Call the longestCommonPrefix function
    string prefix = sol.longestCommonPrefix(arr, N);

    // Print the result
    cout << "Longest Common Prefix: " << prefix << endl;

    return 0;
}

size_t getLen (string arr[]) {
    return sizeof(arr)/ sizeof(int);
}