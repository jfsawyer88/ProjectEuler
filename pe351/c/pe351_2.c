#include <stdio.h>
#include <sys/time.h>
#include <malloc.h>
//int phi[N+1];

main()
{
  struct timeval tv1, tv2;
  gettimeofday(&tv1, NULL);

  int i, j;
  long output = 1;   // phi[1] = 1, but we won't sieve with it
  int N = 100000000; //input
  //int phi[N + 1];
  int *phi;
  phi = (int *)malloc(sizeof(int)*(N+1));

  // initialize phi[i] = i
  for (i = 0; i < N + 1; i++)
    {
      phi[i] = i;
    }
  
  for (i = 2; i < N + 1; i++)
    {
      // when i is prime sieve
      if (phi[i] == i)
	{
	  for (j = i; j < N + 1; j += i)
	    {
	      phi[j] = phi[j] - phi[j]/i;
	    }
	}
      output += phi[i];
    }

  gettimeofday(&tv2, NULL);

  printf("Execution time: %f seconds\n",
         (double) (tv2.tv_usec - tv1.tv_usec)/1000000 + 
         (double) (tv2.tv_sec - tv1.tv_sec));
  printf("Output: %ld\n", output);

  return 0;

}
