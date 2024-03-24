### 批量创建文件夹

##### 简介

一个通过go语言编写可以从excel文件中读取数据然后批量创建文件夹的软件

##### 使用方式

1. 将要批量创建的文件夹的名称写入到excel文件的sheet1中的第一列中，一个文件夹一行
2. 将excel文件命名为data.xlsx
3. 将data.xlsx文件与程序createFolder放在同一个目录下
4. 运行createFolder程序。注意windows系统运行createFolder.exe；macos系统运行createFolder
5. 运行完毕之后会在运行目录上创建created文件夹，所有批量创建的文件夹都在这个文件夹下面

##### 备注

1. 本项目已打包好了macos系统、windows系统下的可执行文件，存放于src目录下，可以直接下载使用
