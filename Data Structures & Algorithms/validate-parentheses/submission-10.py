class Solution:
    def reverse (self,s:str)->str:
        if s == '(':
            return ')'
        if  s == '{':
            return '}'
        if s =='[':
            return ']'
        print("wrong strategie ")
        return 'F'
    def itsOpen(self,s:str)->bool:
        if s =='(' or s == '{' or s=='[':
            return True 
        return False

    def itsClose(self,s:str)->bool:
        if s =='}' or s == ')' or s ==']':
            return True 
        return False

    def isValid(self, s: str) -> bool:
        # i=0
        j=len(s)
        if j % 2 ==1:
            return False
        # while i !=j:
        #     if s[j-1]!= self.reverse(s[i]) and s[i+1] != self.reverse(s[i]):
        #         return False
        #     i=i+1
        #     j=j-1
        # return True
        stack = []
        i=0
        while i <len(s):
            if self.itsOpen(s[i]):
                stack.append(s[i])
            if self.itsClose(s[i]):
                if not stack :
                    return False
                check=stack.pop()
                if s[i]!=self.reverse(check):
                    return False
            i=i+1
        return len(stack) ==0
                


        



        

        