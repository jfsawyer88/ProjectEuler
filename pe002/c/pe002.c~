#include <stdio.h>
#include <sys/time.h>

main()
{
  struct timeval tv1, tv2;
  gettimeofday(&tv1, NULL);

  int f1 = 1;
  int f2 = 2;
  int res = 0;

  while (f2 < 4000000)
    {
      if (0 == f2 % 2)
	{
	  res += f2;
	}

      f2 = f1 + f2;
      f1 = f2 - f1;
    }

  gettimeofday(&tv2, NULL);

  printf("Execution time: %f seconds\n",
	 (double) (tv2.tv_usec - tv1.tv_usec)/1000000 +
	 (double) (tv2.tv_sec - tv1.tv_sec));
  printf("Output: %d\n", res);

  return 0;
}
