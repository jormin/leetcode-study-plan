package day09

// 请你仅使用两个栈实现先入先出队列。队列应当支持一般队列支持的所有操作（push、pop、peek、empty）：
//
// 实现 MyQueue 类：
// void push(int x) 将元素 x 推到队列的末尾
// int pop() 从队列的开头移除并返回元素
// int peek() 返回队列开头的元素
// boolean empty() 如果队列为空，返回 true ；否则，返回 false
//
// 说明：
// 你只能使用标准的栈操作 —— 也就是只有push to top,peek/pop from top,size, 和is empty操作是合法的。
// 你所使用的语言也许不支持栈。你可以使用 list 或者 deque（双端队列）来模拟一个栈，只要是标准的栈操作即可。
//
// 进阶：
// 你能否实现每个操作均摊时间复杂度为 O(1) 的队列？换句话说，执行 n 个操作的总时间复杂度为 O(n) ，即使其中一个操作可能花费较长时间。
//
// 示例：
// 输入：
// ["MyQueue", "push", "push", "peek", "pop", "empty"]
// [[], [1], [2], [], [], []]
// 输出：
// [null, null, null, 1, 1, false]
// 解释：
// MyQueue myQueue = new MyQueue();
// myQueue.push(1); // queue is: [1]
// myQueue.push(2); // queue is: [1, 2] (leftmost is front of the queue)
// myQueue.peek(); // return 1
// myQueue.pop(); // return 1, queue is [2]
// myQueue.empty(); // return false
//
// 提示：
// 1 <= x <= 9
// 最多调用 100 次 push、pop、peek 和 empty
// 假设所有操作都是有效的 （例如，一个空的队列不会调用 pop 或者 peek 操作）
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/implement-queue-using-stacks
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// MyQueue 队列
type MyQueue struct {
	in  []int
	out []int
}

// Constructor  构造器
func Constructor() MyQueue {
	return MyQueue{}
}

// Push 将元素 x 推到队列的末尾
func (q *MyQueue) Push(x int) {
	q.in = append(q.in, x)
}

// move 将元素从in移动到out
func (q *MyQueue) move() {
	for len(q.in) > 0 {
		q.out = append(q.out, q.in[len(q.in)-1])
		q.in = q.in[:len(q.in)-1]
	}
}

// Pop 从队列的开头移除并返回元素
func (q *MyQueue) Pop() int {
	if len(q.out) == 0 {
		q.move()
	}
	x := q.out[len(q.out)-1]
	q.out = q.out[:len(q.out)-1]
	return x
}

// Peek 返回队列开头的元素
func (q *MyQueue) Peek() int {
	if len(q.out) == 0 {
		q.move()
	}
	return q.out[len(q.out)-1]
}

// Empty 队列是否为空
func (q *MyQueue) Empty() bool {
	return len(q.in) == 0 && len(q.out) == 0
}
