import math

def main():
    while True:
        display_calculator_art()
        choice = get_input()
        
        if choice == 7:
            print("Exiting")
            break
        
        perform_operation(choice)
        input("Press Enter to continue...")
    return

def display_calculator_art():
    print("""
     _____________________
    |  _________________  |
    | | Calc            | |
    | |_________________| |
    |  ___ ___ ___   ___  |
    | | 7 | 8 | 9 | | + | |
    | |___|___|___| |___| |
    | | 4 | 5 | 6 | | - | |
    | |___|___|___| |___| |
    | | 1 | 2 | 3 | | x | |
    | |___|___|___| |___| |
    | | . | 0 | = | | / | |
    | |___|___|___| |___| |
    |_____________________|
    """)
    
def get_input():
    print("Please select an opperation to perform")
    print("1. Addition")
    print("2. Subtraction")
    print("3. Multiplication")
    print("4. Division")
    print("5. Exponentiate")
    print("6. Find the Square Root (n1^n2)")
    print("7. Exit")
    
    while True:
        try:
            choice = int(input("Enter choice (#): "))
            if choice in [1, 2, 3, 4, 5, 6, 7]:
                return choice
            else:
                print("Invalid input.")
        except ValueError:
            print("Invalid Input")
            
def add(x, y):
    return x + y

def subtract(x, y):
    return x - y

def multiply(x, y):
    return x * y

def divide(x, y):
    if (y == 0):
        return "Cannot divide by 0"
    return x / y

def exponent(x, y):
    return x**y

def square_rt(x):
    return math.sqrt(x)

def perform_operation(choice):
    try:
        num1 = float(input("Enter first number: "))
        if choice not in [6]:
            num2 = float(input("Enter second number: "))
        
        if choice == 1:
            print(add(num1, num2))
        elif choice == 2:
            print(subtract(num1, num2))
        elif choice == 3:
            print(multiply(num1, num2))
        elif choice == 4:
            print(divide(num1, num2))
        elif choice == 5:
            print(exponent(num1, num2))
        elif choice == 6:
            print(square_rt(num1))
            
    except ValueError:
        print("Invalid Input")
        
if __name__ == "__main__":
    main()