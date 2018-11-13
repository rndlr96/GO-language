#include <stdio.h>

int MatrixMulti(int i,int j,int Mat[]){

  if (i == j)
     return 0;

   int count, k;
   int min = 10000000;

  for (k = i; k < j; k++) {
    count = MatrixMulti(i, k, Mat) + MatrixMulti(k+1, j, Mat) +
            Mat[i-1] * Mat[k] * Mat[j];

    if (count < min)
      min = count;
  }
  return min;
}



int* input_array_number(int n)  // 행렬에 크기 넣기
{
   int i, *size;
   printf("행렬의 행의 값을 입력하시오:\n");
   size = (int*)malloc(sizeof(int)*n + 1);
   for (i =0; i < n; i++)
   {
      printf("%d번째 행렬의 행의 값은 :", i );
      scanf("%d", &size[i]);
      printf("%d번째 행렬의 열의 값은 :", i );
      scanf("%d", &size[i+1]);
   }
   for (i = 0; i < n; i++) {
      printf("행렬 %d:%d,%d\n",i,size[i],size[i+1]);
   }

   return size;
}


int main()
{
   int *Matrix;
   int number, min;

   printf("행렬의 개수를 입력하시오\n");
   scanf("%d", &number);

   Matrix = input_array_number(number);

   min = MatrixMulti(1, number, Matrix);

   printf("\n행렬 곱셈의 최솟값은 :%d", min);

   return 0;
}
