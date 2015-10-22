#include <stdio.h>
#include <sys/time.h>


int is_prime(int n)
{
  int d = 3;
  while (d *d <= n)
    {
      if (0 == n % d)
	{
	  return 0;
	}
      d += 2;
    }
  return 1;
}


main()
{
  struct timeval tv1, tv2;
  gettimeofday(&tv1, NULL);

  int i = 3;
  int cnt = 1;
  while (cnt < 10001)
    {
      if (is_prime(i))
	{
	  cnt += 1;
	}
      i += 2;
    }

  gettimeofday(&tv2, NULL);

  printf("Execution time: %f seconds\n",
         (double) (tv2.tv_usec - tv1.tv_usec)/1000000 + 
         (double) (tv2.tv_sec - tv1.tv_sec));
  printf("Output: %d\n", i - 2);

  return 0;
}
