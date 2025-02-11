def stocGradAscent1 (dataMatrix, classLabels,numIter = 150):
    m,n = shape(dataMatrix)
    weights = ones(n) 
    for j in range(numIter): dataIndex = range(m)
        for i in range(m):
            alpha = 4 / (1.0 + j + i) + 0.01
            randIndex = int(random.uniform(0,len(dataIndex)))
            h = sigmoid(sum(dataMatrix[randIndex]*weights))
            error = classLabels[randIndex] - h
            weights = weights + alpha * error * dataMatrix[randIndex]
            del(dataIndex[randIndex])
    return weights