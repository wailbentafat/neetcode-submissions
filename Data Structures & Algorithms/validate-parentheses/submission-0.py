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

    def isValid(self, s: str) -> bool:
        # stack =[]
        # reverseStack=[]
        # for i in s:
        #     stack.append(i)
        # j=len(s)
        # while j >=0 :
        #     reverseStack.append(s[j])
        #     j=j-1
        i=0
        j=len(s)
        if j % 2 ==1:
            return False
        while i !=j:
            if s[j-1]!= self.reverse(s[i]):
                return False
            i=i+1
            j=j-1
        return True
        



        

        