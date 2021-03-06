#include <stdio.h>
#include <sys/time.h>

long collatz(long n)
{
  if (0 == n % 2)
    {
      return n / 2;
    }
  else
    {
      return 3 * n + 1;
    }
}


main()
{
  struct timeval tv1, tv2;
  gettimeofday(&tv1, NULL);

  long collatz_array[1000000];
  collatz_array[1] = 1;

  long i, s, c;
  long m = 1;
  for (i = 2; i < 1000000; i++)
    {
      //printf("%ld", i);
      s = collatz(i);
      //printf("->%ld", s);
      c = 1;
      while (s >= i)
	{
	  s = collatz(s);
	  c += 1;
	  //printf("->%ld", s);
	}
      //printf("\n");
      collatz_array[i] = c + collatz_array[s];
      if (collatz_array[m] < collatz_array[i])
	{
	  m = i;
	}
    }

  gettimeofday(&tv2, NULL);

  printf("Execution time: %f seconds\n",
         (double) (tv2.tv_usec - tv1.tv_usec)/1000000 + 
         (double) (tv2.tv_sec - tv1.tv_sec));
  printf("Output: %ld\n", m);

  return 0;
}
