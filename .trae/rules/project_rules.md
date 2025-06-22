# 项目规则
- 不能使用cgo
- 涉及sqlite的操作必须使用github.com/glebarez/sqlite 库，而不是其他数据库。


# 前端项目规则
- 项目名称叫复式记账系统
- 前端项目必须使用vue3 + pinia + tailwind，且使用vite进行开发，使用pnpm进行包管理。
- 前端项目在web目录下
- 使用axios进行封装，封装的axios实例在src/api目录下，封装的axios实例必须使用baseURL进行请求，baseURL必须使用环境变量进行配置。
- api请求参数使用小写snake_case进行命名，返回数据使用camelCase进行命名。
- 前端项目必须使用pinia进行状态管理，使用pinia进行状态管理的组件必须使用组件库进行开发。

# js项目规则
- 项目名称叫复式记账系统
- 项目必须使用nodejs进行开发，使用pnpm进行包管理。
- 对象使用camelCase进行命名，数组使用snake_case进行命名。
- 转换为json的时候使用小写snake_case进行命名。