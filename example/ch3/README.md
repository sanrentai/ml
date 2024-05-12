# 决策树

## 优点：计算复杂度不高，输出结果易于理解，对中间值的缺失不敏感，可以处理不相关特征数据

## 缺点：可能会产生过度匹配问题

## 适用数据类型：数值型和标称型

首要问题，找到数据集上哪个特征划分数据分类时起决定性作用

### 决策树的一般流程

1. 收集数据：可以使用任何方法。
2. 准备数据：树构造算法只适用于标称型数据，因此数值型数据必须离散化。
3. 分析数据：可以使用任何方法，构造树完成之后，我们应该检查图形是否符合预期。
4. 训练算法：构造树的数据结构。
5. 测试算法：使用经验树计算错误率。
6. 使用算法：此步骤可以适用于任何监督学习算法，而使用决策树可以更好地理解数据的内在含义。

### 决策树构造算法

1. 找到具有最高信息增益的属性，该属性将用于分割数据集
2. 对每个属性划分后的数据集，递归调用步骤1，直到数据集中的所有实例都被标注了标签。
3. 构建决策树

### 伪代码createBranch

检测数据集中的每个子项是否属于同一分类：
    If so return 
    Else
        寻找到划分数据集的最好特征
        划分数据集
        创建分支节点
            for 每个划分的子集
                调用函数createBranch并增加返回结果到分支节点中
        return 分支节点

### 信息增益

信息增益 information gain

熵 entropy

熵越高，则混合的数据也越多

另一个度量集合无序程度的方法是基尼不纯度（Gini impurity）,简单地说就是从一个数据集中随机选取子项，度量其被错误分类到其他分组的概率。
这里不采用这个方法。

### 划分数据集

### 递归建决策树

原理：得到原始数据，然后基于最好的属性值划分数据集，由于特征值可能多于两个，因此可能存在大于两个分支的数据集划分。


## 代码实现

```python
from sklearn.tree import DecisionTreeClassifier
from sklearn.model_selection import train_test_split
from sklearn.metrics import classification_report
from sklearn import tree
import pandas as pd
import numpy as np

# 1. 准备数据
data = pd.read_csv('data.csv')
data.columns = ['年龄', '有工作', '有自己的房子', '信贷情况', '类别']

# 2. 数据预处理
# 假设数据已经预处理完毕，如果有缺失值，需要进行处理

# 3. 划分特征和标签
X = data[['年龄', '有工作', '有自己的房子', '信贷情况']]
y = data['类别']

# 4. 划分训练集和测试集 
X_train, X_test, y_train, y_test = train_test_split(X, y, test_size=0.3, random_state=1)

# 5. 创建决策树分类器
clf = DecisionTreeClassifier()

# 6. 训练模型
clf.fit(X_train, y_train)

# 7. 预测
y_pred = clf.predict(X_test)

# 8. 评估模型   
print(classification_report(y_test, y_pred))

# 9. 可视化决策树
tree.export_graphviz(clf, out_file='tree.dot', feature_names=X.columns, class_names=['不通过', '通过'], filled=True, rounded=True)

# 10. 模型保存
from sklearn.externals import joblib
joblib.dump(clf, 'model.pkl')

# 11. 模型加载
clf = joblib.load('model.pkl')
```
    
## 参考

- [决策树](https://www.cnblogs.com/pinard/p/6052321.html)
- [sklearn决策树](https://www.cnblogs.com/pinard/p/6149677.html)
- [sklearn决策树可视化](https://www.cnblogs.com/pinard/p/6150212.html)
