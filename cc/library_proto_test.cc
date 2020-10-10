#include <iostream>

#include "proto/library/library.pb.h"

int main() {
  library::Book book;
  book.set_name("The C++ Programming Language");
  std::cout << book.ShortDebugString() << std::endl;
  return 0;
}
