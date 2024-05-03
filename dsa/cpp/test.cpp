#include <iostream>
#include <unordered_map>
using namespace std;

// Function declaration
size_t arraySize(int arr[]);

int main() {
    std::unordered_map<int, int> myMap1;
    std::unordered_map<int, int> myMap2;

    int a1[] = {1, 2, 3, 4, 4, 5, 6};
    int a2[] = {1, 2, 4};

    int a1Size = arraySize(a1);
    int a2Size = arraySize(a2);

    for (int i = 0; i < a1Size; i++) {
        myMap1.insert
    }
        

    return 0;
}

bool checkKeyExists (const std::unordered_map<int, int> &myMap, const int &key) {
    if(myMap.find(key) != myMap.end()) {
        return true;
    }
    return false;
}

// Function definition
size_t arraySize(int arr[]) {
    // Calculate the size of the array
    size_t size = sizeof(arr) / sizeof(arr[0]);
    return size;
}
