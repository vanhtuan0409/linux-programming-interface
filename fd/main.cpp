#include <stdio.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <unistd.h>
#include <fcntl.h>
using namespace std;

int main() {
  char const *file = "tmp.txt";
  int fd1 = open(file, O_RDWR | O_CREAT | O_TRUNC, S_IRUSR | S_IWUSR);
  int fd2 = dup(fd1);
  int fd3 = open(file, O_RDWR);
  write(fd1, "Hello,", 6);
  write(fd2, "world", 6);
  lseek(fd2, 0, SEEK_SET);
  write(fd1, "HELLO,", 6);
  write(fd3, "Gidday", 6);
}
