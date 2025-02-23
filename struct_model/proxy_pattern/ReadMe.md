# 代理模式（结构型设计模式）
## 角色组成
1. 抽象主题角色（Subject）：声明了真实主题和代理主题的共同接口，这样在任何使用真实主题的地方都可以使用代理主题。
2. 真实主题角色（RealSubject）：定义了代理角色所代表的真实对象。
3. 代理主题角色（Proxy）：代理主题角色内部包含对真实主题角色的引用，从而可以在任何时候操作真实主题对象；代理主题角色提供与真实主题角色相同的接口，以便在任何时候都可以替代真实主题角色；代理主题角色还可以控制对真实主题角色的使用，负责在需要的时候创建和删除真实主题角色，并且可以在真实主题角色处理请求之前或之后执行其他操作。
4. 客户端（Client）：通过代理主题角色间接访问真实主题角色。

## 使用场景
1. 延迟初始化 （虚拟代理）。 如果你有一个偶尔使用的重量级服务对象， 一直保持该对象运行会消耗系统资源时， 可使用代理模式。
2. 访问控制 （保护代理）。 如果你只希望特定客户端使用服务对象， 这里的对象可以是操作系统中非常重要的部分， 而客户端则是各种已启动的程序 （包括恶意程序）， 此时可使用代理模式。
3. 本地执行远程服务 （远程代理）。 适用于服务对象位于远程服务器上的情形。
4. 日志记录请求 （日志记录代理）。 适用于记录请求日志的情形。
5. 缓存请求结果 （缓存代理）。 适用于缓存请求结果并在后续请求时返回缓存结果的情形。
6. 智能引用。 可在没有客户端使用某个重量级对象时立即销毁该对象。
## 实战举例
Nginx 这样的 Web 服务器可充当应用程序服务器的代理：

提供了对应用程序服务器的受控访问权限。
可限制速度。
可缓存请求。