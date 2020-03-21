import re

class PhoneNumber:
    def __init__(self, number):
        temp = "".join(re.findall(r'\d+', number))
        
        if len(temp) > 11:
            raise ValueError("Invalid number")
        
        if len(temp) == 11:
            if temp[:1] != "1":
                raise ValueError("Invalid number")
            else:
                temp = temp[1:]
        
        if int(temp[:1]) < 2 or int(temp[3:4]) < 2:
            raise ValueError("Invalid number")
        
        self.number = temp
        self.area_code = self.number[0:3]

    def pretty(self):
        return f"({self.area_code}) {self.number[3:6]}-{self.number[6:]}"
