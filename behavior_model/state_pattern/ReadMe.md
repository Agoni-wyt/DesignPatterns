## 状态模式

状态模式是一种行为设计模式， 让你能在一个对象的内部状态变化时改变其行为， 使其看上去就像改变了自身所属的类一样。

## 角色组成

- Context（上下文）：是包含状态的对象，它可以处理请求，并将请求委托给当前状态对象。
- State（抽象状态）：定义一个接口，用于封装与Context的一个特定状态相关的行为。
- ConcreteState（具体状态）：每一个具体状态都提供了上下文的一部分行为。
- Client（客户端）：通过Context和State接口与上下文中的状态进行交互。
- 优点
  - 单一职责原则。 将与特定状态相关的代码放在单独的类中。
  - 开闭原则。 无需更改现有状态类和上下文就能引入新状态。
  - 简化代码。 将与特定状态相关的代码分布在状态类的方法中， 使得代码更加清晰。

## 应用场景
1. 如果对象需要根据自身当前状态进行不同行为， 同时状态的数量非常多且与状态相关的代码会频繁变更的话， 可使用状态模式。
2. 如果某个类需要根据成员变量的当前值改变自身行为， 从而需要使用大量的条件语句时， 可使用该模式。 
3. 当相似状态和基于条件的状态机转换中存在许多重复代码时， 可使用状态模式。