# 观察者模式
## 角色组成
1. 抽象主题：用于存储观察者对象的聚焦类和用于增加、删除观察者对象的方法，以及通知所有观察者的抽象方法
2. 观察者：包含更新自己的抽象方法，具体主题更改-->自身更新
3. 具体观察者：实现观察者抽象类或接口中定义的抽象方法，目标更改-->自身更新
4. 客户端：创建抽象主题对象和观察者对象，为观察者对象注册抽象主题信息

## 使用场景
1. 对象一对多时   主题1->观察者n
2. 一个对象改变需要改变其它对象
3. 一些对象需要观察其它对象时

订阅者和被订阅物

## 观察者模式组成部分
1. 任意事件发生时发布事件 的 主题
2. 订阅了主题事件并在事件发生时接收通知 的 观察者

