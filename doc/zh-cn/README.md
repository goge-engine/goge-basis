# 简介  
**Goge Basis**是一个用于**Goge Engine**的基础包。  
该包提供了一些基础的结构  

## 文档  

`type Warning interface`一个警告接口，用于返回警告信息，它的子项为：  
- `Warning() string`返回警告信息

`type Warning struct`一个警告结构，用于记录警告信息，有一个`type Warning interface`的实现，它的子项为：  
- `message string`警告信息

`func NewWarning(warningMessage string) Warning`返回一个`warning`结构，参数为警告信息  

`func ReturnWarningForTest() Warning`返回一个`warning`结构，用于测试  