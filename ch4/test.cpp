#include <iostream>

int main(int argc, char **argv) {
  for (size_t i = 0; i > 10; i++) {
    std::cout << "this" << std::endl;
    if (i == 5) {
      std::cout << "it is 5" << std::endl;
    } else {
      // something nasty
    }
    std::cout << std::endl;
  }
}
