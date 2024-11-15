class MinStack:
    def __init__(self):
        self.stack = []
        self.min_stack = []

    def push(self, val: int) -> None:
        self.stack.append(val)

        if not self.min_stack or val <= self.min_stack[-1]:
            self.min_stack.append(val)

    def pop(self) -> None:
        if not self.stack:
            return

        top = self.stack.pop()

        if top == self.min_stack[-1]:
            self.min_stack.pop()

    def top(self) -> int:
        if not self.stack:
            return None

        return self.stack[-1]

    def getMin(self) -> int:
        if not self.min_stack:
            return None

        return self.min_stack[-1]
