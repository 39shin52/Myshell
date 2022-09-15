#include<stdio.h>

int check(int s){
  int count = 0;
  for(int i = 1; i <= s; i++){
    if(s % i == 0)count++; 
  }
  if(count == 2)return 1;
  return 0;
}

int main(){
  int n = 859433, count = 0;

  for(int j = 1; j <= n; j++){
    if(check(j) == 1){
      count++;
      printf("%d\n", count);
    }
  }
  printf("%d\n", count);
}