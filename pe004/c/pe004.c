#include <stdio.h>
#include <sys/time.h>
#include <string.h>

int reverse_int(int n)
{
  int res = 0;
  while (0 < n)
    {
      res = (res * 10) + (n % 10);
      n /= 10;
    }

  return res;
}

int is_palindrome(int n)
{
  return n == reverse_int(n);
}

main()
{

  int a;
  int b;

  struct timeval tv1, tv2;
  gettimeofday(&tv1, NULL);

  a = 999;
  int largest_palindrome = 11;
  while (a > 99)
    {
      b = 999;
      while (b > 99)
	{
	  if (a * b < largest_palindrome)
	    {
	      break;
	    }
	  if (is_palindrome(a * b))
	    {
	      largest_palindrome = a * b;
	    }
	  b -= 1;
	}
      a -= 1;
    }

  gettimeofday(&tv2, NULL);

  printf("Execution time: %f seconds\n",
	 (double) (tv2.tv_usec - tv1.tv_usec)/1000000 +
	 (double) (tv2.tv_sec - tv1.tv_sec));
  printf("Output: %d\n", largest_palindrome);

  return 0;
}
