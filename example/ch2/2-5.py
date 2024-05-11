def classifyPerson()：
    resultList=['not at all','in small doses','in large doses']
    percentTats=float(raw_input(\
        "percentage of time spent playing video games?"))
    ffMiles=float(raw_input("freguent flier miles earned per year?"))
    iceCream=float(raw_input("liters of ice cream consumed per year?"))
    datingDataMat,datingLabels=file2matrix('datingTestSet2.txt')
    normMat,ranges,minVals=autoNorm(datingDataMat)
    inArr=array([ffMiles,percentTats,iceCream])
    classifierResult=classifyO((inArr-\
        minVals)/ranges,normMat,datingLabels,3)
    print "You will probably like this person：",\.
     resultList[classifierResult-1]
