# 使用教程
## 安装依赖

```
下载依赖，地址：
https://codeload.github.com/go-vgo/Mingw/zip/master

将压缩包解压到本地

下载依赖文件，地址： https://github.com/wppzxc/fanli_test/raw/master/third_party/file2clip.exe
将下载好的 file2clip.exe 文件放到上一步解压后的 bin 路径下
```
![image](https://github.com/wppzxc/fanli_test/blob/master/image/save_file2clip.png)
```
将解压后的 bin 路径（如：D:\software\sdk\mingw\bin）加入到系统环境变量 Path中
```
![image](https://github.com/wppzxc/fanli_test/blob/master/image/change_env.png)

![image](https://github.com/wppzxc/fanli_test/blob/master/image/save_env.png)
```
打开 cmd 窗口，输入 gcc -v 敲回车，如果有大量返回则依赖安装成功
```
![image](https://github.com/wppzxc/fanli_test/blob/master/image/check_env.png)

## 修改配置文件

```
将 config.yaml 文件与 fanli.exe 放到一起，用记事本打开 config.yaml 按需求修改相关配置 
例如：
```
![image](https://github.com/wppzxc/fanli_test/blob/master/image/change_config.png)


## 运行程序

```
将上一步配置文件配置的用户窗口单独拖出来
并将 restart.bat 与fanli.exe 放到一起，双击 restart.bat 即可运行程序
```
