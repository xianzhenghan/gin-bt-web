#include <iostream>
#include <string>

using namespace std;

class Adapter1 {
public:
    bool req(int input1, string input2, int* indd);
    bool rsp(string input2, int* cnfd);
};

bool Adapter1::req(int input1, string input2, int* indd) {
    // Implementation of req
    return true;
}

bool Adapter1::rsp(string input2, int* cnfd) {
    // Implementation of rsp
    return true;
}


int main() {
    Adapter1 myComponent;
    int input1 = 123;
    string input2 = "Request Data";
    int indd = 456;
    string input2 = "Response Data";
    int cnfd = 101112;

    // Calling REQ function
    myComponent.req(input1, input2, &indd);

    // Calling RSP function
    myComponent.rsp(input2, &cnfd);
    return 0;
}
