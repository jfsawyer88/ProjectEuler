#include <stdio.h>
#include <sys/time.h>

long mpow(long a, long b, long n)
{
  long out = 1;
  while (b > 0)
    {
      out = (out * a) % n;
      b--;
    }
  return out;
}

main()
{
  struct timeval tv1, tv2;
  gettimeofday(&tv1, NULL);

  long res = 0;
  int i;
  for (i = 1; i <= 1000; i++)
    {
      res = (res + mpow(i, i, 10000000000)) % 10000000000;
    }

  gettimeofday(&tv2, NULL);

  printf("Execution time: %f seconds\n",
         (double) (tv2.tv_usec - tv1.tv_usec)/1000000 + 
         (double) (tv2.tv_sec - tv1.tv_sec));
  printf("Output: %ld\n", res);

  return 0;
}

