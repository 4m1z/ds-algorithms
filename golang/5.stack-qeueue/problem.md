## Problem-1: Valid Parentheses

**Description:** Given a string containing only parentheses, determine if it is valid. The string is valid if and only if all parentheses close in the correct order.

**Example:**
- Input: `"()"`, Output: `true`
- Input: `"()[]{}"`, Output: `true`
- Input: `"(]"`, Output: `false`
- Input: `"([)]"`, Output: `false`

## Problem-2: Valid HTML Tags

**Description:** Given a string containing HTML tags, determine if the tags are valid. Tags are valid if each opening tag has its corresponding closing tag.

**Example:**
- Input: `"<html><body></body></html>"`, Output: `true`
- Input: `"<div><p></div></p>"`, Output: `false`
- Input: `"<a href='example.com'>Link</a>"`, Output: `true`
- Input: `"This is <b>bold</b> text."`, Output: `true`

## Problem-3: Valid Round Brackets

**Description:** Given a string containing only round brackets "()" and lowercase characters, remove the least amount of brackets so that the string becomes valid. A string is valid if it is either empty or all round brackets are closed in the correct order.

**Example:**
- Input: `"()())()"`, Output: `"()()()"`
- Input: `"(a(b(c)d)"`, Output: `"(a(b(c)d))"`
- Input: `"abc)"`, Output: `"abc"`
- Input: `"()("`, Output: `"()"`

## Queue with Stack

**Description:** Implement a queue using two stacks. A queue is a data structure that follows the First-In-First-Out (FIFO) principle.

## Queue

**Description:** A queue is a data structure that stores elements and allows for two main operations: enqueue (add an element to the back) and dequeue (remove and return the element from the front). It follows the First-In-First-Out (FIFO) principle.

## Stack

**Description:** A stack is a data structure that stores elements and allows for two main operations: push (add an element to the top) and pop (remove and return the element from the top). It follows the Last-In-First-Out (LIFO) principle.






## problem-4:  

Find the nearest smaller element to the right of each element in an array. If there is no smaller element, return -1. 


Time Complexity: O(n)


## problem-5: 

Find the largest rectangle area in a histogram.

Time Complexity: O(n)




## problem-6: 


https://www.geeksforgeeks.org/problems/next-larger-element-1587115620/1


Expected Time Complexity : O(n)
Expected Auxiliary Space : O(n)




✅ Problem-7: Async Queue

Description:
Implement an asynchronous queue class with the following API:

    enqueue(item):
        if there are pending dequeue requests:
            resolve the oldest pending dequeue request with the item
        else:
            add the item to the internal queue
    
    dequeue():
        if the internal queue is not empty:
            return and remove the oldest item from the internal queue
        else:
            create a pending promise
            add the promise to the list of pending dequeue requests
            return the pending promise
    If the queue is empty, it waits until an item is available.

This is especially useful in producer-consumer scenarios where producers and consumers operate asynchronously.

Your queue must support the following behavior:

    If dequeue() is called before an item is enqueued, the promise should wait (i.e., it should be pending) until an item is added.

    If enqueue() is called and there are pending dequeue() calls, the item should be delivered to the oldest pending dequeue().

    If enqueue() is called and there are no pending dequeue() calls, the item should be added to the internal queue.

🔧 Constraints:
    The queue should preserve the order of both items and pending dequeues (FIFO for both).
