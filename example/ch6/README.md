# 支持向量机

## SVM Support Vector Machine

### SMO Sequential Minimal Optimization 

序列最小优化

创建一个alpha向量并将其初始化为0向量
当迭代次数小于最大迭代次数时（外循环）
    对数据集中的每个数据向量（内循环）
        如果该数据向量可以被优化：
            随机选择另一个数据向量
            同时优化这两个向量
            如果两个向量都不能被优化，退出内循环
    如果所有向量都不能被优化，增加迭代数目，继续下一次循环

### 核函数 kernel

### 支持向量 support vector
