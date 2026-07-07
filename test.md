

## 1. how to print

``` cpp
#include <iostream>

int main()

{

    std::cout<<"hello deamon" <<std::endl; //or you can use <<"\n"
    std::cout<<"welcome to the code";

    return 0;

}
```

##  2. variables

```cpp
#include <iostream>

int main() {

    int x ; // Declare an integer variable x without initialization

    x= 10; // Assign the value 10 to x

    std::cout<< x << "\n";

  
  

    return 0;

}


// or i can do-

#include <iostream>

int main() {

    int x =10;
    std::cout<<x;
    return 0;

}

```

# WHILE USING THE .sh  SCRIPTS MAKE SURE USE ```chmod +x filename.sh

# a script for compiling
```bash 
#!/bin/bash

  

# Change this to your C++ file name

SOURCE_FILE="variable.cpp"

OUTPUT_FILE="variableoutput.out"

  

# Compile the C++ code

echo "Compiling $SOURCE_FILE..."

g++ "$SOURCE_FILE" -o "$OUTPUT_FILE"

  

# Check if compilation was successful

if [ $? -eq 0 ]; then

    echo "Compilation successful."

    echo "Running the program:"

    echo "---------------------"

    ./"$OUTPUT_FILE"

else

    echo "Compilation failed. Please fix the errors above."

fi
```

## ANOTHER TIP INSTEAD OF USING  `std::cout<<"hi";`
USE using namespace std; (# you now u dont have to write that std:: thing again)

EXAMPLE
```cpp
#include <iostream>

using namespace std;

int main() {

    int x ; // Declare an integer variable x without initialization

    int y;

    x= 10; // Assign the value 10 to x

    y=29;

    int sum = x + y;

  

    cout<< x << "\n";

    cout<< y <<"\n";

    cout<< sum <<"\n";

  
  

    return 0;

}
```

# DATA TYPES


```cpp

int myNum = 5;               // Integer (whole number)  

float myFloatNum = 5.99;     // Floating point number  

double myDoubleNum = 9.98;   // Floating point number  

char myLetter = 'D';         // Character  

bool myBoolean = true;       // Boolean  

string myText = "Hello";     // String

  
#3 const //constants which u donnot want to change in future

  

#include <iostream>


using namespace std;

int main() {


    const double PI = 3.14 ; //here is a constant PI wegenerally use uppercase

    double r= 10;
    double circumference = 2*PI*r;

    cout<< circumference<<" cm" ;
    return 0;

}

```

 # Key Differences

- **Precision:**
    
    - **float** provides approximately 7 decimal digits of precision.
        

- **double** provides about 15 decimal digits of precision, offering nearly twice the accuracy of float
    

**Memory Size:**

- **float** is a 32-bit (4 bytes) data type

**double** is a 64-bit (8 bytes) data type
  
```
``` 

# NAME SPACE

(gives multiple values to a same variable) or it provides preventing conflict in names

```cpp

#include <iostream>

using namespace std;

namespace first{
    int x = 1;

}

namespace second{

  

    int x=2;

  

}

  

  

int main() {
    int x= 0;
   cout<< x; /* with cout<<x; output will be 0 where as if you write     cout<<first:: x; (first name space will be returned) */
    return 0;
}```




# CIN>> FOR INPUT
```cpp
#include <iostream>

  

using namespace std;

int main() {

    std::string name;

    std::getline(std::cin, name); ///to get strings with spaces use getline

    int age;

    std::cout<<"what is your name   ";  

    std::cin>> name;

    std::cout<<"what is your age  ";

    std::cin>> age;

    std::cout<<"hello  " <<name;

    std::cout<<"you are " <<age <<"years old";

  

}
```

max
```cpp
#include <iostream>

  

using namespace std;

int main()

{

    double a = 5;

    double b = 2;

    double z;

    z= std::max(a,b);

    std::cout <<z;

    return 0;

}
```

maths 

```cpp
#include <iostream>

#include <cmath>

using namespace std;

int main()

{

    double a = 5.2;

    double b = 2;

    double z;

    //z= std::min(a,b);

    //z= pow(3,3);

    //z= sqrt(16);

    //z =abs(-5); //absolute value

    //z= round(a);

    //z= ceil(a); //ceil value round up to upper value

    //z=floor(a); //floor value round down to lower value

    std::cout <<z;

    return 0;

}
```


hypotenuse cal

```cpp
#include <iostream>

#include <cmath>

using namespace std;

int main()

{

    cout<<"hypotenuse calculator"<<endl;

    double a;

    double b;

    double c;

    cout << "enter the side a" << endl;

    cin>>a;

    cout << "enter the side b" << endl;

    cin>>b;

    a= pow(a,2);

    b= pow(b,2);

    c= sqrt(a+b);

    cout << "hypotenuse is" << endl;

    cout<<c<<endl;

    return 0;

}
```

if , else , else if

```cpp
#include <iostream>

#include <cmath>

using namespace std;

int main()

{

    int age;

    cout<<"Enter your age: ";

    cin>>age;

    if(age>=18){

        cout<<"you are eligible for voting"<<endl;

    }

    else if(age<0){

        cout<<"invalid age"<<endl;

    }

    else{

        cout<<"you are not eligible for voting"<<endl;

    }

    return 0;

}
```


switches (alternate for if else)


```cpp
#include <iostream>

#include <cmath>

using namespace std;

int main()

{

    int month;

    cout<<"Enter your month(1-12): ";

    cin>>month;

    switch(month){

        case 1:

            cout<<"jan";

            break;

  

        case 2:

            cout << "feb";

            break;

  

        case 3:

            cout << "mar";

            break;

        case 4:

            cout << "apr";

            break;

        case 5:

            cout << "may";

            break;

        case 6:

            cout << "jun";

            break;

  

        case 7:

            cout << "jul";

            break;

  

        case 8:

            cout << "aug";

            break;

  

        case 9:

            cout << "sep";

            break;

        case 10:

            cout << "oct";

            break;

  

        case 11:

            cout << "nov";

            break;

  

        case 12:

            cout << "dec";

            break;
            
             
        default:

        cout<<"enter only 1-12";

            break;    

        }

  
  

    return 0;

}
```


simple cal program

```cpp
#include <iostream>

#include <cmath>

using namespace std;

int main()

{

    char op;

    double num1;

    double num2;

    double result;

    cout << "**********calculator**********\n";

    cout << "Enter operator (+, -, *, /): ";

    cin>> op;

  
  

    cout<<"enter num1  ";

    cin>>num1;

  
  

    cout << "enter num2  ";

    cin >> num2;

  

    switch(op){

        case '+':

            result = num1+num2;

            cout << result << endl;

            break;

  

        case '-':

            result = num1 - num2;

            cout << result << endl;

            break;

  

        case '*':

            result = num1 * num2;

            cout << result << endl;

            break;

  

        case '/':

            result = num1 / num2;

            cout << result << endl;

  

            break;

        default:

            cout << "enter a given vaild operator" << endl;

            break;

        }

  

    cout << "****************************";

  

    return 0;

}
```


&& operator

```cpp
#include <iostream>

using namespace std;

int main()

{

    //&& checks if both conditions are true

    // || checks if at least one condition is true like an or statement

    // ! reverse the logical state of its operand

  

    int temp;

    cout<<"Enter the temperature: ";

    cin>>temp;

    if(temp>0 && temp<30){

  

        cout<<"the temp is fine";

    }

    else{

        cout<<"temp is bad";

    }

  

    return 0;

}
```

||
```cpp
#include <iostream>

using namespace std;

int main()

{

    //&& checks if both conditions are true

    // || checks if at least one condition is true like an or statement

    // ! reverse the logical state of its operand

  

    int temp;

    cout<<"Enter the temperature: ";

    cin>>temp;

    if(temp<=0 || temp>=30){

  

        cout<<"the temp is fine";

    }

    else{

        cout<<"temp is bad";

    }

  

    return 0;

}
```


TEMP PROGRAM:

```cpp
#include <iostream>

using namespace std;

  

int main ()

  
  

{

    double temp ;

    char unit;

    cout<<"** **temp conversion program** **\n";

    cout<<"** ** F= Farenheit ** **\n";

    cout << "** ** C= Celsius ** **\n";

    cout << "** ** What unit you want to convert to** **\n";

    cin>> unit;

    if(unit == 'F' || unit == 'f'){

        cout<<"enter the temp in celsius \n";

        cin>>temp;

  

        temp = (1.8 *temp + 32);

        cout<<"Temp is :" <<temp <<"F\n";

    }

    else if(unit == 'C' || unit == 'c'){

        cout << "enter the temp in farenheit \n";

        cin >> temp;

        temp =(temp - 32 )/1.8 ;

        cout << "Temp is :" << temp << "C\n";

    }

    else {

        cout<<"**pls enter only C or F**\n";

    }

  

    return 0;

}
```



#### Why Use `getline` Instead of `cin >> name`?

The key difference is how they handle **whitespace** (spaces, tabs, etc.).

- `cin >> name;` stops reading as soon as it encounters any whitespace. If you type "Ada Lovelace", `cin` will only read "Ada" into the `name` variable.
    
- `getline(cin, name);` reads the entire line until you press Enter. It will correctly read "Ada Lovelace" into the `name` variable.

```cpp
#include <iostream>

using namespace std;

  

int main ()

  
  

{

   string name;

   cout<<"enter your name \n";

   getline(cin, name);

   if (name.length()>12){

    cout<<"you name cant be greater than 12 characters ";

   }

   else{

    cout<<"welcome  " <<name;

   }

    return 0;

}
```

while loop

```cpp
#include <iostream>

#include <string>

  

int main()

{

    std::string name;

  

    while (name.empty())

    {

        std::cout << "Enter your name: ";

  

        std::getline(std::cin, name);

    }

  

    std::cout << "Hello, " << name << "!";

  

    return 0;

} 
```


```cpp
#include <iostream>

  
  

int main()

{

    int number;

    do{

        std::cout<<"enter a +ve number";

        std::cin>> number;

    }

    while (number < 0);

    std::cout << "number is : " << number;

  

    return 0;

  

}
```



NESTED LOOPS

```cpp
#include <iostream>

using namespace std;

  
  

int main()

{   int rows;

    int columns;

    char symbol;

    cout<<"enter the number of rows :";

    cin>>rows;

  

    cout<<"enter the number of columns :";

    cin>>columns;

    cout << "enter the Symbol you want to use :";

    cin >> symbol;

  
  

    for(int i = 1; i<=rows; i++){

        for(int j =1; j<=columns; j++){

            cout << symbol;

        }

    cout<<"\n";}

}
```


RANDOMNESS

```cpp
#include <iostream>

using namespace std;

#include <ctime>

  

int main()

{

    srand(time(NULL));

    int num = (rand()%6) +1;

    cout<<num;

}
```


RANDOM LUCKY DRAW USING CTIME AND RAND ALSO SWITCHES

```cpp 
#include <iostream>

using namespace std;

#include <ctime>

  

int main()

{

    srand(time(0));

    int randnum = (rand()%5) +1;

    switch(randnum){

        case 1 :cout <<"you win a water bottle !! \n";

            break ;

        case 2:

            cout << "better luck next time !! \n";

            break;

        case 3:

            cout << "you win a gift card !! \n";

            break;

        case 4:

            cout << "you win a free lunch !! \n";

            break;

        case 5:

            cout << "you win a trip to london !! \n";

            break;

        }

  
  

}
```


GUESS THE NUMBER GAME


```cpp
#include <iostream>

using namespace std;

#include <ctime>

  

int main()

{

   int num;

   int tries = 0;

   int guess;

    srand(time(NULL));

    num = (rand()%100) +1;

    cout<<"** ** ** GUESS THE NUMBERR! ** ** ** \n";

  
  

    do{

        cout<<"enter a guess bw 1 to 100 \n";

        cin>> guess;

        tries++;

  

        if(guess>num)

{

    cout<<"you are too high \n";

}

else if (guess<num)

{

    cout<<"too low\n";

  

}

else {

    cout<<"YOU WINN!!! NUMBER OF TRIES : "<< tries <<"\n";

  

}

  

    }while(guess != num);

    return 0;

}
```


FUNCTIONS

```CPP
#include <iostream>

  

void happybday(std::string name , int age);

int main(){

  

    std::string name= "Hydroo";

    int age = 18;

    happybday(name, age);

    return 0;

}

void happybday(std::string name , int age ){

    std::cout<<"happybday \n" <<name <<"\n";

    std::cout << "happybday \n"<<name;

    std::cout << "happybday \n"<< name;

    std::cout << "happybday \n"<<name;

    std::cout << "happybday \n"<<name;

}
```

return keyword and other uses for function
```cpp
//square fn

//cube fn

  

#include <iostream>

double square (double length);

double cube (double length);

int main(){

  

    double length = 5.0;

    double area = square(length);

    double volume = cube(length);

    std::cout<<"the area is -\n" <<area <<"\n";

    std::cout<<"the volume is -\n" <<volume <<"\n";

    return 0;

  
  
  

}

double square (double length){

    return length*length;

  

}

double cube(double length)

{

    return length * length * length;}
```

BANKING
```cpp
#include <iostream>

#include <string> // Required library

void showBalance(double balance);

double deposit();

double withdraw(double balance);

using namespace std;

  

int main() {

    double balance = 0;

    int choice =0;

  
  

    do{

    cout<<"Enter you Choice: \n";

    cout<<"1 -- Show Balance\n";

    cout<<"2 -- deposit money\n";

    cout<<"3 -- withdraw money\n";

    cout<<"4 -- exit\n";

    cin >> choice;

  

    switch(choice){

        case 1: showBalance(balance);

                break;

        case 2: balance += deposit();

                showBalance(balance);

                break;

        case 3: balance-=withdraw(balance);

                showBalance(balance);

                break;

        case 4: cout<< "thanks for visiting \n";

                break;

        default: cout<< "invalid choice \n";

                break;

  
  

    }

  

    } while(choice != 4);

  
  

    return 0;

}

void showBalance(double balance){

    cout<< "Your balance is $" <<balance <<'\n';

}

double deposit(){

    double amount = 0;

  

    cout<<"Enter amount to deposit\n" ;

    cin>> amount;

  

   if (amount>0){

     return amount ;

   }

   else{

    cout<<"thats negative amount sir \n";

    return 0;

   }

}

double withdraw(double balance){

    double amount = 0;

    cout<< "ENTER TO AMOUNT TO BE WITHDRAWN\n";

    cin >> amount;

    if(amount>balance){

        cout<<"insufficient funds";

        return 0;

    }

    else if(amount < 0){

        cout<< "thats not a vaild amount \n";

        return 0;

  

    }

    else{

        return amount;

        }

    return amount;

}
```


ROCK PAPERS SCISSORS

```cpp
#include <iostream>

#include<ctime>

using namespace std;

char getuserchoice();

char getcomputerchoice();

void showchoice(char choice);

void choosewinner(char player , char computer);

  
  

int main(){

    char player;

    char computer;

    player = getuserchoice();

    cout << "\nYou chose: " << player << endl;

    showchoice(player);

  
  

    computer = getcomputerchoice();

    cout << "\ncomputers choice: " << computer << endl;

    showchoice(computer);

    choosewinner(player ,computer);

  
  

    return 0;

}

char getuserchoice(){

    char player;

    do{

        cout<< "rock paper scissors\n";

        cout<<"R for rock\n";

        cout<<"P for paper\n";

        cout<<"S for scissors\n";

        cin>> player;

        cout<< player;

    }while(player!='R' && player!='P' && player!='S');

    return player;

  

}

char getcomputerchoice(){

  

    srand(time(0));

    int num = rand() % 3 +1;

    switch(num){

        case 1: return 'R';

        case 2 : return 'P';

        case 3 : return 'S';

  

    }

    return 0;

  

}

void showchoice(char choice){

    switch(choice){

        case 'R' : cout<<"rock\n";

                break;

        case 'P' : cout<<"paper\n";

                break;

        case 'S' : cout<<"scissors\n";

                break;

  

        }

  

}

void choosewinner(char player , char computer){

    switch(player){

        case 'R': if(computer=='R'){

            cout<<"tie\n";

  

        }

        else if(computer=='P'){

            cout<<"you lose\n";

  
  

        }

        else{

            cout<<"you win \n";

        }

        break;

        case 'P': if(computer=='R'){

            cout<<"Win\n";

  

        }

        else if(computer=='P'){

            cout<<"Tie\n";

  
  

        }

        else{

            cout<<"Lose \n";

        }

        break;

  

        case 'S': if(computer=='R'){

            cout<<"lose\n";

  

        }

        else if(computer=='P'){

            cout<<"win\n";

  
  

        }

        else{

            cout<<"tie \n";

        }

        break;

  

    }

  

}
```


Arrays
```cpp
#include <iostream>

#include <string>

using namespace std;

  

int main(){

    string fruits[3] ;

    fruits[0] = "grapes";

    fruits[1] = "apples";

    fruits[2]= "kela";

  

    cout<<fruits[0]<<'\n';

    cout<<fruits[1]<<'\n';

    cout<<fruits[2]<<'\n';

    return 0;

}
```

ARRAY ITTERATION

```cpp
#include <iostream>

#include <string>

using namespace std;

#include <cstddef>  

  

int main(){

    string animals[] = {"dog" ,"cat" , "lion" , "zebra"};

  

    size_t count = sizeof(animals) / sizeof(string);

    for(size_t i= 0 ; i<count; i++){

  
  

        cout<<animals[i] << '\n';

  

    }

}
```



Memory location
```cpp
#include <iostream>

#include <string>

using namespace std;

int main(){

    string name = "hydro";

    int age = 32;

    bool student = true;

    cout << &name;

}
```

# boolean takes 1 byte of memory
# integers take 4 byte of memory


pass by using memory reference
```cpp
#include <iostream>
#include <string>
using namespace std;
void swap(string &x, string &y);
int main(){
    string x = "hello";
    string y = "niggers";
    
    
    swap(x,y);
    cout<< x << "\n";
    cout<< y <<"\n";
    
}

void swap(string &x, string &y){
    string temp;
    temp = x;
    x = y;
    y= temp;
     
}
```

#variable_swap

#pointers
```cpp
#include <iostream>

using namespace std;

#include <string>

int main(){

    string name = "bro";

    string *pName =&name;

    cout<<pName <<"\n";

    cout <<*pName;

}

  

//& address-of operator

//* dereference operator
```


#function #templates

```cpp
#include <iostream>
template <typename T>
T max(T x, T y){
    return (x > y) ? x : y;
}
int main()
{
    std::cout << max(1, 2) << '\n';
    std::cout << max(1.1, 2.2) << '\n';
    std::cout << max('1', '2') << '\n';
 
    return 0;
}
```

# pattern printing
```cpp
#include <iostream>

using namespace std;

void print1(int n){

for (int i = 0; i < n; i++)

{

for (int j = 0; j < n; j++)

{

cout << "*";

}

cout << endl;

}

}

int main()

{

int n;

cin>>n;

print1(n);

}
```


# triangle

```cpp
#include <iostream>

using namespace std;

void print1(int n){

for (int i = 0; i < n; i++)

{

for (int j = 0; j < i; j++)

{

cout << "*";

}

cout << endl;

}

}

int main()

{

int n;

cin>>n;

print1(n);

}
```
#number print

```cpp
#include <iostream>

using namespace std;

void print1(int n){

for (int i = 1; i < n; i++)

{

for (int j = 1; j <= i; j++)

{

cout << j << " ";

}

cout << endl;

}

}

int main()

{

int n;

cin>>n;

print1(n);

}
```

```cpp
#include <iostream>
using namespace std;
void print1(int n){
    
    for (int i=1; i<=n; i++){
        for (int j=0; j<n-i+1; j++){
            cout<<"*";
        }cout<<endl;

    }
    
}
int main()
{
    int n;
    cin>>n;
    print1(n);
}`
```

![[Pasted image 20260603175913.png]]

#palindrome

![[Pasted image 20260603223229.png]]

# armstrong numbers
![[Pasted image 20260603223350.png]]z

# gcd equiledian
![[Pasted image 20260604023839.png]]