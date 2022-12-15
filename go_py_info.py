#! coding:utf-8

# "" 等价   ''  ,"""用于备注

# list列表--->golang :array
list_ = [1,2,3]

# 索引对比golang的切片slice
list_1 = list_[0:2]
# 字典dict对比golang的map --> make(map[string]int)
dict_ = {}
dict_["a"] =1  #  dict_----->{'a': 1}
# def 函数对比golang的函数func，def __name()---> func name()  ; def name()----> func Name()
# 内部函数对比golang 的小写函数，一般函数对比大写函数
def test(*args,**kwargs):
    print("test")  # pass break continue
#  lambda argument_list: expression
lambda_ = lambda x,y:x+y
# 内置函数filter,map,sort,reduce与匿名函数lambda表达式一起使用  --->  类比golang中函数作为参数进行传递
# filter函数。此时lambda函数用于指定过滤列表元素的条件。
filter(lambda x: x % 3 == 0, [1, 2, 3])
sorted([1, 2, 3, 4, 5, 6, 7, 8, 9], key=lambda x: abs(5-x))
map(lambda x: x+1, [1,2,3])
reduce(lambda a, b: '{}, {}'.format(a, b), [1, 2, 3, 4, 5, 6, 7, 8, 9])
# 例如指定将列表[1,2,3]中能够被3整除的元素过滤出来，其结果是[3]。
# sorted函数。此时lambda函数用于指定对列表中所有元素进行排序的准则。例如sorted([1, 2, 3, 4, 5, 6, 7, 8, 9], key=lambda x: abs(5-x))将列表[1, 2, 3, 4, 5, 6, 7, 8, 9]按照元素与5距离从小到大进行排序，其结果是[5, 4, 6, 3, 7, 2, 8, 1, 9]。
# map函数。此时lambda函数用于指定对列表中每一个元素的共同操作。例如map(lambda x: x+1, [1, 2,3])将列表[1, 2, 3]中的元素分别加1，其结果[2, 3, 4]。
# reduce函数。此时lambda函数用于指定列表中两两相邻元素的结合条件。例如reduce(lambda a, b: '{}, {}'.format(a, b), [1, 2, 3, 4, 5, 6, 7, 8, 9])将列表 [1, 2, 3, 4, 5, 6, 7, 8, 9]中的元素从左往右两两以逗号分隔的字符的形式依次结合起来，其结果是'1, 2, 3, 4, 5, 6, 7, 8, 9'。


# class 类 对比 golang中的接口，类型和方法
# method的继承的部分，就是对比golang中的类型嵌套（匿名struct）
# method重写---> 重写
class C():
    def __init__(self):
        pass

    def func1(self):
        pass
    @staticmethod
    def fun2(args):
        print(args)
# 闭包回头再说
# golang没有的 ,yield ,set() 去重
