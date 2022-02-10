#include "stdio.h"

void swap(int *a, int *b) {
	int t = *a;
	*a = *b;
	*b = t;
	return;
};

int main(void) {
	int a = 21;
	int b = 17;

	swap(&a, &b);
	printf("main: a = %d, b = %d\n", a, b);
	return 0;
}
