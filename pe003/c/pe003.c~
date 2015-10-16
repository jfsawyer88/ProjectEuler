#include <stdio.h>
#include <sys/time.h>

main()
{
  struct timeval tv1, tv2;
  gettimeofday(&tv1, NULL);

  long N = 600851475143;
  int factor = 2;
  int last_factor = 1;

  while (N > 1)
    {
      if (0 == N % factor)
	{
	  last_factor = factor;
	  N /= factor;
	  while (0 == N % factor)
	    {
	      N /= factor;
	    }
	}
      factor += 1;
    }

  gettimeofday(&tv2, NULL);

  printf("Execution time: %f seconds\n",
	 (double) (tv2.tv_usec - tv1.tv_usec)/1000000 +
	 (double) (tv2.tv_sec - tv1.tv_sec));
  printf("Output: %d\n", last_factor);

  return 0;
}
