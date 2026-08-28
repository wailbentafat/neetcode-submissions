class Solution:
    def dailyTemperatures(self, temperatures):
        n = len(temperatures)
        answer = [0] * n
        stack = []

        for i in range(n):
            while stack and temperatures[i] > temperatures[stack[-1]]:#so look the problem was here when we dont find the stack  we will append first operation one of the errors tht i haev to fix it for myself is thinking to use stack u have to do it in two times fill it then use it  meanwhile the solution is trying to do both in the same time  so  wht i have to do now is  we stack the day number 0 1 2  the operation goes this way fill the first day compare the next day if its bigger then the days before then if its true the difference between the start of the stack and the i tht u use it is the waiting days 
                waiting_day = stack.pop()
                answer[waiting_day] = i - waiting_day
            stack.append(i)

        return answer