class BankAccount:
    def __init__(self):
        self.balance = 0
        self.closed = True

    def get_balance(self):
        if self.closed:
            raise ValueError("Account closed")
        
        return self.balance

    def open(self):
        if not self.closed:
            raise ValueError("Account closed")
        
        self.closed = False

    def deposit(self, amount):
        if self.closed:
            raise ValueError("Account closed")
        
        if amount < 0:
            raise ValueError("Amount can't be negative")
        
        self.balance += amount
        
    def withdraw(self, amount):
        if self.closed:
            raise ValueError("Account closed")
        
        if amount < 0:
            raise ValueError("Amount can't be negative")
        
        if self.balance < amount:
            raise ValueError("Insufficient balance")

        self.balance -= amount

    def close(self):
        if self.closed:
            raise ValueError("Account already closed")
 
        self.balance = 0
        self.closed = True
