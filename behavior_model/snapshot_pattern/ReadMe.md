# 备忘录模式
## 角色组成
1. `Originator`：发起人，负责创建一个备忘录Memento，用以记录当前时刻它的内部状态，并可以使用备忘录恢复内部状态。Originator可以根据需要决定Memento存储Originator的哪些内部状态。
2. `Memento`：备忘录，用于存储Originator的内部状态，并且可以防止Originator以外的对象访问Memento。在备忘录Memento中有两个接口，Caretaker只能看到备忘录的窄接口，它只能将备忘录传递给其他对象。Originator可以看到一个宽接口，允许它访问返回先前状态所需的所有数据。
3. `Caretaker`：管理者，负责保存备忘录Memento，不能对备忘录的内容进行访问或者操作。
4. `Client`：客户端，负责创建备忘录Memento，用于记录Originator的内部状态，或者使用备忘录Memento恢复Originator的内部状态。

## 使用场景
备忘录模式让我们可以保存对象状态的快照。 你可使用这些快照来将对象恢复到之前的状态。 这在需要在对象上实现撤销-重做操作时非常实用。


