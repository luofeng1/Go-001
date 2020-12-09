## 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

- 不应该抛出一个sql.ErrNoRows
- sql.ErrNoRows 是数据库返回的信息,但是对于service层来说,不关心sql返回的信息,只关心在当前业务场景中,它应该返回的信息
