#include <stdio.h>
#include <sys/time.h>

long gcd(long a, long b)
{
  long c;
  while (a != 0)
    {
      c = a;
      a = b % a;
      b = c;
    }
  return b;
}

long lcm(long a, long b)
{
  return (a * b) / gcd(a, b);
}

main()
{
  struct timeval tv1, tv2;
  gettimeofday(&tv1, NULL);

  long out = 1;
  long i;
  for (i = 1; i <= 20; i++)
    {
      out = lcm(out, i);
    }

  gettimeofday(&tv2, NULL);  

  printf("Execution time: %f seconds\n",
         (double) (tv2.tv_usec - tv1.tv_usec)/1000000 + 
         (double) (tv2.tv_sec - tv1.tv_sec));
  printf("Output: %ld\n", out);

  return 0;
}
