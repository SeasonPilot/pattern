package main

// 创建原型接口， 并在其中声明 克隆方法。 如果你已有类层次结构， 则只需在其所有类中添加该方法即可。
type inode interface {
	print(string)
	clone() inode
}
