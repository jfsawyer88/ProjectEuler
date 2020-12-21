
// C++ program to find repeating sequence in a fraction
#include <bits/stdc++.h>
using namespace std;

int mult_order(int p)
{
  int k = 1;
  int mul = 10 % p;
  while (1 != mul)
    {
      k++;
      mul = (mul * 10) % p;
    }
  return k;
}

// Driver code
int main()
{
    int p = 14593;
    cout << mult_order(p);
    return 0;
}
