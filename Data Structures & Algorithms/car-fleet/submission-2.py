class Solution:
    def carFleet(self, target: int, position: List[int], speed: List[int]) -> int:
        if len(position)!=len(speed) or len(speed)==0 or len(position)==0:
            return 0
        stack=[]
        i=0
        filter ={}

        order=[]
        while i<len(position):
            order.append(i)
            i=i+1
        order.sort(key=lambda x: position[x], reverse=True)

        i=0
        while i<len(order):
            stack.append((target-position[order[i]])/speed[order[i]])
            i=i+1

        counter=0
        fleet_time=-1
        i=0
        while i<len(stack):
            if stack[i]>fleet_time:
                counter+=1
                fleet_time=stack[i]
            i=i+1
        return counter