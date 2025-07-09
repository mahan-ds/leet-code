#include <iostream>
#include<vector>
#include<unordered_map>
using namespace std;

int main()
{
    int size;
    int target;
    
    cout << "enter size of array : ";
    cin >> size;
    
    vector<int> numArray(size);

    for (int i = 0; i < size; i++) {
        cout << "enter index [" << i << "] of array : \n";
        cin >> numArray[i];
    }

    cout << "enter your target : ";
    cin >> target;

    unordered_map<int, int> map;

    for(int i = 0; i < size; i++)
    {
        int complement = target - numArray[i];

        if (map.find(complement) != map.end()) {
            cout << "output : [" << map[complement] << "," << i << "]";
        }

        map[numArray[i]] = i;
    }
   
}
