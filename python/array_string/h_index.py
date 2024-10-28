from typing import List


class Solution:
    def h_index(self, citations: List[int]) -> int:
        citations.sort(reverse=True)

        h = 0

        for i, c in enumerate(citations):
            if c >= i + 1:
                h = i + 1
            else:
                break

        return h